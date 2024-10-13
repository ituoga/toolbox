package toolbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/valyala/bytebufferpool"
)

func ReaderFromRequest(r *http.Request) (io.ReadSeeker, error) {
	if r.Body == nil {
		return nil, fmt.Errorf("request body is nil")
	}
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	return bytes.NewReader(buf.Bytes()), nil
}

func BodyUnmarshal(r io.ReadSeeker, store any) error {
	r.Seek(0, 0)
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
