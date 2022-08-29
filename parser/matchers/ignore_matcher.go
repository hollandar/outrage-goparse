package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type IgnoreMatcher struct {
	Ignore Matcher
}

func (m IgnoreMatcher) Matches(input *content.Source) Match {
	match := m.Ignore.Matches(input)
	if match.Success {
		return NewMatchSuccess([]tokens.Token{})
	}

	return match
}
