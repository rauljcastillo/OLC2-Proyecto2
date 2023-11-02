// Code generated from Gramar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Gramar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GramarParser.
type GramarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GramarParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by GramarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GramarParser#productions.
	VisitProductions(ctx *ProductionsContext) interface{}

	// Visit a parse tree produced by GramarParser#pbrk.
	VisitPbrk(ctx *PbrkContext) interface{}

	// Visit a parse tree produced by GramarParser#pswitch.
	VisitPswitch(ctx *PswitchContext) interface{}

	// Visit a parse tree produced by GramarParser#ccase.
	VisitCcase(ctx *CcaseContext) interface{}

	// Visit a parse tree produced by GramarParser#pdefault.
	VisitPdefault(ctx *PdefaultContext) interface{}

	// Visit a parse tree produced by GramarParser#pfor.
	VisitPfor(ctx *PforContext) interface{}

	// Visit a parse tree produced by GramarParser#prin.
	VisitPrin(ctx *PrinContext) interface{}

	// Visit a parse tree produced by GramarParser#cexpr.
	VisitCexpr(ctx *CexprContext) interface{}

	// Visit a parse tree produced by GramarParser#swhile.
	VisitSwhile(ctx *SwhileContext) interface{}

	// Visit a parse tree produced by GramarParser#asign.
	VisitAsign(ctx *AsignContext) interface{}

	// Visit a parse tree produced by GramarParser#pdeclara.
	VisitPdeclara(ctx *PdeclaraContext) interface{}

	// Visit a parse tree produced by GramarParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by GramarParser#pif.
	VisitPif(ctx *PifContext) interface{}

	// Visit a parse tree produced by GramarParser#pelse.
	VisitPelse(ctx *PelseContext) interface{}

	// Visit a parse tree produced by GramarParser#pfuncion.
	VisitPfuncion(ctx *PfuncionContext) interface{}

	// Visit a parse tree produced by GramarParser#pparams.
	VisitPparams(ctx *PparamsContext) interface{}

	// Visit a parse tree produced by GramarParser#pparamet.
	VisitPparamet(ctx *PparametContext) interface{}

	// Visit a parse tree produced by GramarParser#pllamada.
	VisitPllamada(ctx *PllamadaContext) interface{}

	// Visit a parse tree produced by GramarParser#argumento.
	VisitArgumento(ctx *ArgumentoContext) interface{}

	// Visit a parse tree produced by GramarParser#tipoArg.
	VisitTipoArg(ctx *TipoArgContext) interface{}

	// Visit a parse tree produced by GramarParser#parent.
	VisitParent(ctx *ParentContext) interface{}

	// Visit a parse tree produced by GramarParser#Literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by GramarParser#Acceso.
	VisitAcceso(ctx *AccesoContext) interface{}

	// Visit a parse tree produced by GramarParser#aritmetic.
	VisitAritmetic(ctx *AritmeticContext) interface{}
}
