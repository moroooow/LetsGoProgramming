package calculator

import "testing"

func TestAdd(t *testing.T) {
	var expected float64 = 10
	var a float64 = 5
	var b float64 = 5
	res, _ := Calculate(a, b, Add)
	if expected != res {
		t.Errorf("Expected %f,received %f", expected, res)
	}
}

func TestSub(t *testing.T) {
	var expected float64 = 0
	var a float64 = 5
	var b float64 = 5
	res, _ := Calculate(a, b, Sub)
	if expected != res {
		t.Errorf("Expected %f,received %f", expected, res)
	}
}

func TestMul(t *testing.T) {
	var expected float64 = 25
	var a float64 = 5
	var b float64 = 5
	res, _ := Calculate(a, b, Mul)
	if expected != res {
		t.Errorf("Expected %f,received %f", expected, res)
	}
}

func TestDiv(t *testing.T) {
	var expected float64 = 1
	var a float64 = 5
	var b float64 = 5
	res, _ := Calculate(a, b, Div)
	if expected != res {
		t.Errorf("Expected %f,received %f", expected, res)
	}
}

func TestDivByZero(t *testing.T) {
	var a float64 = 5
	var b float64 = 0
	_, err := Calculate(a, b, Div)
	if err.Error() != ErrDivisionByZero.Error() {
		t.Errorf("Expected %s,received %s", ErrDivisionByZero.Error(), err.Error())
	}
}
