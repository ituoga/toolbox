package env

import (
	"errors"
	"os"

	"github.com/ituoga/toolbox/sconv"
)

var ErrNotExist = errors.New("environment variable does not exist")

func Get(k string) *sconv.ConverterWithoutErrors {
	if v, ok := os.LookupEnv(k); ok {
		return sconv.String(v)
	}
	return sconv.String("")
}

func GetWithError(k string) (*sconv.ConverterWithoutErrors, error) {
	if v, ok := os.LookupEnv(k); ok {
		return sconv.String(v), nil
	}
	return nil, ErrNotExist
}
