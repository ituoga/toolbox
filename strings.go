package toolbox

import "github.com/delaneyj/toolbelt"

var Pascal = toolbelt.Pascal
var Camel = toolbelt.Camel
var Snake = toolbelt.Snake
var ScreamingSnake = toolbelt.ScreamingSnake
var Kebab = toolbelt.Kebab
var Upper = toolbelt.Upper
var Lower = toolbelt.Lower

type CasedString = toolbelt.CasedString

var ToCasedString = toolbelt.ToCasedString

type CasedFn func(string) string

func Cased(s string, fn ...CasedFn) string {
	for _, f := range fn {
		s = f(s)
	}
	return s
}
