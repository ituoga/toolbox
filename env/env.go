package env

import (
	"os"

	"github.com/ituoga/toolbox/sconv"
)

func Get(k string) *sconv.ConverterWithoutErrors {
	if v, ok := os.LookupEnv(k); ok {
		return sconv.String(v)
	}
	return nil
}
