package toolbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/valyala/bytebufferpool"
)

func ReaderFromRequest(r *http.Request) (*bytes.Buffer, error) {
	if r.Body == nil {
		return nil, fmt.Errorf("request body is nil")
	}
	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	return buf, nil
}

func BodyUnmarshal(r io.Reader, store any) error {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	if _, err := buf.ReadFrom(r); err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}
	b := buf.Bytes()
	if err := json.Unmarshal(b, store); err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}
	return nil
}
