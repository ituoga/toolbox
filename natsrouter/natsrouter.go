// Package natsrouter provides a simple routing layer for NATS with unified request handling.
package natsrouter

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

// HandlerFunc defines a function to handle incoming messages and return a response.
type HandlerFunc func(msg *Message) (any, error)

// Router routes request/reply messages based on subject names.
type Router struct {
	nc       *nats.Conn
	handlers map[string]HandlerFunc
}

// New creates a new Router with the given NATS connection.
func New(nc *nats.Conn) *Router {
	return &Router{
		nc:       nc,
		handlers: make(map[string]HandlerFunc),
	}
}

// Handle registers a request/reply handler for a specific subject.
func (r *Router) Handle(subject string, handler HandlerFunc) {
	r.handlers[subject] = handler
}

// Listen subscribes to all registered subjects and processes responses if needed.
func (r *Router) Listen() error {
	for subject, handler := range r.handlers {
		_, err := r.nc.Subscribe(subject, func(m *nats.Msg) {
			if m.Reply == "" {
				return // skip if no reply expected
			}

			msg := WrapMessage(m)
			resp, err := handler(msg)
			if err != nil {
				_ = msg.RespondJSON(map[string]string{"error": err.Error()})
				return
			}
			if resp == nil {
				return // no response to send
			}
			if err := msg.RespondAny(resp); err != nil {
				log.Printf("[natsrouter] failed to respond: %v", err)
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
	if m.Reply == "" {
		return nil
	}
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
	case *nats.Msg:
		return m.Msg.RespondMsg(val)
	default:
		return m.RespondJSON(val)
	}
}
