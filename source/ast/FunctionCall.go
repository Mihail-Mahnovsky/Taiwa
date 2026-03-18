package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type FunctionCall struct {
	Name string
	args []Expression
}

func MakeFuctionCall(name *string, args []Expression) *FunctionCall {
	return &FunctionCall{
		Name: *name,
		args: args,
	}
}

func (i *FunctionCall) Codegen(ctx *codegen.Context) llvm.Value {

	return llvm.Value{}
}
