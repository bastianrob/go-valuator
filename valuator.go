package valuator

import (
	"encoding/json"
	"errors"
	"strings"
)

//Error collections
var (
	ErrUnrecognized  = errors.New("Valuator for operator is not recognized")
	ErrInvalidFormat = errors.New("Expression value is not valid")
)

//Expression of a rule
type expression struct {
	inverse     bool
	property    string
	operator    string
	value       interface{}
	description string
}

//Evaluator expression evaluator
type Evaluator interface {
	//Evaluate returns true if expression is valid
	Evaluate(obj map[string]interface{}) bool
}

//NewValuator picked based on operator
func NewValuator(prop, op, val, desc string) (Evaluator, error) {
	inverse := strings.HasPrefix(op, "!")
	if inverse {
		runes := []rune(op)
		op = string(runes[1:])
	}

	expr := expression{inverse, prop, op, val, desc}
	switch op {
	case "=":
		return &equality{expr}, nil
	case "in":
		arr := []interface{}{}
		err := json.Unmarshal([]byte(val), &arr)
		if err != nil {
			return nil, ErrInvalidFormat
		}

		expr.value = arr
		return &in{expr}, nil
	case "intersect":
		arr := []interface{}{}
		err := json.Unmarshal([]byte(val), &arr)
		if err != nil {
			return nil, ErrInvalidFormat
		}

		expr.value = arr
		return &intersect{expr}, nil
	case ">":
		return &gt{expr}, nil
	case ">=":
		return &gte{expr}, nil
	case "<":
		return &lt{expr}, nil
	case "<=":
		return &lte{expr}, nil
	}

	return nil, ErrUnrecognized
}
