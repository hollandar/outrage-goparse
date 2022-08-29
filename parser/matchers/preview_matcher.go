package matchers

import "github.com/hollandar/outrage-goparse/parser/content"

type PreviewMatcher struct {
	Preview Matcher
}

func (m PreviewMatcher) Matches(input *content.Source) Match {
	source := input.Clone()
	return m.Preview.Matches(source)
}
