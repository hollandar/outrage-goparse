package calctokens

import (
	parseTokens "github.com/hollandar/outrage-goparse/parser/tokens"
)

const (
	Function_Sqrt = iota
)

type FunctionToken struct {
	parseTokens.Token
	Function   int
	Parameters []parseTokens.Token
}

func (t *FunctionToken) GetTokens() []parseTokens.Token {
	return t.Parameters
}

func (t *FunctionToken) SetTokens(tokens []parseTokens.Token) {
	t.Parameters = tokens
}
