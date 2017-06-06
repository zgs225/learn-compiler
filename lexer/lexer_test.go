package lexer

import (
	"bytes"
	"testing"
)

func TestLexer(t *testing.T) {
	s := "3 + 429 - i * 9"
	r := bytes.NewReader([]byte(s))
	l := NewLexer(r)

	token, _ := l.Scan()
	if token.Type() != TokenTag_NUM {
		t.Error("Lexer scan tag error")
	} else {
		v := token.(*NumToken)
		if v.V != 3 {
			t.Error("Lexer scan value error")
		}
	}

	token, _ = l.Scan()
	if byte(token.Type()) != '+' {
		t.Error("Lexer scan tag error")
	}

	token, _ = l.Scan()
	if token.Type() != TokenTag_NUM {
		t.Error("Lexer scan tag error")
	} else {
		v := token.(*NumToken)
		if v.V != 429 {
			t.Error("Lexer scan value error")
		}
	}

	token, _ = l.Scan()
	if byte(token.Type()) != '-' {
		t.Error("Lexer scan tag error")
	}

	token, _ = l.Scan()
	if token.Type() != TokenTag_ID {
		t.Error("Lexer scan tag error")
	} else {
		v := token.(*WordToken)
		if v.V != "i" {
			t.Error("Lexer scan value error")
		}
	}

	token, _ = l.Scan()
	if byte(token.Type()) != '*' {
		t.Error("Lexer scan tag error")
	}

	token, _ = l.Scan()
	if token.Type() != TokenTag_NUM {
		t.Error("Lexer scan tag error")
	} else {
		v := token.(*NumToken)
		if v.V != 9 {
			t.Error("Lexer scan value error")
		}
	}

	token, _ = l.Scan()
	if token.Type() != TokenTag_EOF {
		t.Error("Lexer scan tag error")
	}
}
