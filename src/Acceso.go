package src

import (
	generador "main/Generador"
	"main/parser"
	"strconv"
)

func (l *Visitor) VisitAcceso(ctx *parser.AccesoContext, amb *Ambiente, gen *generador.Generador) interface{} {
	variab := amb.getVar(ctx.ID().GetText()).(*Variable)
	tmp := gen.NewTemporal()
	gen.AddCommment("Acceso a variable")
	gen.GetStack(tmp, strconv.Itoa(variab.posicion))
	gen.Addbr()
	if variab.tipo == BOOL {
		lbltrue := gen.NewLabel()
		lblfalse := gen.NewLabel()
		gen.Addif(tmp, "1", "==", lbltrue)
		gen.AddGoto(lblfalse)
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: lbltrue, labelfalse: lblfalse}
	}
	return Valor{tipo: variab.tipo, val: tmp, isTemp: true, position: variab.posicion}
}
