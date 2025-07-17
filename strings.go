package toolbox

import "github.com/delaneyj/toolbelt"

func Pascal(s string) string {
	return toolbelt.Pascal(s)
}

func Camel(s string) string {
	return toolbelt.Camel(s)
}

func Snake(s string) string {
	return toolbelt.Snake(s)
}

func ScreamingSnake(s string) string {
	return toolbelt.ScreamingSnake(s)
}

func Kebab(s string) string {
	return toolbelt.Kebab(s)
}

func Upper(s string) string {
	return toolbelt.Upper(s)
}

func Lower(s string) string {
	return toolbelt.Lower(s)
}

type CasedString struct {
	Original       string
	Pascal         string
	Camel          string
	Snake          string
	ScreamingSnake string
	Kebab          string
	Upper          string
	Lower          string
}

func ToCasedString(s string) CasedString {
	return CasedString{
		Original:       s,
		Pascal:         Pascal(s),
		Camel:          Camel(s),
		Snake:          Snake(s),
		ScreamingSnake: ScreamingSnake(s),
		Kebab:          Kebab(s),
		Upper:          Upper(s),
		Lower:          Lower(s),
	}
}

type CasedFn func(string) string

func Cased(s string, fn ...CasedFn) string {
	for _, f := range fn {
		s = f(s)
	}
	return s
}
