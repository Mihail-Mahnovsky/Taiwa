package ast

import (
	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type BinaryOperation struct {
	left      Expression
	right     Expression
	operation byte
}

func MakeBinOp(left, right Expression, op byte) *BinaryOperation {
	return &BinaryOperation{
		left:      left,
		right:     right,
		operation: op,
	}
}

func (b *BinaryOperation) Codegen(ctx *codegen.Context) llvm.Value {
	leftVal := b.left.Codegen(ctx)
	rightVal := b.right.Codegen(ctx)

	switch b.operation {
	case '+':
		return ctx.Builder.CreateAdd(leftVal, rightVal, "addtmp")
	case '-':
		return ctx.Builder.CreateSub(leftVal, rightVal, "subtmp")
	case '*':
		return ctx.Builder.CreateMul(leftVal, rightVal, "multmp")
	case '/':
		return ctx.Builder.CreateSDiv(leftVal, rightVal, "divtmp")
	default:
		panic("don`t support op in bin op")
	}
}
