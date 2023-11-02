package src

type Variable struct {
	tipo     int
	id       string
	posicion int
}

type Funcion struct {
	tipo   int
	ID     string
	Params any
	Block  any
}
