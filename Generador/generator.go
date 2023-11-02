package generador

import (
	"fmt"
	"strings"
)

type Generador struct {
	temporal      int
	label         int
	listaTemporal []string
	code          []string
	nativas       []string
	functions     []string
	print         bool
	printMath     bool
	concat        bool
}

func NewGenerator() *Generador {
	return &Generador{listaTemporal: make([]string, 0), print: false, printMath: false, concat: false}
}

func (g *Generador) NewTemporal() string {
	temp := "t" + fmt.Sprint(g.temporal)
	g.temporal = g.temporal + 1
	g.listaTemporal = append(g.listaTemporal, temp)
	return temp
}

// Genera nueva etiqueta
func (g *Generador) NewLabel() string {
	tmp := "L" + fmt.Sprint(g.label)
	g.label = g.label + 1
	return tmp
}

func (g *Generador) AddLabel(label string) {
	g.code = append(g.code, label+":"+"\n")
}

func (g *Generador) Change(oldlbl, newlbl string) {
	g.code[len(g.code)-1] = strings.ReplaceAll(g.code[len(g.code)-1], oldlbl, newlbl)
}

// Agrega al arreglo de Codigo
func (g *Generador) AddCode(code string) {
	g.code = append(g.code, code)
	//g.code = append(g.code, code+"\n")
}

func (g *Generador) SetHeap(position, value string) {
	temp := "heap[" + "(int)" + position + "] = " + value + ";" + "\n"
	g.code = append(g.code, temp)
}

func (g *Generador) GetHeap(temp, position string) {
	tem := temp + " = " + "heap[(int)" + position + "]" + ";" + "\n"
	g.code = append(g.code, tem)
}

func (g *Generador) SetStack(position, value string) {
	temp := "stack[" + position + "] = " + value + ";" + "\n"
	g.code = append(g.code, temp)
}

func (g *Generador) GetStack(temp, position string) {
	tmp := temp + " = " + "stack[" + position + "]" + ";" + "\n"
	g.code = append(g.code, tmp)
}

// Genera Goto
func (g *Generador) AddGoto(label string) {
	temp := "goto " + label + ";" + "\n"
	g.code = append(g.code, temp)
	//g.code = append(g.code, "goto "+label+";"+"\n")
}

// Genera la llamada a una funcion
func (g *Generador) Call(name string) {
	g.code = append(g.code, name+"();\n")
}

// Genera la expresion id = exp op exp
func (g *Generador) AddExpresion(temp, izq, der, op string) {
	exp := temp + " = " + izq + op + der + ";" + "\n"
	g.code = append(g.code, exp)
	//g.code = append(g.code, exp)
}

// tipovalor %d% o %c%,etc
func (g *Generador) Printf(tipoc, value string) {
	temp := "printf(" + "\"%" + tipoc + "\"" + "," + value + ")" + ";" + "\n"
	g.code = append(g.code, temp)
	//g.code = append(g.code, temp)

}

// Crear el if para un condicional
func (g *Generador) Addif(l, r, op, label string) {
	temp := "if(" + l + op + r + ") goto " + label + ";" + "\n"
	g.code = append(g.code, temp)
	//g.code = append(g.code, temp)
}

func (g *Generador) Addbr() {
	g.code = append(g.code, "\n")
}

func (g *Generador) AddCommment(comment string) {
	g.code = append(g.code, "//"+comment+"\n")
}

func (g *Generador) GetCode() string {
	final := "double " + strings.Join(g.listaTemporal, ", ") + ";"
	final = final + "\n\n" + strings.Join(g.nativas, "")
	final = strings.Join(g.functions, "") + final
	main := "int main(){\n"
	main = main + strings.Join(g.code, "")
	main = main + "}\n"
	final = final + main
	return final
}

func (g *Generador) Contiene(name string) bool {
	for _, value := range g.functions {
		if strings.Contains(value, name) {
			return true
		}
	}
	return false
}

// Mueve lo que hay en g.code a Functions
func (g *Generador) MovFunction(index int) {
	g.functions = append(g.functions, g.code[index:]...)
	g.functions = append(g.functions, "}\n\n")
	g.code = g.code[:index]
}

// Devuelve el tamaño de g.code
func (g *Generador) LenCode() int {
	return len(g.code)
}

func (g *Generador) AddFunction(data string) {
	g.functions = append(g.functions, data+"\n")
}

