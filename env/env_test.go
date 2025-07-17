package env_test

import (
	"os"
	"testing"

	"github.com/ituoga/toolbox/env"
)

func TestMain(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "5")

	if env.Get("TEST_ENV_VAR").Int64() != 5 {
		t.Errorf("Expected 5, got %d", env.Get("TEST_ENV_VAR").Int64())
	}

	if env.Get("NON_EXISTENT_VAR").Int() != 0 {
		t.Error("Expected empty string for non-existent variable")
	}

	if _, err := env.GetWithError("TEST_ENV_VAR"); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if _, err := env.GetWithError("NON_EXISTENT_VAR"); err != env.ErrNotExist {
		t.Errorf("Expected ErrNotExist, got %v", err)
	}
	if v, err := env.GetWithError("TEST_ENV_VAR"); err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if v.Int() != 5 {
		t.Errorf("Expected 5, got %v", v.Int())
	}
}
