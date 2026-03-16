package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type IntLiteral struct {
	value int32
}

func MakeIntLiteral(val int32) *IntLiteral {
	return &IntLiteral{value: val}
}

func (i *IntLiteral) Codegen(ctx *codegen.Context) llvm.Value {
	return llvm.ConstInt(ctx.Ctx.Int32Type(), uint64(i.value), false)
}
