package calctokens

import (
	parseTokens "github.com/hollandar/outrage-goparse/parser/tokens"
)

type BracketsToken struct {
	parseTokens.Token
	Value []parseTokens.Token
}

func (t *BracketsToken) GetTokens() []parseTokens.Token {
	return t.Value
}

func (t *BracketsToken) SetTokens(tokens []parseTokens.Token) {
	t.Value = tokens
}
