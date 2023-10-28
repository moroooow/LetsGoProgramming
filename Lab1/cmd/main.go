package main

import (
	calc "Lab1/pkg/calculator"
	"fmt"
	"os"
)

func main() {
	var a float64
	var b float64
	var result float64
	var sign string
	var err error

	fmt.Print("Enter first number:")
	_, err = fmt.Scanf("%f\n", &a)
	if err != nil {
		fmt.Println(calc.ErrIncorrectInput.Error())
		os.Exit(1)
	}

	fmt.Print("Enter second number:")
	_, err = fmt.Scanf("%f\n", &b)
	if err != nil {
		fmt.Println(calc.ErrIncorrectInput.Error())
		os.Exit(1)
	}

	fmt.Printf("Enter action (%s, %s, %s, %s):", calc.Add, calc.Sub, calc.Div, calc.Mul)
	_, err = fmt.Scanf("%s\n", &sign)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err = calc.Calculate(a, b, sign)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Result", result)
}
