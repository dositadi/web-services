package main

import "fmt"

type MyInterface interface {
	DoSomething()
}

type MyStruct struct {
	value string
}

func (m *MyStruct) DoSomething() {
	if m == nil {
		fmt.Println("DoSomething is called on a nil type")
	} else {
		fmt.Println("DoSomething is doing something with the value `", m.value, "`")
	}
}

// An Interface is nil only when it is not assigned a type at all
func main() {
	var i MyInterface // in this case i is a nil pointer
	fmt.Println("i: ", i, ", Is i nil: ", i == nil)

	var s *MyStruct // In this case the struct s is nil
	fmt.Println("s: ", s, ", Is s nil: ", s == nil)

	i = s // In this case i is not nil but s is nil
	fmt.Println("Is i nil: ", i == nil, ", is s nil: ", s == nil)
	i.DoSomething()

	nonNil := &MyStruct{
		value: "Hello",
	}
	i = nonNil
	fmt.Println("Is i nil: ", i == nil, ", is nonNil nil: ", nonNil == nil)
	i.DoSomething()
}
