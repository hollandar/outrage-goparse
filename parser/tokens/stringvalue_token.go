package tokens

type StringValueToken struct {
	Token
	Value []rune
}

func NewStringValueToken(slice []rune) (result StringValueToken) {
	result.Value = slice
	return
}
