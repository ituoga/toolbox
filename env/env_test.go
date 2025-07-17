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

}
