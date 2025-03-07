package main

import "fmt"

func PrimitiveTypesExample() {
	// Integers
	var intNumber int = 42
	var int8Number int8 = 127 // Ranges from -128 to 127
	var int16Number int16 = 32767
	var int32Number int32 = 2147483647
	var int64Number int64 = 9223372036854775807

	// Unsigned integers (only positive numbers)
	var uintNumber uint = 42
	var uint8Number uint8 = 255 // Ranges from 0 to 255
	var uint16Number uint16 = 65535
	var uint32Number uint32 = 4294967295
	var uint64Number uint64 = 18446744073709551615

	// Floating point numbers
	var float32Number float32 = 3.14
	var float64Number float64 = 3.14159265359

	// Boolean values
	var isActive bool = true
	var isDisabled bool = false

	// Strings
	var message string = "Hello, Golang!"

	// Rune (represents a Unicode character)
	var letter rune = 'G' // Runes are like chars in other languages

	// Printing values
	fmt.Println("=== Primitive Types in Go ===")
	fmt.Println("Integer (int):", intNumber)
	fmt.Println("Integer (int8):", int8Number)
	fmt.Println("Integer (int16):", int16Number)
	fmt.Println("Integer (int32):", int32Number)
	fmt.Println("Integer (int64):", int64Number)
	fmt.Println("Unsigned Integer (uint):", uintNumber)
	fmt.Println("Unsigned Integer (uint8):", uint8Number)
	fmt.Println("Unsigned Integer (uint16):", uint16Number)
	fmt.Println("Unsigned Integer (uint32):", uint32Number)
	fmt.Println("Unsigned Integer (uint64):", uint64Number)
	fmt.Println("Floating Point (float32):", float32Number)
	fmt.Println("Floating Point (float64):", float64Number)
	fmt.Println("Boolean (isActive):", isActive)
	fmt.Println("Boolean (isDisabled):", isDisabled)
	fmt.Println("String (message):", message)
	fmt.Println("Rune (letter):", string(letter)) // Convert rune to string
}
