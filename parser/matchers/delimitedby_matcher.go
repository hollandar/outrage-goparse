package matchers

import (
	"fmt"

	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type DelimitedByMatcher struct {
	Term          Matcher
	Delimiter     Matcher
	MinOccurrence int
	MaxOccurrence int
}

func (m DelimitedByMatcher) Matches(input *content.Source) Match {
	var occurrences int
	occurrences = 0
	var tokens []tokens.Token
	for {

		termMatch := m.Term.Matches(input)
		if termMatch.Success {
			tokens = append(tokens, termMatch.Tokens...)
			occurrences++
		} else {
			return termMatch
		}

		delimiterMatch := m.Delimiter.Matches(input)
		if delimiterMatch.Success {
			tokens = append(tokens, delimiterMatch.Tokens...)
			continue
		} else {
			break
		}

	}

	if occurrences <= m.MinOccurrence || occurrences >= m.MaxOccurrence {
		return NewMatchError(fmt.Sprintf("expected between %d and %d occurrences, found %d", m.MinOccurrence, m.MaxOccurrence, occurrences))
	}

	return NewMatchSuccess(tokens)
}
