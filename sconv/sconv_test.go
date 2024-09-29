package sconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithErrors(t *testing.T) {
	t.Run("test error", func(t *testing.T) {
		_, err := StringWithError("teksras").Int()
		assert.True(t, err != nil)
	})
	t.Run("int", func(t *testing.T) {
		resp, _ := StringWithError("1").Int()

		if resp != int(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("int8", func(t *testing.T) {
		resp, _ := StringWithError("1").Int8()

		if resp != int8(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("int16", func(t *testing.T) {
		resp, _ := StringWithError("1").Int16()

		if resp != int16(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("int32", func(t *testing.T) {
		resp, _ := StringWithError("1").Int32()

		if resp != int32(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("int64", func(t *testing.T) {
		resp, _ := StringWithError("1").Int64()

		if resp != int64(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint", func(t *testing.T) {
		resp, _ := StringWithError("1").UInt()

		if resp != uint(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint8", func(t *testing.T) {
		resp, _ := StringWithError("1").UInt8()

		if resp != uint8(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint16", func(t *testing.T) {
		resp, _ := StringWithError("1").UInt16()

		if resp != uint16(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint32", func(t *testing.T) {
		resp, _ := StringWithError("1").UInt32()

		if resp != uint32(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint64", func(t *testing.T) {
		resp, _ := StringWithError("1").UInt64()

		if resp != uint64(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})

	t.Run("float32", func(t *testing.T) {
		resp, _ := StringWithError("1.123").Float32()

		if resp != float32(1.123) {
			t.Fatalf("expected 1.123 got %f", resp)
		}
	})
	t.Run("float64", func(t *testing.T) {
		resp, _ := StringWithError("1.123").Float64()

		if resp != float64(1.123) {
			t.Fatalf("expected 1.123 got %f", resp)
		}
	})

	t.Run("bool", func(t *testing.T) {
		resp, _ := StringWithError("").Bool()
		assert.False(t, resp)

		resp, _ = StringWithError("0").Bool()
		assert.False(t, resp)

		resp, _ = StringWithError("false").Bool()
		assert.False(t, resp)

		resp, _ = StringWithError("no").Bool()
		assert.False(t, resp)

		resp, _ = StringWithError("1").Bool()
		assert.True(t, resp)

		resp, _ = StringWithError("true").Bool()
		assert.True(t, resp)

		resp, _ = StringWithError("yes").Bool()
		assert.True(t, resp)
	})

}

func TestWithoutErrors(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		resp := String("1").Int()

		if resp != int(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("int8", func(t *testing.T) {
		resp := String("1").Int8()

		if resp != int8(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("int16", func(t *testing.T) {
		resp := String("1").Int16()

		if resp != int16(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("int32", func(t *testing.T) {
		resp := String("1").Int32()

		if resp != int32(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("int64", func(t *testing.T) {
		resp := String("1").Int64()

		if resp != int64(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint", func(t *testing.T) {
		resp := String("1").UInt()

		if resp != uint(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint8", func(t *testing.T) {
		resp := String("1").UInt8()

		if resp != uint8(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint16", func(t *testing.T) {
		resp := String("1").UInt16()

		if resp != uint16(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint32", func(t *testing.T) {
		resp := String("1").UInt32()

		if resp != uint32(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})
	t.Run("uint64", func(t *testing.T) {
		resp := String("1").UInt64()

		if resp != uint64(1) {
			t.Fatalf("expected 1 got %d", resp)
		}
	})

	t.Run("float32", func(t *testing.T) {
		resp := String("1.123").Float32()

		if resp != float32(1.123) {
			t.Fatalf("expected 1.123 got %f", resp)
		}
	})
	t.Run("float64", func(t *testing.T) {
		resp := String("1.123").Float64()

		if resp != float64(1.123) {
			t.Fatalf("expected 1.123 got %f", resp)
		}
	})

	t.Run("bool", func(t *testing.T) {
		resp := String("").Bool()
		assert.False(t, resp)

		resp = String("0").Bool()
		assert.False(t, resp)

		resp = String("false").Bool()
		assert.False(t, resp)

		resp = String("no").Bool()
		assert.False(t, resp)

		resp = String("1").Bool()
		assert.True(t, resp)

		resp = String("true").Bool()
		assert.True(t, resp)

		resp = String("yes").Bool()
		assert.True(t, resp)
	})

}
