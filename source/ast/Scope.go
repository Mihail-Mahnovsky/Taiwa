package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type Scope struct {
	exps []Expression
}

func MakeScope(exps []Expression) Scope {
	return Scope{
		exps: exps,
	}
}

func (s *Scope) Codegen(ctx *codegen.Context) llvm.Value {
	var last llvm.Value
	for _, expr := range s.exps {
		last = expr.Codegen(ctx)
	}

	if last.IsNil() {
		last = llvm.ConstInt(ctx.Ctx.Int32Type(), 0, false)
	}
	return last
}
