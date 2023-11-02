package src

import (
	generador "main/Generador"
	"main/parser"
)

func (l *Visitor) VisitPif(ctx *parser.PifContext, amb *Ambiente, gen *generador.Generador) interface{} {
	condition := l.Visit(ctx.Expr(), amb, gen).(Valor)

	newAmb := NewAmbiente(amb)
	gen.AddLabel(condition.labelTrue)
	l.Visit(ctx.Block(), newAmb, gen)
	if ctx.Pelse() != nil {
		lblexit := gen.NewLabel()
		gen.AddGoto(lblexit)
		gen.AddLabel(condition.labelfalse)
		l.Visit(ctx.Pelse(), amb, gen)
		gen.AddLabel(lblexit)
		return 0
	}
	gen.AddLabel(condition.labelfalse)
	return 0
}

func (l *Visitor) VisitPelse(ctx *parser.PelseContext, amb *Ambiente, gen *generador.Generador) interface{} {
	if ctx.Pif() != nil {
		return l.Visit(ctx.Pif(), amb, gen)
	} else {
		newAmb := NewAmbiente(amb)
		l.Visit(ctx.Block(), newAmb, gen)
		return 0
	}
}
