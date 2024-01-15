package main

import (
	"fmt"
	"os"
	"strconv"

	dmas "github.com/Pdv2323/Calculator/DMAS"
)

func main() {
	// a, b := Input()

	// fmt.Print("What action you want to perform?")

	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <num1> <operator> <num2>")
		os.Exit(1)
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error parsing num1 : ", err)
		os.Exit(1)
	}

	operator := os.Args[2]

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error parsing num2 : ", err)
	}

	var result float64
	var operation string

	switch operator {
	case "+":
		result = dmas.Add(num1, num2)
		operation = "Addition"

	case "-":
		result = dmas.Subtract(num1, num2)
		operation = "Subtraction"

	case "*":
		result = dmas.Multiply(num1, num2)
		operation = "Multiplication"

	case "/":
		res, err := dmas.Divide(num1, num2)
		if err != nil {
			fmt.Println("Error : ", err)
			os.Exit(1)
		}
		result = res
		operation = "Division"

	default:
		fmt.Println("Invalid Operator : ", operation)
		os.Exit(1)
	}

	fmt.Printf("Result of %s %s %s is : %0.2f\n", os.Args[1], operator, os.Args[3], result)
}
