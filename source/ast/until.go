package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

const (
	TypeI32 = iota + 100
	TypeF32
	TypeString
	TypeBool
	TypeVoid
)

func LLVMType(ctx *codegen.Context, t int) llvm.Type {
	switch t {
	case TypeI32:
		return ctx.Ctx.Int32Type()
	case TypeF32:
		return ctx.Ctx.FloatType()
	case TypeBool:
		return ctx.Ctx.Int1Type()
	case TypeString:
		return llvm.PointerType(ctx.Ctx.Int8Type(), 0)
	default:
		panic("unknown type")
	}
}
