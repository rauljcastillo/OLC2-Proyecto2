package src

import (
	"fmt"
	generador "main/Generador"
	"main/parser"
	"strconv"
	"strings"
)

func (l *Visitor) VisitLiteral(ctx *parser.LiteralContext, amb *Ambiente, gen *generador.Generador) Valor {
	switch {
	case ctx.INT() != nil:
		return Valor{tipo: INT, val: ctx.INT().GetText(), isTemp: false}
	case ctx.DOUBLE() != nil:
		return Valor{tipo: FLOAT, val: ctx.DOUBLE().GetText(), isTemp: false}
	case ctx.CADENA() != nil:
		tmp := gen.NewTemporal()

		gen.AddExpresion(tmp, "H", "", "")
		cadena := ctx.GetText()[1 : len(ctx.GetText())-1]
		cadena = strings.ReplaceAll(cadena, `\n`, "\n")
		for _, charac := range []byte(cadena) {
			gen.SetHeap("H", fmt.Sprint(charac))
			gen.AddExpresion("H", "H", "1", "+")
		}
		gen.SetHeap("H", "-1")
		gen.AddExpresion("H", "H", "1", "+")
		gen.Addbr()

		return Valor{tipo: STRING, val: tmp, isTemp: true}
	case ctx.BOOLEAN() != nil:
		temp, _ := strconv.ParseBool(ctx.BOOLEAN().GetText())
		truelabel := gen.NewLabel()
		falselabel := gen.NewLabel()
		if !temp {
			gen.AddGoto(falselabel)
			return Valor{tipo: BOOL, val: "0", isTemp: false, labelTrue: truelabel, labelfalse: falselabel}
		}
		return Valor{tipo: BOOL, val: "1", isTemp: false, labelTrue: truelabel, labelfalse: falselabel}

	case ctx.CHAR() != nil:

	}

	return Valor{}
}
