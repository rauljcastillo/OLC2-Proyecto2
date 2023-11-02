package src

type Valor struct {
	tipo       int
	val        interface{}
	isTemp     bool
	labelTrue  string
	labelfalse string
	position   int
}

type Sentences struct {
	expresion string
}

const (
	STRING = iota
	FLOAT
	INT
	BOOL
	CHAR
	VOID
)

func Num(tipo string) int {
	switch tipo {
	case "Float":
		return FLOAT
	case "Int":
		return INT
	case "String":
		return STRING
	case "Bool":
		return BOOL
	case "Character":
		return CHAR
	}
	return 0
}

func Str(tipo int) string {
	switch tipo {
	case FLOAT:
		return "Float"
	case INT:
		return "Int"
	case STRING:
		return "String"
	case BOOL:
		return "Bool"
	}
	return ""
}
