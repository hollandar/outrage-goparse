package parser

import (
	"fmt"

	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/matchers"
)

func Parse(input string, rootMatcher matchers.Matcher) matchers.Match {
	source := content.NewSource(input)
	match := rootMatcher.Matches(source)

	if match.Success == false {
		return matchers.NewMatchError(fmt.Sprintf("Line %d, Col %d; %s; at %s", source.Line, source.Column, match.Error, string(source.AtPosition(30))))
	} else {
		return match
	}
}
