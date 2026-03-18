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

	retType := LLVMType(ctx, f.ReturnType)

	fnType := llvm.FunctionType(
		retType,
		argTypes,
		false,
	)

	fn := llvm.AddFunction(ctx.Module, f.Name, fnType)

	entry := ctx.Ctx.AddBasicBlock(fn, "entry")
	ctx.Builder.SetInsertPointAtEnd(entry)

	val := f.Body.Codegen(ctx)

	if val.Type() != retType {
		switch f.ReturnType {
		case TypeF32:
			val = ctx.Builder.CreateSIToFP(val, retType, "")
		case TypeI32:
			val = ctx.Builder.CreateFPToSI(val, retType, "")
		}
	}

	ctx.Builder.CreateRet(val)

	return fn
}
