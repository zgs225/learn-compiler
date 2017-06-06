package lexer

type TokenTag int32

const (
	TokenTag_NUM TokenTag = iota + 256
	TokenTag_ID
	TokenTag_TRUE
	TokenTag_FALSE
	TokenTag_EOF
)

type Tokenizor interface {
	Type() TokenTag
}

type Token TokenTag

func (t Token) Type() TokenTag {
	return TokenTag(t)
}

type NumToken struct {
	Token
	V int
}

type WordToken struct {
	Token
	V string
}
