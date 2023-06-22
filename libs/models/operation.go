package models

type Operator string

const (
	ADDITION       Operator = "ADDITION"
	SUBSTRACTION   Operator = "SUBSTRACTION"
	MULTIPLICATION Operator = "MULTIPLICATION"
	DIVISION       Operator = "DIVISION"
	SQUARE_ROOT    Operator = "SQUARE_ROOT"
	RANDOM_STRING  Operator = "RANDOM_STRING"
)

type Operation struct {
	Operation    Operator `faker:"operation" json:"operation" bson:"operation"`
	Operand1     float64  `faker:"operand1" json:"operand1" bson:"operand1"`
	Operand2     float64  `faker:"operand2" json:"operand2" bson:"operand2"`
	Result       *float64 `faker:"result" json:"result" bson:"result"`
	StringResult *string  `faker:"stringResult" json:"stringResult" bson:"stringResult"`
	Cost         float64  `faker:"cost" json:"cost" bson:"cost"`
}

func NewOperation(operation Operator, operand1 float64, operand2 float64) *Operation {
	return &Operation{
		Operation: operation,
		Operand1:  operand1,
		Operand2:  operand2,
	}
}
