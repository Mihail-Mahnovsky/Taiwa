package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type Assign struct {
	Name  string
	Right Expression
}

func MakeAssign(name string, value Expression) *Assign {
	return &Assign{
		Name:  name,
		Right: value,
	}
}

func (a *Assign) Codegen(ctx *codegen.Context) llvm.Value {
	val := a.Right.Codegen(ctx)

	alloc, ok := ctx.GetVariable(a.Name)
	if !ok {
		panic("undefined variable: " + a.Name)
	}

	ctx.Builder.CreateStore(val, alloc)

	return val
}
