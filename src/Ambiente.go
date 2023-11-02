package src

import "fmt"

type Ambiente struct {
	anterior  *Ambiente
	variables map[string]*Variable
	functions map[string]Funcion
	Position  int
	//map[]type
}

func NewAmbiente(ant *Ambiente) *Ambiente {

	return &Ambiente{anterior: ant, variables: make(map[string]*Variable), Position: ant.Position}

}

func Global() *Ambiente {
	return &Ambiente{anterior: nil, variables: make(map[string]*Variable), functions: make(map[string]Funcion), Position: 0}
}

func (l *Ambiente) SaveVariable(tipo int, id string) {
	if _, ok := l.variables[id]; !ok {
		l.variables[id] = &Variable{tipo: tipo, id: id, posicion: l.Position}
		l.Position = l.Position + 1
	} else {
		fmt.Println("Ya existe la variable " + id)
	}

}

func (l *Ambiente) saveFunctions(tipe int, id string, param any, bloque any) bool {
	if _, ok := l.functions[id]; !ok {
		l.functions[id] = Funcion{tipo: tipe, ID: id, Params: param, Block: bloque}
		return true
	}
	fmt.Println("Ya existe la función")
	return false
}

func (l *Ambiente) getFunction(id string) interface{} {
	for l != nil {
		if valor, ok := l.functions[id]; ok {
			return valor
		}
		l = l.anterior
	}

	return nil
}

func (l *Ambiente) getVar(id string) interface{} {
	for l != nil {
		if variab, ok := l.variables[id]; ok {
			return variab
		}
		l = l.anterior
	}
	fmt.Println("No existe la variable " + id)
	return nil
}

func (l *Ambiente) SetSize() {
	l.Position = l.Position + 1
}
