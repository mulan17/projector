package main

import "fmt"

type Age int

func main() {
	// name := Name("N")
	// // n := "Sasha"

	// age := 36

	person := Person {
		Name: "Nastya",
		Age: 36,
		Education: "NAU",

	}

	person.print()
	person.printEducation()

	var nonamePerson Person

	nonamePerson.print()
	nonamePerson.printEducation()

	p := Person{"Stepan", 45, "Phd"}

	p.print()
	p.printEducation()

	// printPerson(name, Age(age))
	// name.print()
	// printName(name)
	// printName(Name(n))

	// fmt.Println()
}
type Name string

func (n Name) print() {

	fmt.Printf("The name is %v\n", n)
}

type Person struct {
	Name Name
	Age Age
	Education string
}

func (p Person) print(){

	p.Name.print()
	fmt.Printf("And their age is %v\n", p.Age )

}

func (p Person) printEducation(){
	fmt.Println(p.Education)
}