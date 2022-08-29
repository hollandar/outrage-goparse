package tokens

type Wrapper interface {
	GetTokens() []Token
	SetTokens(tokens []Token)
}
