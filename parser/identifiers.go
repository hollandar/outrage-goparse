package parser

import (
	"github.com/hollandar/outrage-goparse/parser/matchers"
)

func Name() matchers.Matcher {
	return Or(
		LetterOrDigit(),
		Underscore(),
	)
}
