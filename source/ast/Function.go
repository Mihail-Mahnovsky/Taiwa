package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type Function struct {
	Name       string
	ReturnType llvm.Type
	Args       llvm.Type
	Body       Scope
}

func (f *Function) Codegen(ctx *codegen.Context) llvm.Value {
	fnType := llvm.FunctionType(f.ReturnType, nil, false)
	fn := llvm.AddFunction(ctx.Module, f.Name, fnType)
	entry := ctx.Ctx.AddBasicBlock(fn, "entry")
	ctx.Builder.SetInsertPointAtEnd(entry)

	ctx.Builder.CreateRet(f.Body.Codegen(ctx))
	return fn
}
