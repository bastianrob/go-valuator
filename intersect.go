package valuator

import "github.com/bastianrob/arrayutil"

type intersect struct {
	expr expression
}

//Evaluate returns true if expression value intersects with obj and is not inversed
func (e *intersect) Evaluate(obj map[string]interface{}) bool {
	objVal, ok := obj[e.expr.property]
	if !ok {
		return false
	}

	return arrayutil.Intersect(e.expr.value, objVal) == !e.expr.inverse
}
