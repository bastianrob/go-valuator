package valuator

import "fmt"

type equality struct {
	expr expression
}

//Evaluate returns true if expression equals to obj
//Both object and expression calue will be treated as string
func (e *equality) Evaluate(obj map[string]interface{}) bool {
	objVal, ok := obj[e.expr.property]
	if !ok {
		return false
	}

	objStringVal := fmt.Sprintf("%v", objVal)
	return objStringVal == e.expr.value == !e.expr.inverse
}
