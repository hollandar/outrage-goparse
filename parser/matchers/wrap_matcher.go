package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type WrapFunc func(match Match) ([]tokens.Token, error)

type WrapMatcher struct {
	Inner Matcher
	Wrap  WrapFunc
}

func (m WrapMatcher) Matches(input *content.Source) Match {
	match := m.Inner.Matches(input)

	if !match.Success {
		return match
	} else {
		match, err := m.Wrap(match)
		if err != nil {
			return NewMatchError(err.Error())
		}
		return NewMatchSuccess(match)
	}
}
