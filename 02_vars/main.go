package main

import "fmt"

// name := "Brad" non-declaration statement outside funciton body

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte alias for uint8
	// rune alias for int32
	// float32 float64
	// complex64 comples128

	var name string = "Brad"
	// same as
	// var name = "Brad"

	// in go vars need to be used otherwise we get an error
	// var age int = 37

	// verbs are format pointers
	// %T go syntax representation of the type of the value

	var isCool = true

	const notCool = false
	// notCool = true error cennot assign to notCool

	// short hand method can be only used in the function
	marko := "Marko"

	size := 1.3 // sets automatically to float64

	name, email := "Brad", "marko.antolovic@gmail.com"

	fmt.Println(name)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isCool)
	fmt.Println(marko)
	fmt.Printf("%T\n", marko)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", email)
}
