package valuator

import (
	"github.com/bastianrob/arrayutil"
)

type in struct {
	expr expression
}

//Evaluate returns true if expression value contains obj and is not inversed
func (e *in) Evaluate(obj map[string]interface{}) bool {
	objVal, ok := obj[e.expr.property]
	if !ok {
		return false
	}

	return arrayutil.Contains(e.expr.value, objVal) == !e.expr.inverse
}
