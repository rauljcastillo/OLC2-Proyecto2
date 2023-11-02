package src

import (
	generador "main/Generador"
	"main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type CASE struct {
	val  string
	bloc antlr.ParseTree
}

func (l *Visitor) VisitPswitch(ctx *parser.PswitchContext, amb *Ambiente, gen *generador.Generador) interface{} {
	exp := l.Visit(ctx.Expr(), amb, gen).(Valor)
	exit := gen.NewLabel()
	for i := 0; ctx.Ccase(i) != nil; i++ {
		temp := l.Visit(ctx.Ccase(i), amb, gen).(CASE)
		lbl1 := gen.NewLabel()
		lbl2 := gen.NewLabel()
		gen.Addif(exp.val.(string), temp.val, "==", lbl1)
		gen.AddGoto(lbl2)
		gen.AddLabel(lbl1)
		l.Visit(temp.bloc, amb, gen)
		gen.AddGoto(exit)
		gen.AddLabel(lbl2)
	}
	l.Visit(ctx.Pdefault(), amb, gen)
	gen.AddLabel(exit)
	return 0

}

func (l *Visitor) VisitCcase(ctx *parser.CcaseContext, amb *Ambiente, gen *generador.Generador) interface{} {
	a := l.Visit(ctx.Expr(), amb, gen).(Valor)
	return CASE{val: a.val.(string), bloc: ctx.Block()}
}

func (l *Visitor) VisitPdefault(ctx *parser.PdefaultContext, amb *Ambiente, gen *generador.Generador) interface{} {
	return l.Visit(ctx.Block(), amb, gen)
}
