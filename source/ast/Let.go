package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type Let struct {
	Name  string
	Value Expression
	Type  int
}

func MakeLet(name string, value Expression, Type int) *Let {
	return &Let{
		Name:  name,
		Value: value,
		Type:  Type,
	}
}

func (l *Let) Codegen(ctx *codegen.Context) llvm.Value {

	val := l.Value.Codegen(ctx)

	alloc := ctx.Builder.CreateAlloca(LLVMType(ctx, l.Type), l.Name)

	ctx.Builder.CreateStore(val, alloc)

	ctx.AddVariable(l.Name, alloc)

	return val
}
