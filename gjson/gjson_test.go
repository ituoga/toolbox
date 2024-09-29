package gjson_test

import (
	"testing"

	"github.com/ituoga/toolbox/gjson"
)

type testData struct {
	Name string `json:"name"`
}

func TestUnmarshal(t *testing.T) {
	data := []byte(`{"name":"test"}`)
	v, err := gjson.Unmarshal[testData](data)
	if err != nil {
		t.Fatal(err)
	}
	if v.Name != "test" {
		t.Fatalf("expected %q, got %q", "test", v.Name)
	}
}

func TestMustUnmarshal(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, got nil")
		}
	}()

	data := []byte(`{"name":"test"}`)
	v := gjson.MustUnmarshal[testData](data)
	if v.Name != "test" {
		t.Fatalf("expected %q, got %q", "test", v.Name)
	}

	data = []byte(`{"nameinvalid:"test"}`)
	v = gjson.MustUnmarshal[testData](data)
	if v.Name != "test" {
		t.Fatalf("expected %q, got %q", "test", v.Name)
	}
}

func TestString(t *testing.T) {
	data := []byte(`"some value"`)
	v, err := gjson.Unmarshal[string](data)
	if err != nil {
		t.Fatal(err)
	}
	if v != `some value` {
		t.Fatalf("expected %q, got %q", `some value`, v)
	}
}
