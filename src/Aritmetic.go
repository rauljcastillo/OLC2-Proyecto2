package src

import (
	generador "main/Generador"
	"main/parser"
	"strconv"
)

func (v *Visitor) VisitAritmetic(ctx *parser.AritmeticContext, amb *Ambiente, gen *generador.Generador) interface{} {
	switch ctx.GetOp().GetText() {
	case "+":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)

		dom := TablaT(l.tipo, r.tipo)
		if dom == FLOAT {
			temp := gen.NewTemporal()
			gen.AddExpresion(temp, l.val.(string), r.val.(string), "+")
			return Valor{val: temp, tipo: dom, isTemp: true}
		} else if dom == INT {
			temp := gen.NewTemporal()
			gen.AddExpresion(temp, l.val.(string), r.val.(string), "+")
			return Valor{val: temp, tipo: dom, isTemp: true}
		} else if dom == STRING {
			tmp1 := gen.NewTemporal()
			tmp2 := gen.NewTemporal()
			tmp3 := gen.NewTemporal()
			gen.Concat()
			gen.AddCommment("Contenacion")
			gen.AddExpresion(tmp1, "P", strconv.Itoa(amb.Position), "+")
			gen.AddExpresion(tmp2, tmp1, "1", "+")
			gen.SetStack("(int)"+tmp1, l.val.(string))
			gen.SetStack("(int)"+tmp2, r.val.(string))
			gen.AddExpresion("P", "P", strconv.Itoa(amb.Position), "+")
			gen.Call("concat")
			gen.GetStack(tmp3, "(int)P")
			gen.AddExpresion("P", "P", strconv.Itoa(amb.Position), "-")
			gen.Addbr()
			return Valor{val: tmp3, tipo: dom, isTemp: true}
		}
	case "*":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)
		temp := gen.NewTemporal()
		dom := TablaT(l.tipo, r.tipo)
		if dom == FLOAT {

			gen.AddExpresion(temp, l.val.(string), r.val.(string), "*")
			return Valor{val: temp, tipo: dom, isTemp: true}
		} else if dom == INT {
			gen.AddExpresion(temp, l.val.(string), r.val.(string), "*")
			return Valor{val: temp, tipo: dom, isTemp: true}
		}
	case "-":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)
		temp := gen.NewTemporal()
		dom := TablaT(l.tipo, r.tipo)
		if dom == FLOAT {
			gen.AddExpresion(temp, l.val.(string), r.val.(string), "-")
			return Valor{val: temp, tipo: dom}
		} else if dom == INT {
			gen.AddExpresion(temp, l.val.(string), r.val.(string), "-")
			return Valor{val: temp, tipo: dom}
		}

	case "/":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)
		temp := gen.NewTemporal()
		dom := TablaT(l.tipo, r.tipo)
		if dom == FLOAT {
			gen.AddExpresion(temp, l.val.(string), r.val.(string), "/")
			return Valor{val: temp, tipo: dom, isTemp: true}
		} else if dom == INT {
			gen.AddExpresion(temp, l.val.(string), r.val.(string), "/")
			return Valor{val: temp, tipo: dom, isTemp: true}
		}
	case "<":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)

		//dom := TablaT(l.tipo, r.tipo)
		truelbl := gen.NewLabel()
		falselbl := gen.NewLabel()
		gen.Addif(l.val.(string), r.val.(string), "<", truelbl)
		gen.AddGoto(falselbl)
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: truelbl, labelfalse: falselbl}
	case ">":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)

		//dom := TablaT(l.tipo, r.tipo)

		truelbl := gen.NewLabel()
		falselbl := gen.NewLabel()
		gen.Addif(l.val.(string), r.val.(string), ">", truelbl)
		gen.AddGoto(falselbl)
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: truelbl, labelfalse: falselbl}

	case ">=":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)

		//dom := TablaT(l.tipo, r.tipo)
		truelbl := gen.NewLabel()
		falselbl := gen.NewLabel()
		gen.Addif(l.val.(string), r.val.(string), ">=", truelbl)
		gen.AddGoto(falselbl)
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: truelbl, labelfalse: falselbl}

	case "<=":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)

		//dom := TablaT(l.tipo, r.tipo)
		truelbl := gen.NewLabel()
		falselbl := gen.NewLabel()
		gen.Addif(l.val.(string), r.val.(string), "<=", truelbl)
		gen.AddGoto(falselbl)
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: truelbl, labelfalse: falselbl}

	case "==":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)

		dom := TablaT(l.tipo, r.tipo)
		if dom == STRING {
			tmp1 := gen.NewTemporal()
			tmp2 := gen.NewTemporal()
			gen.Comparar()
			gen.AddExpresion("P", "P", strconv.Itoa(amb.Position), "+")
			gen.AddExpresion(tmp1, "P", "1", "+")
			gen.SetStack("(int)"+"P", l.val.(string))
			gen.SetStack("(int)"+tmp1, r.val.(string))
			gen.Call("compare_string")
			gen.GetStack(tmp2, "(int)P")
			gen.AddExpresion("P", "P", strconv.Itoa(amb.Position), "-")
			gen.Addbr()
			truelbl := gen.NewLabel()
			falselbl := gen.NewLabel()
			gen.Addif(tmp2, "1", "==", truelbl)
			gen.AddGoto(falselbl)
			return Valor{tipo: BOOL, val: tmp2, isTemp: true, labelTrue: truelbl, labelfalse: falselbl}
		}
		truelbl := gen.NewLabel()
		falselbl := gen.NewLabel()
		gen.Addif(l.val.(string), r.val.(string), "==", truelbl)
		gen.AddGoto(falselbl)
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: truelbl, labelfalse: falselbl}

	case "!=":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)

		//dom := TablaT(l.tipo, r.tipo)
		truelbl := gen.NewLabel()
		falselbl := gen.NewLabel()
		gen.Addif(l.val.(string), r.val.(string), "!=", truelbl)
		gen.AddGoto(falselbl)
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: truelbl, labelfalse: falselbl}
	case "&&":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)
		gen.AddLabel(l.labelTrue)

		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)

		gen.Change(r.labelfalse, l.labelfalse)
		r.labelfalse = l.labelfalse
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: r.labelTrue, labelfalse: r.labelfalse}
	case "||":
		l := v.Visit(ctx.GetLeft(), amb, gen).(Valor)

		gen.AddLabel(l.labelfalse)

		r := v.Visit(ctx.GetRight(), amb, gen).(Valor)
		return Valor{tipo: BOOL, val: "", isTemp: false, labelTrue: r.labelTrue, labelfalse: r.labelfalse}
	}

	return 0
}

func (l *Visitor) VisitParent(ctx *parser.ParentContext, amb *Ambiente, gen *generador.Generador) interface{} {
	return l.Visit(ctx.Expr(), amb, gen)
}

//1+2*3
//+
/*
t0=
t1=2*3
*/
