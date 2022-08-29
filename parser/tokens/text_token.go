package tokens

type TextToken struct {
	Token
	Value string
}

func NewTextToken(value string) (result TextToken) {
	result.Value = value
	return
}
