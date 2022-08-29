package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type CharacterMatchFunc func(rune) (ok bool, err string)
type CharacterTokenizeFunc func([]rune) []tokens.Token

type CharacterMatcher struct {
	Match    CharacterMatchFunc
	Tokenize CharacterTokenizeFunc
}

func (m CharacterMatcher) Matches(input *content.Source) Match {
	if input.Length() >= 1 {
		value := input.Take(1)
		if len(value) > 0 {
			ok, error := m.Match(value[0])
			if ok {
				match := NewMatchSuccess(m.Tokenize(value))
				input.Advance(len(value))
				return match
			} else {
				return NewMatchError(error)
			}
		}
	}

	return NewMatchError("End of input reached.")
}
