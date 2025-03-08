package main

import "fmt"

func main() {
	var name string = "Philippe"
	age := 28 // Inferência de tipo

	const PI = 3.1415 // Constante

	fmt.Println("Nome:", name)
	fmt.Println("Idade:", age)
	fmt.Println("PI:", PI)

	// Calling the function from primitive_types.go
	PrimitiveTypesExample()
	ZeroValuesExample()
	TypeInferenceExample()

}
