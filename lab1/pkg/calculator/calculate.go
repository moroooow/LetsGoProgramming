package calculator

import "errors"

var ErrIncorrectInput = errors.New("Incorrect format")
var ErrDivisionByZero = errors.New("Division by zero")

const Add, Sub, Div, Mul = "+", "-", "/", "*"

func Calculate(a float64, b float64, sign string) (float64, error) {
	switch sign {
	case Add:
		return a + b, nil
	case Sub:
		return a - b, nil
	case Mul:
		return a * b, nil
	case Div:
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, ErrIncorrectInput
	}
}
