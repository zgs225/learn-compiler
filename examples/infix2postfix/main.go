package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 将9+5-2这样的中缀表达式翻译后后缀表达式

type parser struct {
	L byte
	R io.ByteReader
}

func (p *parser) Expr() {
	p.L = p.readByte()
	p.Term()
	for {
		if p.L == '+' {
			p.Match('+')
			p.Term()
			fmt.Print("+")
			continue
		} else if p.L == '-' {
			p.Match('-')
			p.Term()
			fmt.Print("-")
			continue
		} else {
			break
		}
	}
}

func (p *parser) Term() {
	if p.L >= '0' && p.L <= '9' {
		fmt.Printf("%c", p.L)
		p.Match(p.L)
	} else {
		panic("Syntax error")
	}
}

func (p *parser) Match(b byte) {
	if p.L == b {
		p.L = p.readByte()
	} else {
		panic(fmt.Sprintf("Syntax error: expect %v got %v\n", p.L, b))
	}
}

func (p *parser) readByte() byte {
	b, err := p.R.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	p := &parser{
		R: bytes.NewReader(b),
	}
	p.Expr()
	fmt.Print("\n")
}
