package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
)

type BlockMatcher struct {
	InnerMatcher Matcher
}

func (m BlockMatcher) Matches(input *content.Source) Match {
	return m.InnerMatcher.Matches(input)
}
