package main

import "fmt"

func TypeInferenceExample() {
	// Inferência de tipo usando :=
	name := "Philippe"  // Go infere que é uma string
	age := 28           // Go infere que é um int
	pi := 3.1415        // Go infere que é um float64
	isDeveloper := true // Go infere que é um bool

	// Inferência com operações
	number1 := 10
	number2 := 3.5
	result := float64(number1) + number2 // Go infere como float64 após conversão

	fmt.Println("=== Type Inference in Go ===")
	fmt.Printf("name (%T): %v\n", name, name)
	fmt.Printf("age (%T): %v\n", age, age)
	fmt.Printf("pi (%T): %v\n", pi, pi)
	fmt.Printf("isDeveloper (%T): %v\n", isDeveloper, isDeveloper)
	fmt.Printf("result (%T): %v\n", result, result)
}
