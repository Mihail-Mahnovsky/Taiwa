package compiler

import (
	"bufio"
	"os"
	"os/exec"

	"mako.com/MahnoLang/source/codegen"
)

type Compiler struct {
	tb     TokensBox
	ctx    *codegen.Context
	parser Parser
}

func MakeCompiler() *Compiler {
	return &Compiler{
		tb:     MakeTokensBox(),
		ctx:    codegen.MakeContext(),
		parser: MakeParser(),
	}
}

func (c *Compiler) MakeProgramm(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lineNum int64 = 1

	for scanner.Scan() {
		lineText := scanner.Text()
		if err := MakeTokens(lineText, &c.tb, lineNum); err != nil {
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

	c.parser.Tokens = c.tb.tokens

	c.MakeLLVM()
	c.SaveIR("output.ll")
	c.Run("output.ll", "output")
}

func (c *Compiler) MakeLLVM() {
	c.parser.Tokens = c.tb.tokens

	astBox := c.parser.StatementList()

	for _, expr := range astBox.Expressions {
		expr.Codegen(c.ctx)
	}
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

	cmd := exec.Command("./" + exeFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
