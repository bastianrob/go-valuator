package valuator

type lt struct {
	expr expression
}

//Evaluate returns true if object value < expression value and not inversed
//Both object and expression value must be of type float64
//If object or expression value is not a number, returns false ignoring inverse
func (e *lt) Evaluate(obj map[string]interface{}) bool {
	objVal, ok := obj[e.expr.property]
	if !ok {
		return false
	}

	objFloat64Val, err := getFloatFromInterface(objVal)
	if err != nil {
		return false
	}

	exprFloat64Val, err := getFloatFromInterface(e.expr.value)
	if err != nil {
		return false
	}

	return objFloat64Val < exprFloat64Val == !e.expr.inverse
}
