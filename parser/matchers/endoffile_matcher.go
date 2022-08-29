package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type EndOfFileMatcher struct {
}

func (m EndOfFileMatcher) Matches(input *content.Source) Match {
	if input.Length() == 0 {
		return NewMatchSuccess([]tokens.Token{tokens.NewEndOfFileToken()})
	}

	return NewMatchError("End of file expected.")
}
