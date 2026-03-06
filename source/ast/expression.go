package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type Expression interface {
	Codegen(ctx *codegen.Context) llvm.Value
}
