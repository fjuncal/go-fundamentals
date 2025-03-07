package main

import "fmt"

func ZeroValuesExample() {
	// Declaração de variáveis sem inicialização
	var intNumber int
	var floatNumber float64
	var booleanValue bool
	var stringValue string
	var runeValue rune
	var pointer *int

	fmt.Println("=== Zero Values in Go ===")
	fmt.Println("Default int:", intNumber)       // 0
	fmt.Println("Default float64:", floatNumber) // 0.0
	fmt.Println("Default bool:", booleanValue)   // false
	fmt.Println("Default string:", stringValue)  // ""
	fmt.Println("Default rune:", runeValue)      // 0 (interpreted as '\u0000')
	fmt.Println("Default pointer:", pointer)     // nil
}
