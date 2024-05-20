package main

import "fmt"

type Name string

func main() {
	name := Name("Nastya")
	// n := "Sasha"

	name.print()
	// printName(name)
	// printName(Name(n))

	// fmt.Println()
}

func (n Name) print() {

	fmt.Printf("The name is %v\n", n)
}
