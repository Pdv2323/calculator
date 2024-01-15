package main

import "fmt"

func Input() (float64, float64) {
	var a, b float64

	fmt.Print("Enter 1st integer : ")
	fmt.Scanln(&a)

	fmt.Print("Enter 2nd integer : ")
	fmt.Scanln(&b)

	return a, b
}
