package src

import (
	generador "main/Generador"
	"main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	antlr.ParseTreeVisitor
	Errores []string
}

func NewVisitor() *Visitor {
	return &Visitor{ParseTreeVisitor: &parser.BaseGramarVisitor{}}
}

func (l *Visitor) Visit(tree antlr.ParseTree, amb *Ambiente, gen *generador.Generador) interface{} {
	switch val := tree.(type) {
	case *parser.SContext:
		return l.VisitS(val, amb, gen)
	case *parser.BlockContext:
		return l.VisitBlock(val, amb, gen)
	case *parser.ProductionsContext:
		return l.VisitProduction(val, amb, gen)
	case *parser.AritmeticContext:
		return l.VisitAritmetic(val, amb, gen)
	case *parser.PrinContext:
		return l.VisitPrin(val, amb, gen)
	case *parser.LiteralContext:
		return l.VisitLiteral(val, amb, gen)
	case *parser.ParentContext:
		return l.VisitParent(val, amb, gen)
	case *parser.PifContext:
		return l.VisitPif(val, amb, gen)
	case *parser.PelseContext:
		return l.VisitPelse(val, amb, gen)
	case *parser.PdeclaraContext:
		return l.VisitPdeclara(val, amb, gen)
	case *parser.AccesoContext:
		return l.VisitAcceso(val, amb, gen)
	case *parser.SwhileContext:
		return l.VisitSwhile(val, amb, gen)
	case *parser.AsignContext:
		return l.VisitAsign(val, amb, gen)
	case *parser.PforContext:
		return l.VisitPfor(val, amb, gen)
	case *parser.PswitchContext:
		return l.VisitPswitch(val, amb, gen)
	case *parser.CcaseContext:
		return l.VisitCcase(val, amb, gen)
	case *parser.PdefaultContext:
		return l.VisitPdefault(val, amb, gen)
	case *parser.PfuncionContext:
		return l.VisitPfuncion(val, amb, gen)
	case *parser.PllamadaContext:
		return l.VisitPllamada(val, amb, gen)
	case *parser.PparamsContext:
		return l.VisitPparams(val, amb, gen)
	case *parser.ArgumentoContext:
		return l.VisitArgumento(val, amb, gen)
	case *parser.TipoArgContext:
		return l.VisitTipoArg(val, amb, gen)
	case *parser.PparametContext:
		return l.VisitPparamet(val, amb, gen)
	}

	return 0
}

func (l *Visitor) VisitS(ctx *parser.SContext, amb *Ambiente, gen *generador.Generador) interface{} {
	return l.Visit(ctx.Block(), amb, gen)
}

func (l *Visitor) VisitBlock(ctx *parser.BlockContext, amb *Ambiente, gen *generador.Generador) interface{} {
	var temp string
	for i := 0; ctx.Productions(i) != nil; i++ {
		l.Visit(ctx.Productions(i), amb, gen)
	}
	return temp
}

func (l *Visitor) VisitProduction(ctx *parser.ProductionsContext, amb *Ambiente, gen *generador.Generador) interface{} {
	if ctx.Prin() != nil {
		return l.Visit(ctx.Prin(), amb, gen)
	} else if ctx.Pif() != nil {
		return l.Visit(ctx.Pif(), amb, gen)
	} else if ctx.Pdeclara() != nil {
		return l.Visit(ctx.Pdeclara(), amb, gen)
	} else if ctx.Asign() != nil {
		return l.Visit(ctx.Asign(), amb, gen)
	} else if ctx.Swhile() != nil {
		return l.Visit(ctx.Swhile(), amb, gen)
	} else if ctx.Pfor() != nil {
		return l.Visit(ctx.Pfor(), amb, gen)
	} else if ctx.Pswitch() != nil {
		return l.Visit(ctx.Pswitch(), amb, gen)
	} else if ctx.Pfuncion() != nil {
		return l.Visit(ctx.Pfuncion(), amb, gen)
	} else if ctx.Pllamada() != nil {
		return l.Visit(ctx.Pllamada(), amb, gen)
	}
	return 0
}
