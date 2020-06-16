package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// assigne
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// declare and assigne
	vegetArra := [2]string{"Carrots", "Salata"}

	// slices
	meet := []string{"Cow", "Sheep", "Pig"}


	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(vegetArra)
	fmt.Println(len(meet))
	fmt.Println(meet[1:3])
}

