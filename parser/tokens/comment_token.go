package tokens

type CommentToken struct {
	Token
	Value string
}

func NewCommentToken(value string) (result CommentToken) {
	result.Value = value

	return
}
