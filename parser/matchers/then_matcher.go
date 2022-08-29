package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type ThenMatcher struct {
	InitialMatcher Matcher
	ThenMatcher    Matcher
}

func (m ThenMatcher) Matches(input *content.Source) Match {
	trackingSource := input.Clone()
	var tokens []tokens.Token

	initialMatch := m.InitialMatcher.Matches(trackingSource)
	if initialMatch.Success {
		tokens = append(tokens, initialMatch.Tokens...)
	} else {
		return initialMatch
	}

	thenMatch := m.ThenMatcher.Matches(trackingSource)
	if thenMatch.Success {
		tokens = append(tokens, thenMatch.Tokens...)
	} else {
		return thenMatch
	}

	input.Advance(trackingSource.Position - input.Position)

	return NewMatchSuccess(tokens)
}
