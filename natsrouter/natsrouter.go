// Package natsrouter provides a simple routing layer for NATS with unified request handling.
package natsrouter

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

// HandlerFunc defines a function to handle incoming messages and return a response.
type HandlerFunc func(msg *Message) (any, error)

// Router routes request/reply messages based on subject names.
type Router struct {
	nc       *nats.Conn
	prefix   string
	handlers map[string][]HandlerFunc
}

// New creates a new Router with the given NATS connection.
func New(nc *nats.Conn) *Router {
	return &Router{
		nc:       nc,
		handlers: make(map[string][]HandlerFunc),
	}
}

// WithHealth starts a goroutine that periodically publishes node health info.
func (r *Router) WithHealth(interval time.Duration, nodeInfo map[string]any) *Router {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			nodeInfo["timestamp"] = time.Now().Unix()
			data, err := json.Marshal(nodeInfo)
			if err != nil {
				log.Printf("[natsrouter] failed to marshal health info: %v", err)
				continue
			}
			err = r.nc.Publish("health.node", data)
			if err != nil {
				log.Printf("[natsrouter] failed to publish health info: %v", err)
			}
			<-ticker.C
		}
	}()
	return r
}

// NewWithPrefix creates a new Router with a specific prefix for subjects.
func (r *Router) Group(prefix string) *Router {
	return &Router{
		nc:       r.nc,
		prefix:   joinSubject(r.prefix, prefix),
		handlers: r.handlers, // bendras map (arba izoliuotas, jei reikia)
	}
}

// Handle registers a handler for a specific subject.
func (r *Router) Handle(subject string, handler HandlerFunc) {
	full := joinSubject(r.prefix, subject)
	r.handlers[full] = append(r.handlers[full], handler)
}

// HandleMany registers a handler for multiple subjects.
func (r *Router) HandleMany(subjects []string, handler HandlerFunc) {
	for _, s := range subjects {
		r.Handle(s, handler)
	}
}

// Listen subscribes to all registered subjects and processes responses if needed.
func (r *Router) Listen() error {
	for subject, handlers := range r.handlers {
		_, err := r.nc.Subscribe(subject, func(m *nats.Msg) {
			for _, handler := range handlers {
				msg := WrapMessage(m)
				resp, err := handler(msg)
				if err != nil {
					if m.Reply == "" {
						msg.MarkError(err)
						_ = msg.RespondSelf()
					}
					return
				}
				if resp == nil {
					return // no response to send
				}
				if m.Reply == "" {
					return // skip if no reply expected
				}
				if err := msg.RespondAny(resp); err != nil {
					log.Printf("[natsrouter] failed to respond: %v", err)
				}
			}
		})
		if err != nil {
			return err
		}
		log.Printf("[natsrouter] listening for: %s", subject)
	}
	return nil
}

func NewMessage(subject string) *Message {
	msg := nats.NewMsg(subject)
	return WrapMessage(msg)
}

// Message wraps a *nats.Msg and provides helper methods.
type Message struct {
	*nats.Msg
}

// WrapMessage creates a Message from *nats.Msg
func WrapMessage(m *nats.Msg) *Message {
	return &Message{Msg: m}
}

// JSON parses the incoming data as JSON into v.
func (m *Message) JSON(v interface{}) error {
	return json.Unmarshal(m.Data, v)
}

// Respond sends a raw data reply.
func (m *Message) Respond(data []byte) error {
	if m.Reply != "" {
		return m.Msg.Respond(data)
	}
	return nil
}

// RespondJSON sends a JSON-encoded reply.
func (m *Message) RespondJSON(v interface{}) error {
	if m.Reply == "" {
		return nil
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return m.Msg.Respond(data)
}

func (m *Message) RespondSelf() error {
	return m.Msg.RespondMsg(m.Msg)
}

// RespondAny sends a response depending on the type.
func (m *Message) RespondAny(v any) error {
	if v == nil || m.Reply == "" {
		return nil
	}

	switch val := v.(type) {
	case []byte:
		return m.Msg.Respond(val)
	case string:
		return m.Msg.Respond([]byte(val))
	case *Message:
		return m.Msg.RespondMsg(val.Msg)
	case *nats.Msg:
		return m.Msg.RespondMsg(val)
	default:
		return m.RespondJSON(val)
	}
}

func (m *Message) MarkError(err error) {
	if m.Header == nil {
		m.Header = nats.Header{}
	}
	m.Header.Set("X-NatsRouter-Error", err.Error())
}

func (m *Message) Error() error {
	if m.Header == nil {
		return nil
	}
	if errStr := m.Header.Get("X-NatsRouter-Error"); errStr != "" {
		return errors.New(errStr)
	}
	return nil
}

// IsError checks if the message has an error flag set.
func (m *Message) IsError() bool {
	return m.Header.Get("X-NatsRouter-Error") != ""
}

func joinSubject(parts ...string) string {
	var out []string
	for _, p := range parts {
		if p != "" {
			out = append(out, p)
		}
	}
	return strings.Join(out, ".")
}
