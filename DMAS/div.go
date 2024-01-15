package dmas

import "fmt"

func Divide(a, b float64) (float64, error) {
	// sum := fmt.Sprintf("%.2f", a/b)
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	} else if a == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	} else if a == 0 && b == 0 {
		return 0, fmt.Errorf("both the values are zero")
	}
	return a / b, nil
}
