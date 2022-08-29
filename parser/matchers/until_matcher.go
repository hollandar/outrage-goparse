package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type UntilMatcher struct {
	Matcher Matcher
	Until   Matcher
}

func (m UntilMatcher) Matches(input *content.Source) Match {

	innerMatch := NewMatchEmpty()
	untilMatch := NewMatchEmpty()
	var advance int
	untilSource := input.Clone()
	for {
		if untilSource.Length() == 0 {
			return NewMatchError("Reached the end of the file matching until")
		}

		untilPosition := untilSource.Position
		untilMatch = m.Until.Matches(untilSource)
		if untilMatch.Success {
			innerSource := input.Constrain(untilPosition)
			innerMatch = m.Matcher.Matches(innerSource)

			if untilPosition != innerSource.Position {
				return NewMatchError("terminator not reached matching until")
			}
		} else {
			untilSource.Advance(1)
		}

		advance = untilSource.Position - input.Position

		if innerMatch.Success {
			break
		}
	}

	input.Advance(advance)

	var tokens []tokens.Token
	tokens = append(tokens, innerMatch.Tokens...)
	tokens = append(tokens, untilMatch.Tokens...)
	return NewMatchSuccess(tokens)
}
