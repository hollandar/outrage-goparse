package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
)

type Matcher interface {
	Matches(input *content.Source) Match
}