func (g *Generador) Comparar() {
	index := len(g.code)
	g.nativas = append(g.nativas, "void compare_string(){\n")
	tmp1 := g.NewTemporal()
	tmp2 := g.NewTemporal()
	tmp3 := g.NewTemporal()
	tmp4 := g.NewTemporal()
	tmp5 := g.NewTemporal()
	g.GetStack(tmp1, "(int)P")
	g.AddExpresion(tmp2, "P", "1", "+")
	g.GetStack(tmp3, "(int)"+tmp2)
	lbl1 := g.NewLabel()
	lbl2 := g.NewLabel()
	lbl3 := g.NewLabel()
	lbl4 := g.NewLabel()
	lbl5 := g.NewLabel()
	lbl6 := g.NewLabel()
	g.AddLabel(lbl1)
	g.GetHeap(tmp4, tmp1)
	g.GetHeap(tmp5, tmp3)
	g.Addif(tmp4, tmp5, "==", lbl2)
	g.AddGoto(lbl5)
	g.AddLabel(lbl2)
	g.Addif(tmp5, "-1", "!=", lbl3)
	g.AddGoto(lbl4)
	g.AddLabel(lbl3)
	g.AddExpresion(tmp1, tmp1, "1", "+")
	g.AddExpresion(tmp3, tmp3, "1", "+")
	g.AddGoto(lbl1)
	g.AddLabel(lbl4)
	g.SetStack("(int)P", "1")
	g.AddGoto(lbl6)
	g.AddLabel(lbl5)
	g.SetStack("(int)P", "0")
	g.AddLabel(lbl6)
	g.nativas = append(g.nativas, g.code[index:]...)
	g.nativas = append(g.nativas, "}\n\n")
	g.code = g.code[:index]

}

func (g *Generador) Concat() {
	if !g.concat {
		index := len(g.code)
		g.nativas = append(g.nativas, "void concat(){\n")

		tmp1 := g.NewTemporal()
		g.GetStack(tmp1, "(int)P")
		tmp2 := g.NewTemporal()
		g.AddExpresion(tmp2, "P", "1", "+")
		tmp3 := g.NewTemporal()
		g.GetStack(tmp3, "(int)"+tmp2)
		tmp4 := g.NewTemporal()
		g.AddExpresion(tmp4, "H", "", "")

		lbl1 := g.NewLabel()
		lbl2 := g.NewLabel()
		lbl3 := g.NewLabel()
		lbl4 := g.NewLabel()
		lbl5 := g.NewLabel()

		g.AddLabel(lbl1)

		tmp5 := g.NewTemporal()
		g.GetHeap(tmp5, tmp1)
		g.Addif(tmp5, "-1", "!=", lbl2)
		g.AddGoto(lbl3)

		g.AddLabel(lbl2)
		g.SetHeap("H", tmp5)
		g.AddExpresion(tmp1, tmp1, "1", "+")
		g.AddExpresion("H", "H", "1", "+")
		g.AddGoto(lbl1)
		g.AddLabel(lbl3)

		tmp6 := g.NewTemporal()
		g.GetHeap(tmp6, tmp3)
		g.Addif(tmp6, "-1", "!=", lbl4)
		g.AddGoto(lbl5)

		g.AddLabel(lbl4)
		g.SetHeap("H", tmp6)
		g.AddExpresion(tmp3, tmp3, "1", "+")
		g.AddExpresion("H", "H", "1", "+")
		g.AddGoto(lbl3)

		g.AddLabel(lbl5)
		g.SetHeap("H", "-1")
		g.AddExpresion("H", "H", "1", "+")
		g.SetStack("(int)P", tmp4)

		g.nativas = append(g.nativas, g.code[index:]...)
		g.nativas = append(g.nativas, "}\n\n")
		g.code = g.code[:index]
		g.concat = true

	}
}

func (g *Generador) Prints() {
	if !g.print {
		index := len(g.code)

		g.nativas = append(g.nativas, "void prints(){\n")

		temp := g.NewTemporal() //Temporal para guardar el apuntador proveniente del stack
		tmph := g.NewTemporal() //Temporal para guardar el valor proveniente del heap
		g.GetStack(temp, "(int)P")

		loop := g.NewLabel()
		lblprint := g.NewLabel()
		lblexit := g.NewLabel()

		g.AddLabel(loop)
		g.GetHeap(tmph, "(int)"+temp)
		g.Addif(tmph, "-1", "!=", lblprint)
		g.AddGoto(lblexit)

		g.AddLabel(lblprint)
		g.Printf("c", "(int)"+tmph)
		g.AddExpresion(temp, temp, "1", "+")
		g.AddGoto(loop)

		g.AddLabel(lblexit)
		g.Printf("c", "10")
		g.nativas = append(g.nativas, g.code[index:]...)
		g.nativas = append(g.nativas, "}\n\n")
		g.code = g.code[:index]
		g.print = true
	}

}
