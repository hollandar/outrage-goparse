package matchers

import "github.com/hollandar/outrage-goparse/parser/content"

type RefFunc func() Matcher

type RefMatcher struct {
	Ref RefFunc
}

func (m RefMatcher) Matches(input *content.Source) Match {
	return m.Ref().Matches(input)
}
