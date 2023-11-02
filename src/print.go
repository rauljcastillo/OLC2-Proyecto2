package src

import (
	generador "main/Generador"
	"main/parser"
	"strconv"
)

func (l *Visitor) VisitPrin(ctx *parser.PrinContext, amb *Ambiente, gen *generador.Generador) interface{} {
	if ctx.Cexpr() != nil {
		l.Visit(ctx.Cexpr(), amb, gen)
	}
	value := l.Visit(ctx.Expr(), amb, gen).(Valor)
	if value.tipo == INT {
		gen.Printf("d", "(int)"+value.val.(string))
		gen.Printf("c", "10")
	} else if value.tipo == FLOAT {
		gen.Printf("f", value.val.(string))
	} else if value.tipo == CHAR {
		gen.Printf("c", "(int)"+value.val.(string))
		gen.Printf("c", "10")
	} else if value.tipo == STRING {
		gen.AddCommment("Impresión")
		gen.AddExpresion("P", "P", strconv.Itoa(amb.Position), "+")
		gen.SetStack("(int)P", value.val.(string))
		gen.Prints()
		gen.Call("prints")
		gen.AddExpresion("P", "P", strconv.Itoa(amb.Position), "-")

	} else if value.tipo == BOOL {
		lblexit := gen.NewLabel()
		gen.AddLabel(value.labelTrue)
		gen.Printf("c", "116")
		gen.Printf("c", "114")
		gen.Printf("c", "117")
		gen.Printf("c", "101")
		gen.Printf("c", "10")
		gen.AddGoto(lblexit)

		gen.AddLabel(value.labelfalse)

		gen.Printf("c", "102")
		gen.Printf("c", "97")
		gen.Printf("c", "108")
		gen.Printf("c", "115")
		gen.Printf("c", "101")
		gen.Printf("c", "10")
		gen.AddLabel(lblexit)
		return 0

	}
	return 0
}
