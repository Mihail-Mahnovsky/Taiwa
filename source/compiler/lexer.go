package compiler

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

const (
	Add = iota
	Sub
	Mul
	Div
	Num
	LParen
	RParen
	LBrace
	RBrace
	Colon
	Comma
	If
	Else
	Elif
	Break
	Continue
	Let
	Fun
	Package
	True
	False
	Id
	Assign
)

var keywords = map[string]int{
	"let":      Let,
	"package":  Package,
	"fun":      Fun,
	"if":       If,
	"else":     Else,
	"elif":     Elif,
	"break":    Break,
	"true":     True,
	"false":    False,
	"continue": Continue,
}

type Token struct {
	t int
	v string
}

type TokenInfo struct {
	t    Token
	line int64
}

func (t TokenInfo) String() string {
	return fmt.Sprintf("[Line %d] Type: %v, Lexeme: %q", t.line, t.t.t, t.t.v)
}

type TokensBox struct {
	tokens []TokenInfo
}

func MakeTokensBox() TokensBox {
	return TokensBox{
		tokens: make([]TokenInfo, 0),
	}
}

func MakeTokens(line string, tb *TokensBox, lineNum int64) error {

	if strings.TrimSpace(line) == "" {
		return nil
	}

	for i := 0; i < len(line); i++ {
		switch line[i] {
		case ' ', '\t', '\r', '\n':
			continue
		case '+':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Add, v: "+"}, line: lineNum})
		case '-':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Sub, v: "-"}, line: lineNum})
		case '*':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Mul, v: "*"}, line: lineNum})
		case '/':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Div, v: "/"}, line: lineNum})
		case '(':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: LParen, v: "("}, line: lineNum})
		case ')':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: RParen, v: ")"}, line: lineNum})
		case '{':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: LBrace, v: "{"}, line: lineNum})
		case '}':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: RBrace, v: "}"}, line: lineNum})
		case ':':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Colon, v: ":"}, line: lineNum})
		case '=':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Assign, v: "="}, line: lineNum})
		case ',':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Comma, v: ","}, line: lineNum})
		default:
			if unicode.IsDigit(rune(line[i])) {
				num := ""
				num += string(line[i])
				i++
				for i < len(line) && unicode.IsDigit(rune(line[i])) {
					num += string(line[i])
					i++
				}
				i--
				tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Num, v: num}, line: 0})
			} else if unicode.IsLetter(rune(line[i])) {
				id := ""
				id += string(line[i])
				i++
				for i < len(line) && (unicode.IsLetter(rune(line[i])) || unicode.IsDigit(rune(line[i]))) {
					id += string(line[i])
					i++
				}
				i--

				if t, ok := keywords[id]; ok {
					tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: t, v: id}, line: 0})
				} else {
					tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Id, v: id}, line: 0})
				}
			} else {
				return errors.New("Unexpected char : " + string(line[i]))
			}
		}
	}

	return nil
}
