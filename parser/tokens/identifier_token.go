package tokens

type IdentitifierToken struct {
	Token
	Value string

	innerTokens []Token
}

func NewIdentifierToken(value string) (result IdentitifierToken) {
	result.Value = value
	return
}

func (t *IdentitifierToken) SetTokens(tokens []Token) {
	t.innerTokens = tokens
}

func (t *IdentitifierToken) GetTokens() []Token {
	return t.innerTokens
}
