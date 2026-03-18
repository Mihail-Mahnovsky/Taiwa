package codegen

import "tinygo.org/x/go-llvm"

type Variable struct {
	Name string
	Val  llvm.Value
}
