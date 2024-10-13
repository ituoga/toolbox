package toolbox

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/valyala/bytebufferpool"
)

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
