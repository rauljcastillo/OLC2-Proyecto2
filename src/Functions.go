package src

import (
	generador "main/Generador"
	"main/parser"
)

type Param struct {
	extern  string
	intern  string
	tipo    int
	isArray bool
	inout   bool
}

func (l *Visitor) VisitPfuncion(ctx *parser.PfuncionContext, amb *Ambiente, gen *generador.Generador) interface{} {
	if ctx.Tipo() != nil {
		tipo := Num(ctx.Tipo().GetText())
		if ctx.Pparams() != nil {
			a := l.Visit(ctx.Pparams(), amb, gen).([]Param)
			if !amb.saveFunctions(tipo, ctx.ID().GetText(), a, ctx.Block()) {
				l.Errores = append(l.Errores, "Ya existe una funcion con el mismo nombre")
			}

		} else {
			if !amb.saveFunctions(tipo, ctx.ID().GetText(), nil, ctx.Block()) {
				l.Errores = append(l.Errores, "Ya existe una funcion con el mismo nombre")
			}
		}

	} else {
		if ctx.Pparams() != nil {
			a := l.Visit(ctx.Pparams(), amb, gen).([]Param)
			if !amb.saveFunctions(VOID, ctx.ID().GetText(), a, ctx.Block()) {
				l.Errores = append(l.Errores, "Ya existe una funcion con el mismo nombre")
			}
		} else {
			if !amb.saveFunctions(VOID, ctx.ID().GetText(), nil, ctx.Block()) {
				l.Errores = append(l.Errores, "Ya existe una funcion con el mismo nombre")
			}
		}
	}

	return 0
}

func (l *Visitor) VisitPparams(ctx *parser.PparamsContext, amb *Ambiente, gen *generador.Generador) interface{} {
	var temp []Param
	for _, elem := range ctx.AllPparamet() {
		a := l.Visit(elem, amb, gen).(Param)
		temp = append(temp, a)
	}
	return temp
}

func (l *Visitor) VisitPparamet(ctx *parser.PparametContext, amb *Ambiente, gen *generador.Generador) interface{} {
	if ctx.ABR() != nil {
		if ctx.GetPid().GetText() != "_" {
			if ctx.INOUT() != nil {
				return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: true, inout: true}
			} else {
				return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: true, inout: false}
			}

		}
		if ctx.INOUT() != nil {
			return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: true, inout: true}
		} else {
			return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: true, inout: false}
		}

	}
	if ctx.GetPid().GetText() != "_" {
		if ctx.INOUT() != nil {
			return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(1).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: false, inout: true}
		} else {
			return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(1).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: false, inout: false}
		}

	}
	if ctx.INOUT() != nil {
		return Param{extern: "_", intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: false, inout: true}
	} else {
		return Param{extern: "_", intern: ctx.GetPid().GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: false, inout: false}
	}

}
