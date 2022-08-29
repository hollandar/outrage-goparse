package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/content"
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type CastFunc func(from tokens.Token) tokens.Token

type CastMatcher struct {
	InnerMatcher Matcher
	Cast         CastFunc
}

func (m CastMatcher) Matches(input *content.Source) Match {
	match := m.InnerMatcher.Matches(input)
	if match.Success {
		tokens := []tokens.Token{}
		for _, token := range match.Tokens {
			newToken := m.Cast(token)
			tokens = append(tokens, newToken)
		}

		return NewMatchSuccess(tokens)
	}

	return match
}
