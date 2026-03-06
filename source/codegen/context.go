package codegen

import "tinygo.org/x/go-llvm"

type Context struct {
	Ctx       llvm.Context
	Module    llvm.Module
	Builder   llvm.Builder
	Variables map[string]llvm.Value
}

func MakeContext() *Context {
	ctx := llvm.NewContext()
	module := ctx.NewModule("main")
	builder := ctx.NewBuilder()

	return &Context{
		Ctx:       ctx,
		Module:    module,
		Builder:   builder,
		Variables: make(map[string]llvm.Value),
	}
}
