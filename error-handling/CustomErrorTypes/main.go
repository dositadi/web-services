package main

import (
	"fmt"
	"strconv"
)

type InvalidInputError struct {
	Input  string
	Reason string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("Invalid input '%s' : '%s'", e.Input, e.Reason)
}

func CustomAtoi(value string) (int, error) {
	if value == "" {
		return 0, &InvalidInputError{Input: value, Reason: "Cannot parse an empty string"}
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, &InvalidInputError{Input: value, Reason: "Invalid input syntax"}
	}
	return result, nil
}

func main() {
	value, err := CustomAtoi("avc4")
	if err != nil {
		fmt.Println("Error occurred: ", err)
	} else {
		fmt.Println(value)
	}

	// Example 2: Failed parsing (non-numeric string)
	val2, err2 := CustomAtoi("abc")
	if err2 != nil {
		// We can use a type assertion to check if the error is our custom type.
		if inputErr, ok := err2.(*InvalidInputError); ok {
			fmt.Printf("Custom error for 'abc': %s (Input: '%s', Reason: '%s')\n",
				inputErr.Error(), inputErr.Input, inputErr.Reason)
		} else {
			fmt.Printf("Other error for 'abc': %v\n", err2)
		}
	} else {
		fmt.Printf("Parsed 'abc' successfully: %d\n", val2)
	}
}
