package main

import "mako.com/MahnoLang/source/compiler"

func main() {
	compiller := compiler.MakeCompiler()
	compiller.MakeProgramm("examples/test.nn")
}
