package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// alternative syntax
	// firstName, lastName, city, gender string
	// age int
}

// Greeting method (value reciever)
// method identifier (p)
func (p Person) greeet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spousedLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spousedLastName
	}
}

func main() {
	// init person
	person1 := Person{firstName: "Semantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	fmt.Println(person1)

	// alternative syntax
	person2 := Person{"Bob", "Johnson", "NY", "m", 30}
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greeet())

	person1.getMarried("Williams")
	fmt.Println(person1.greeet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greeet())
}

// we have two kinde of methods
// value recievers and pointer recievers
