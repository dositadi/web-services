package main

import (
	"errors"
	"fmt"
)

//Sentinel errors are specific, predefined error variables. They are typically defined at the package level and exported for other packages to use. These errors are compared directly using the == operator.

var ErrNotFound = errors.New("Item not Found")

func FindItem(id string) (bool, error) {
	if id == "valid" {
		return true, nil
	}
	return false, ErrNotFound
}

func main() {
	found1, err := FindItem("My own!")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(found1)
	}

	found2, err2 := FindItem("My game!")
	if err2 != nil {
		if errors.Is(err2, ErrNotFound) {
			fmt.Println("Error is: ", ErrNotFound)
		} else {
			fmt.Println("Some Other error!")
		}
	} else {
		fmt.Println(found2)
	}
}
