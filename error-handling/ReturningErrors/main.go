package main

import (
	"fmt"
	"strconv"
)

func ParseInt(value string) (int, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func main() {
	value, err := ParseInt("")
	if err != nil {
		fmt.Println("Error occurred while parsing value: ", err)
		return
	}
	fmt.Println("Int equivalent: ", value)
}
