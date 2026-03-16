package compiler

import (
	"strconv"

	"mako.com/MahnoLang/source/ast"
	"tinygo.org/x/go-llvm"
)

type AstBox struct {
	Expressions []ast.Expression
}

type Parser struct {
	Pos    int
	Tokens []TokenInfo
}

func MakeParser() Parser {
	return Parser{
		Pos: 0,
	}
}

func (p *Parser) current() TokenInfo {
	return p.Tokens[p.Pos]
}

func (p *Parser) eat(wType int) {
	if p.current().t.t == wType && p.Pos < len(p.Tokens) {
		p.Pos += 1
	} else {
		panic("err in eat fun")
	}
}

func (p *Parser) next() int {
	if p.Pos < len(p.Tokens) {
		return p.Tokens[p.Pos+1].t.t
	}
	panic("call next")
}

func (p *Parser) StatementList() AstBox {
	var exprs AstBox

	for p.Pos < len(p.Tokens) && p.current().t.t != RBrace {
		exprs.Expressions = append(exprs.Expressions, p.statement())
	}

	return exprs
}

func (p *Parser) statement() ast.Expression {
	switch p.current().t.t {
	case Fun:
		p.eat(Fun)
		funName := p.current().t.v
		p.eat(Id)
		return p.parseFunction(&funName)
	case Num:
		return p.expression()
	default:
		panic("don`t undestart the token")
	}
}

func (p *Parser) expression() ast.Expression {
	return p.AddSub()
}

func (p *Parser) AddSub() ast.Expression {
	left := p.MulDiv()

	for p.Pos < len(p.Tokens) && (p.current().t.t == Add || p.current().t.t == Sub) {
		op := p.current().t.t
		p.eat(op)

		right := p.MulDiv()

		if op == Add {
			left = ast.MakeBinOp(left, right, '+')
		} else {
			left = ast.MakeBinOp(left, right, '-')
		}
	}

	return left
}
func (p *Parser) MulDiv() ast.Expression {
	left := p.Factor()

	for p.Pos < len(p.Tokens) && (p.current().t.t == Mul || p.current().t.t == Div) {
		op := p.current().t.t
		p.eat(op)

		right := p.Factor()

		if op == Mul {
			left = ast.MakeBinOp(left, right, '*')
		} else {
			left = ast.MakeBinOp(left, right, '/')
		}
	}

	return left
}

func (p *Parser) Factor() ast.Expression {
	switch p.current().t.t {
	case Num:
		val, _ := strconv.ParseInt(p.current().t.v, 10, 32)
		p.eat(Num)
		return ast.MakeIntLiteral(val)
	case LParen:
		p.eat(LParen)
		expr := p.expression()
		p.eat(RParen)
		return expr
	default:
		panic("unexpected token in factor")
	}
}

func (p *Parser) parseFunction(name *string) ast.Expression {
	p.eat(LParen)
	p.eat(RParen)

	p.eat(LBrace)
	statements := p.StatementList()
	p.eat(RBrace)

	//if p.next() == Colon {
	//	p.eat(Colon)

	//}

	scope := ast.MakeScope(statements.Expressions)
	return &ast.Function{
		Name:       *name,
		ReturnType: llvm.GlobalContext().Int32Type(),
		Body:       scope,
	}
}
