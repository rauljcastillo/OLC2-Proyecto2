package src

import (
	generador "main/Generador"
	"main/parser"
	"strconv"
)

func (l *Visitor) VisitAsign(ctx *parser.AsignContext, amb *Ambiente, gen *generador.Generador) interface{} {
	variab := amb.getVar(ctx.ID().GetText()).(*Variable)
	expr := l.Visit(ctx.Expr(), amb, gen).(Valor)
	if expr.tipo == INT && variab.tipo == FLOAT {
		if !expr.isTemp {
			tmp := gen.NewTemporal()
			gen.AddExpresion(tmp, expr.val.(string), "", "")
			gen.SetStack(strconv.Itoa(variab.posicion), tmp)
			return 0
		}

		gen.SetStack(strconv.Itoa(variab.posicion), expr.val.(string))
	} else if expr.tipo == variab.tipo {
		if !expr.isTemp {
			tmp := gen.NewTemporal()
			gen.AddExpresion(tmp, expr.val.(string), "", "")
			gen.SetStack(strconv.Itoa(variab.posicion), tmp)
			return 0
		}
		gen.SetStack(strconv.Itoa(variab.posicion), expr.val.(string))
	} else {
		l.Errores = append(l.Errores, "No se puede asignar "+Str(expr.tipo)+" con "+Str(variab.tipo))

	}
	return 0
}
