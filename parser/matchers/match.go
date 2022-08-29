package matchers

import (
	"github.com/hollandar/outrage-goparse/parser/tokens"
)

type Match struct {
	Success bool
	Tokens  []tokens.Token
	Error   string
}

func NewMatchSuccess(tokens []tokens.Token) (result Match) {
	result.Success = true
	result.Tokens = tokens

	return
}

func NewMatchError(error string) (result Match) {
	result.Success = false
	result.Error = error

	return
}

func NewMatchEmpty() (result Match) {
	result.Success = false
	result.Tokens = make([]tokens.Token, 0)

	return
}
