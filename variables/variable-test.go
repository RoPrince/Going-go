package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Roshni"
	address := "Santa Cruz"
	age := 7
	salary := 50
	fmt.Println("Name & address ", name, " ", address, ".")
	fmt.Println("Age & salary ", age, " ", salary)
	fmt.Println("Type of Age:", reflect.TypeOf(age))

	fmt.Println("Memory of Address", &address)

	var ptr *string = &address
	fmt.Println("Value of Pointer", *ptr)
	fmt.Println("Address points to ", ptr)
}
