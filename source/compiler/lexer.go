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
	Сolon
	If
	Else
	Elif
	Break
	Continue
	Let
	Fun
	Id
	Assign
)

var keywords = map[string]int{
	"let":      Let,
	"fun":      Fun,
	"if":       If,
	"else":     Else,
	"elif":     Elif,
	"break":    Break,
	"continue": Continue,
}

type Token struct {
	t int
	v string
}

type TokenInfo struct {
	t    Token
	line int
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

func MakeTokens(line string, tb *TokensBox) error {

	if strings.TrimSpace(line) == "" {
		return nil
	}

	for i := 0; i < len(line); i++ {
		switch line[i] {
		case ' ', '\t', '\r', '\n':
			continue
		case '+':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Add, v: "+"}, line: 0})
		case '-':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Sub, v: "-"}, line: 0})
		case '*':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Mul, v: "*"}, line: 0})
		case '/':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Div, v: "/"}, line: 0})
		case '(':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: LParen, v: "("}, line: 0})
		case ')':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: RParen, v: ")"}, line: 0})
		case '{':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: LBrace, v: "{"}, line: 0})
		case '}':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: RBrace, v: "}"}, line: 0})
		case ':':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Сolon, v: ":"}, line: 0})
		case '=':
			tb.tokens = append(tb.tokens, TokenInfo{t: Token{t: Assign, v: "="}, line: 0})
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
