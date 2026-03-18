package codegen

import "tinygo.org/x/go-llvm"

type Context struct {
	Ctx        llvm.Context
	Module     llvm.Module
	Builder    llvm.Builder
	Stack      []Variable
	ScopeStack []int
}

func MakeContext() *Context {
	ctx := llvm.NewContext()
	module := ctx.NewModule("main")
	builder := ctx.NewBuilder()

	return &Context{
		Ctx:        ctx,
		Module:     module,
		Builder:    builder,
		Stack:      []Variable{},
		ScopeStack: []int{},
	}
}

func (c *Context) PushScope() {
	c.ScopeStack = append(c.ScopeStack, len(c.Stack))
}

func (c *Context) PopScope() {
	if len(c.ScopeStack) == 0 {
		panic("pop scope called with empty stack")
	}
	start := c.ScopeStack[len(c.ScopeStack)-1]
	c.ScopeStack = c.ScopeStack[:len(c.ScopeStack)-1]
	c.Stack = c.Stack[:start]
}

func (c *Context) AddVariable(name string, val llvm.Value) (llvm.Value, bool) {
	for i := len(c.Stack) - 1; i >= 0; i-- {
		if c.Stack[i].Name == name {
			return c.Stack[i].Val, true
		}
	}

	c.Stack = append(c.Stack, Variable{Name: name, Val: val})
	return val, false
}

func (c *Context) GetVariable(name string) (llvm.Value, bool) {
	for i := len(c.Stack) - 1; i >= 0; i-- {
		if c.Stack[i].Name == name {
			return c.Stack[i].Val, true
		}
	}
	return llvm.Value{}, false
}
