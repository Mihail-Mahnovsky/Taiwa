package ast

type FunctionCall struct {
	Name string
	args []Expression
}

func MakeFuctionCall(name *string, args []Expression) *FunctionCall {
	return &FunctionCall{
		Name: *name,
		args: args,
	}
}
