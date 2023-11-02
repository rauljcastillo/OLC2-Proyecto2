package src

import (
	generador "main/Generador"
	"main/parser"
)

func (l *Visitor) VisitSwhile(ctx *parser.SwhileContext, amb *Ambiente, gen *generador.Generador) interface{} {
	loop := gen.NewLabel()
	gen.AddLabel(loop)
	expr := l.Visit(ctx.Expr(), amb, gen).(Valor)
	gen.AddLabel(expr.labelTrue)
	l.Visit(ctx.Block(), amb, gen)
	gen.AddGoto(loop)
	gen.AddLabel(expr.labelfalse)
	return 0
	/*
		L0:
			if t1>5 goto L1;
			goto L2;
		L1:
			print(5);
			t1=stack[0]
			t1=t1+2
			stack[0]=t1
			goto L0;
		L2:
	*/
}
