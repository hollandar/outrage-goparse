package matchers

import (
	"fmt"

	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type ManyMatcher struct {
	Many           Matcher
	MinimumMatches int
}

func (m ManyMatcher) Matches(input *content.Source) Match {
	var matches int = 0
	var tokens []tokens.Token
	for {
		termMatch := m.Many.Matches(input)
		if termMatch.Success {
			tokens = append(tokens, termMatch.Tokens...)
			matches++
		} else {
			break
		}
	}

	if matches < m.MinimumMatches {
		return NewMatchError(fmt.Sprintf("expected at least %d, found %d", m.MinimumMatches, matches))
	}

	return NewMatchSuccess(tokens)
}
