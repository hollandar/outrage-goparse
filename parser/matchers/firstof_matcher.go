package matchers

import (
	"strings"

	"github.com/hollandar/outrage-goparse/parser/content"
)

type FirstOfMatcher struct {
	Matchers []Matcher
}

func (m FirstOfMatcher) Matches(input *content.Source) Match {
	var errors []string

	for _, matcher := range m.Matchers {
		match := matcher.Matches(input)
		if match.Success {
			return match
		}
		errors = append(errors, match.Error)
	}

	return NewMatchError(strings.Join(errors, " or "))
}
