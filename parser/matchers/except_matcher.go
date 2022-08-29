package matchers

import "github.com/hollandar/outrage-goparse/parser/content"

type ExceptMatcher struct {
	Except Matcher
	Actual Matcher
}

func (m ExceptMatcher) Matches(input *content.Source) Match {
	matcher := PreviewMatcher{
		Preview: m.Except,
	}

	match := matcher.Matches(input)
	if match.Success {
		return NewMatchError("except match was found")
	} else {
		return m.Actual.Matches(input)
	}

}
