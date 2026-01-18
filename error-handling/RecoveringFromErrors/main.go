package main

import "fmt"

func Divide(a, b int) int {
	if b == 0 {
		panic("The divisor is zero")
	}
	return a / b
}

func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("recovered from panic %v", r)
		}
	}()
	result = a / b
	return result, nil
}

func main() {
	a, b := SafeDivide(6, 0)
	fmt.Println(a, b)
}
