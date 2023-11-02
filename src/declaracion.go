package src

import (
	generador "main/Generador"
	"main/parser"
	"strconv"
)

func (l *Visitor) VisitPdeclara(ctx *parser.PdeclaraContext, amb *Ambiente, gen *generador.Generador) interface{} {
	if ctx.Tipo() == nil {
		expr := l.Visit(ctx.Expr(), amb, gen).(Valor)
		if expr.tipo == BOOL {
			lblexit := gen.NewLabel()
			gen.AddLabel(expr.labelTrue)
			gen.SetStack(strconv.Itoa(amb.Position), "1")
			gen.AddGoto(lblexit)
			gen.AddLabel(expr.labelfalse)
			gen.SetStack(strconv.Itoa(amb.Position), "0")
			gen.AddLabel(lblexit)
			amb.SaveVariable(expr.tipo, ctx.ID().GetText())
			return 0
		}
		//amb.SaveVariable(expr.tipo, ctx.ID().GetText())
		gen.AddCommment("Declarando variable")
		if !expr.isTemp {
			tmp := gen.NewTemporal()
			gen.AddExpresion(tmp, expr.val.(string), "", "")
			gen.SetStack(strconv.Itoa(amb.Position), tmp)
			gen.Addbr()
			amb.SaveVariable(expr.tipo, ctx.ID().GetText())
			return 0
		}
		gen.SetStack(strconv.Itoa(amb.Position), expr.val.(string))
		gen.Addbr()
		amb.SaveVariable(expr.tipo, ctx.ID().GetText())
	} else if ctx.Expr() == nil {
		amb.SaveVariable(Num(ctx.Tipo().GetText()), ctx.ID().GetText())
	} else {
		expr := l.Visit(ctx.Expr(), amb, gen).(Valor)
		if expr.tipo != Num(ctx.Tipo().GetText()) {
			l.Errores = append(l.Errores, "No se puede asignar "+Str(expr.tipo)+" a "+ctx.Tipo().GetText()+"\n")
			return 0
		}
		gen.AddCommment("Declaracion de variable")
		if expr.tipo == BOOL {
			lblexit := gen.NewLabel()
			gen.AddLabel(expr.labelTrue)
			gen.SetStack(strconv.Itoa(amb.Position), "1")
			gen.AddGoto(lblexit)
			gen.AddLabel(expr.labelfalse)
			gen.SetStack(strconv.Itoa(amb.Position), "0")
			gen.AddLabel(lblexit)
			amb.SaveVariable(expr.tipo, ctx.ID().GetText())
			return 0
		}
		if !expr.isTemp {
			tmp := gen.NewTemporal()

			//Comentario
			gen.AddExpresion(tmp, expr.val.(string), "", "")

			gen.SetStack(strconv.Itoa(amb.Position), tmp)
			gen.Addbr()
			amb.SaveVariable(expr.tipo, ctx.ID().GetText())
			return 0
		}
		gen.SetStack(strconv.Itoa(amb.Position), expr.val.(string))
		amb.SaveVariable(expr.tipo, ctx.ID().GetText())
		gen.Addbr()
	}

	return 0
}
