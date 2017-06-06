package lexer

import (
	"bytes"
	"io"
	"strconv"
	"unicode"
)

type Lexer struct {
	peek  byte
	words map[string]*WordToken
	R     io.ByteReader
	L     uint32
}

func NewLexer(r io.ByteReader) *Lexer {
	l := &Lexer{
		peek:  ' ',
		words: make(map[string]*WordToken),
		R:     r,
		L:     0,
	}

	l.Reverse(&WordToken{Token(TokenTag_TRUE), "true"})
	l.Reverse(&WordToken{Token(TokenTag_FALSE), "false"})

	return l
}

func (l *Lexer) Reverse(t *WordToken) {
	l.words[t.V] = t
}

func (l *Lexer) Scan() (Tokenizor, error) {
	l.skipBlankCharacters()

	if unicode.IsDigit(rune(l.peek)) {
		b := new(bytes.Buffer)
		b.WriteByte(l.peek)
		for {
			l.peek = l.readByte()
			if unicode.IsDigit(rune(l.peek)) {
				b.WriteByte(l.peek)
			} else {
				break
			}
		}
		v, err := strconv.Atoi(b.String())
		if err != nil {
			return nil, err
		}
		return &NumToken{Token(TokenTag_NUM), v}, nil
	}

	if unicode.IsLetter(rune(l.peek)) {
		b := new(bytes.Buffer)
		b.WriteByte(l.peek)
		for {
			l.peek = l.readByte()
			if r := rune(l.peek); unicode.IsLetter(r) || unicode.IsDigit(r) {
				b.WriteByte(l.peek)
			} else {
				break
			}
		}
		v := b.String()
		t, ok := l.words[v]
		if ok {
			return t, nil
		} else {
			t = &WordToken{Token(TokenTag_ID), b.String()}
			l.words[v] = t
			return t, nil
		}
	}

	if l.peek == byte(0) {
		t := Token(TokenTag_EOF)
		return &t, nil
	}

	t := Token(l.peek)
	return &t, nil
}

func (l *Lexer) skipBlankCharacters() {
	for {
		l.peek = l.readByte()
		if l.peek == ' ' || l.peek == '\t' {
			continue
		}
		if l.peek == '\n' {
			l.L += 1
			continue
		}
		break
	}
}

func (l *Lexer) readByte() byte {
	b, err := l.R.ReadByte()
	if err != nil && err == io.EOF {
		return byte(0)
	}

	if err != nil {
		panic(err)
	}

	return b
}
