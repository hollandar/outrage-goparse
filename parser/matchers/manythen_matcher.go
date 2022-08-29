package matchers

import (
	"fmt"

	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type ManyThenMatcher struct {
	Many           Matcher
	Then           Matcher
	MinimumMatches int
}

func (m ManyThenMatcher) Matches(input *content.Source) Match {
	var matches int
	matches = 0
	var tokens []tokens.Token
	for {
		termMatch := m.Many.Matches(input)
		if termMatch.Success {
			tokens = append(tokens, termMatch.Tokens...)
			matches++
		}

		thenMatch := m.Then.Matches(input)
		if thenMatch.Success {
			tokens = append(tokens, thenMatch.Tokens...)
			break
		}

		if !termMatch.Success && !thenMatch.Success {
			return NewMatchError(fmt.Sprintf("expected at least %s or %s", termMatch.Error, thenMatch.Error))
		}
	}

	if matches <= m.MinimumMatches {
		return NewMatchError(fmt.Sprintf("expected at least %d, found %d", m.MinimumMatches, matches))
	}

	return NewMatchSuccess(tokens)
}
