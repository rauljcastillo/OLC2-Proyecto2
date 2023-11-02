package src

import (
	generador "main/Generador"
	"main/parser"
	"strconv"
)

func (l *Visitor) VisitPfor(ctx *parser.PforContext, amb *Ambiente, gen *generador.Generador) interface{} {
	if ctx.POINTS() != nil {
		exp1 := l.Visit(ctx.Expr(0), amb, gen).(Valor)
		exp2 := l.Visit(ctx.Expr(1), amb, gen).(Valor)

		nuevo := NewAmbiente(amb)
		gen.SetStack(strconv.Itoa(nuevo.Position), exp1.val.(string))
		nuevo.SaveVariable(INT, ctx.ID().GetText())

		loop := gen.NewLabel()
		truelbl := gen.NewLabel()
		falselbl := gen.NewLabel()

		gen.AddLabel(loop)
		tmp := gen.NewTemporal()
		gen.GetStack(tmp, strconv.Itoa(nuevo.Position-1))
		gen.Addif(tmp, exp2.val.(string), "<=", truelbl)
		gen.AddGoto(falselbl)
		gen.AddLabel(truelbl)
		tran := l.Visit(ctx.Block(), nuevo, gen)
		gen.AddExpresion(tmp, tmp, "1", "+")
		gen.SetStack(strconv.Itoa(nuevo.Position-1), tmp)
		gen.AddGoto(loop)
		gen.AddLabel(falselbl)
		if tran.(string) != "" {
			gen.AddLabel(tran.(string))
		}
		return 0
	}
	expr := l.Visit(ctx.Expr(0), amb, gen).(Valor)

	if expr.tipo == STRING {
		nuevo := NewAmbiente(amb)
		nuevo.SaveVariable(CHAR, ctx.ID().GetText())

		loop := gen.NewLabel()
		truelbl := gen.NewLabel()
		falselbl := gen.NewLabel()

		gen.AddLabel(loop)
		tmp := gen.NewTemporal()
		gen.GetHeap(tmp, expr.val.(string))
		gen.SetStack(strconv.Itoa(nuevo.Position-1), tmp)
		gen.Addif(tmp, "-1", "!=", truelbl)
		gen.AddGoto(falselbl)
		gen.AddLabel(truelbl)
		l.Visit(ctx.Block(), nuevo, gen)
		gen.AddExpresion(expr.val.(string), expr.val.(string), "1", "+")
		gen.AddGoto(loop)
		gen.AddLabel(falselbl)

	}
	return 0
}

/*
	t0=H
	heap[H]=117;
	H=H+1;
	heap[H]=-1;
	H=H+1;


	stack[1]=t0

	//Acceso
	t1=stack[1]
	L0:
		t2=heap[(int)t1]
		stack[2]=t2
		if(t2!=-1) goto L1;
		goto L2;
	L1:
		printf("%c",(int)t2)
		t1=t1+1;
		goto L0;
	L2:


*/
