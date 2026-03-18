package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type Arg struct {
	Name string
	Type int
}

type Function struct {
	Name       string
	ReturnType int
	Args       []Arg
	Body       Scope
}

func (f *Function) Codegen(ctx *codegen.Context) llvm.Value {
	var argTypes []llvm.Type

	for _, a := range f.Args {
		argTypes = append(argTypes, LLVMType(ctx, a.Type))
	}

	fnType := llvm.FunctionType(
		LLVMType(ctx, f.ReturnType),
		argTypes,
		false,
	)

	fn := llvm.AddFunction(ctx.Module, f.Name, fnType)

	entry := ctx.Ctx.AddBasicBlock(fn, "entry")
	ctx.Builder.SetInsertPointAtEnd(entry)

	ctx.Builder.CreateRet(f.Body.Codegen(ctx))

	return fn
}
