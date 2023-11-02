package main

import (
	"fmt"
	generador "main/Generador"
	"main/parser"
	"main/src"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	texto := `
	for j in 1...3{
		if j==2{
			break
		}
	}
	` //2>1 || (3>1 && 3<10)
	is := antlr.NewInputStream(texto)
	lexer := parser.NewGramarLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGramarParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()
	visitor := src.NewVisitor()
	gen := generador.NewGenerator()
	global := src.Global()
	visitor.Visit(tree, global, gen)
	if len(visitor.Errores) > 0 {
		fmt.Println(strings.Join(visitor.Errores, ""))
	} else {
		fmt.Println(gen.GetCode())
	}
}
