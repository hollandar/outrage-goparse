package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type AnyCharacterTokenizeFunc func([]rune) []tokens.Token

type AnyCharacterMatcher struct {
	Tokenize AnyCharacterTokenizeFunc
}

func (m AnyCharacterMatcher) Matches(input *content.Source) Match {
	if input.Length() >= 1 {
		var match Match
		match = NewMatchSuccess(m.Tokenize(input.Slice(0, 1)))

		return match
	}

	return NewMatchError("Any character was expected.")

}
