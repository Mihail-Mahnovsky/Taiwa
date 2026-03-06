package compiler

import (
	"bufio"
	"os"
	"os/exec"

	"mako.com/MahnoLang/source/codegen"
	"tinygo.org/x/go-llvm"
)

type Compiler struct {
	tb  TokensBox
	ctx *codegen.Context
}

func MakeCompiler() *Compiler {
	return &Compiler{
		tb:  MakeTokensBox(),
		ctx: codegen.MakeContext(),
	}
}

func (c *Compiler) MakeProgramm(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 1

	for scanner.Scan() {
		lineText := scanner.Text()
		if err := MakeTokens(lineText, &c.tb); err != nil {
			panic(err)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, tokenInfo := range c.tb.tokens {
		println(tokenInfo.String())
	}

	c.MakeLLVM()
	c.SaveIR("output.ll")
	c.Run("output.ll", "output.exe")
}

func (c *Compiler) MakeLLVM() {
	i32 := c.ctx.Ctx.Int32Type()

	fnType := llvm.FunctionType(i32, nil, false)

	fn := llvm.AddFunction(c.ctx.Module, "main", fnType)

	block := c.ctx.Ctx.AddBasicBlock(fn, "entry")

	c.ctx.Builder.SetInsertPointAtEnd(block)

	c.ctx.Builder.CreateRet(llvm.ConstInt(i32, 0, false))
}

func (c *Compiler) SaveIR(filename string) {
	err := os.WriteFile(filename, []byte(c.ctx.Module.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func (c *Compiler) BuildExe(irFile string, exeFile string) {
	cmd := exec.Command("clang", irFile, "-o", exeFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func (c *Compiler) Run(irFile string, exeFile string) {
	c.BuildExe(irFile, exeFile)

	cmd := exec.Command(exeFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
