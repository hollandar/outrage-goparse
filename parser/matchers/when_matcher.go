package matchers

import (
	"fmt"

	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type WhenFunc func([]tokens.Token) bool

type WhenMatcher struct {
	Initial Matcher
	Clause  WhenFunc
	Msg     string
}

func (m WhenMatcher) Matches(input *content.Source) Match {
	initialMatch := m.Initial.Matches(input)
	if initialMatch.Success {
		var finalMatch = m.Clause(initialMatch.Tokens)
		if finalMatch {
			return initialMatch
		} else {
			return NewMatchError(fmt.Sprintf("matched content did not match the when criteria %s", m.Msg))
		}
	} else {
		return initialMatch
	}
}
