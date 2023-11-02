package src

import (
	"fmt"
	generador "main/Generador"
	"main/parser"
	"strconv"
)

type Argument struct {
	name     string
	valor    any
	tipo     int
	position int
}

func (l *Visitor) VisitPllamada(ctx *parser.PllamadaContext, amb *Ambiente, gen *generador.Generador) interface{} {
	temp := amb.getFunction(ctx.ID().GetText())

	if temp == "" {
		l.Errores = append(l.Errores, "No existe la funcion "+ctx.ID().GetText())
		return 0
	}
	if temp.(Funcion).Params == nil {
		if !gen.Contiene("void " + temp.(Funcion).ID + "(){") {
			nuevoamb := NewAmbiente(amb)
			index := gen.LenCode()
			gen.AddFunction("void " + temp.(Funcion).ID + "(){")
			l.Visit(temp.(Funcion).Block.(parser.IBlockContext), nuevoamb, gen)
			gen.MovFunction(index)
		}

		gen.AddExpresion("P", "P", strconv.Itoa(amb.Position), "+")
		gen.Call(ctx.ID().GetText())
		gen.AddExpresion("P", "P", strconv.Itoa(amb.Position), "-")

	} else {
		listArg := l.Visit(ctx.Argumento(), amb, gen)
		//newAmbiente := NewAmbiente(amb)
		fmt.Println(listArg)
		if !gen.Contiene("void " + temp.(Funcion).ID + "(){") {
			nuevoamb := NewAmbiente(amb)
			index := gen.LenCode()
			gen.AddFunction("void " + temp.(Funcion).ID + "(){")
			l.Visit(temp.(Funcion).Block.(parser.IBlockContext), nuevoamb, gen)
			gen.MovFunction(index)
		}

		//newAmbiente := NewAmbiente(amb)
	}
	return 0
}

func (v *Visitor) VisitArgumento(ctx *parser.ArgumentoContext, amb *Ambiente, gen *generador.Generador) interface{} {
	v.Visit(ctx.TipoArg(), amb, gen)
	v.Visit(ctx.Argumento(), amb, gen)
	return 0

}

func (l *Visitor) VisitTipoArg(ctx *parser.TipoArgContext, amb *Ambiente, gen *generador.Generador) interface{} {
	if ctx.ID() != nil {
		exp := l.Visit(ctx.GetExp2(), amb, gen).(Valor)
		return Argument{name: ctx.ID().GetText(), valor: exp.val, tipo: exp.tipo}
	}

	if ctx.PUNTE() != nil {
		exp := l.Visit(ctx.GetExp1(), amb, gen).(Valor)
		return Argument{name: "", valor: exp.val, tipo: exp.tipo, position: exp.position}
	}

	exp := l.Visit(ctx.GetExp3(), amb, gen).(Valor)
	tmp1 := gen.NewTemporal()
	return Valor{tipo: exp.tipo, val: tmp1, isTemp: false}
}

/*
	t0=P+1;
	stack[(int)t0]=3

	t0=t0+1;
	stack[(int)t0]=3
*/
