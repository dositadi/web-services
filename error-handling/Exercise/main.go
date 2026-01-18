package main

import (
	"errors"
	"fmt"
)

var Err = errors.New("operation failed at C")

func funcC() error {
	return Err
}

func funcB() error {
	return fmt.Errorf("Failed in B due to C: %w", funcC())
}

func funcA() error {
	return fmt.Errorf("failed in A due to B: %w", funcB())
}

func main() {
	err := funcA()
	if err != nil {
		fmt.Printf("%+v\n", err)
		if errors.Is(err, Err) {
			fmt.Println("Specific error for Err contained here!")
		}
	}
}
