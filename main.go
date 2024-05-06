package main

import (
	"os"
)

func main() {
	compiler := NewCompiler(os.Args[1])
	GetApplication().SetCommandBuffer(compiler.Compile())
	GetApplication().Run()
}
