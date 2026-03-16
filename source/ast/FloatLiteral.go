package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type FloatLiteral struct {
	value float64
}

func MakeFloatLiteral(val float64) *FloatLiteral {
	return &FloatLiteral{value: val}
}

func (f *FloatLiteral) Codegen(ctx *codegen.Context) llvm.Value {
	return llvm.ConstFloat(ctx.Ctx.FloatType(), f.value)
}
