package matchers

import (
	"fmt"

	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type ExactMatcherTokenizeFunc func([]rune) []tokens.Token

type ExactMatcher struct {
	Match    []rune
	Tokenize ExactMatcherTokenizeFunc
}

func (m ExactMatcher) Matches(input *content.Source) Match {
	if input.Length() >= len(m.Match) {
		value := input.Take(len(m.Match))
		if runeSequenceEqual(value, m.Match) {
			var match = NewMatchSuccess(m.Tokenize(value))
			input.Advance(len(value))

			return match
		}
		return NewMatchError(fmt.Sprintf("%s was expected, got %s.", string(m.Match), string(value)))
	}

	return NewMatchError(fmt.Sprintf("not enough input to match %s", string(m.Match)))
}

func runeSequenceEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
