//1. Розробити програму «Зоопарк». Завдання: 5 чи більше звірів повтікали, наглядач повинен їх зібрати.
//Кожну сутність (наглядач, звір, клітка тощо) представляти окремою структурою (zookeeper, animal, cage).
//Користуємось ембдінгом і методами.

package main

import "fmt"

type Zookeeper struct {
	Name string
}

type Animal struct {
	Name string
	Type string
}

type Cage struct{
	Animal
	Contain int
}

func (a Animal) Print(){
	fmt.Printf("Animal is %v and name is %v\n", a.Type, a.Name)
}

func (c *Cage) CageContain() {
	fmt.Printf("Cage contains: %v\n", c.Contain)
	
}

func (z Zookeeper) NameZookeper(){
	fmt.Printf("Zookeper on duty is %v\n", z.Name)
} 

func (z Zookeeper) CollectAnimal(a Animal, c *Cage) {
	c.Animal = a
	c.Contain++
	fmt.Printf("%v has collected %v the %v\n", z.Name, a.Name, a.Type)
}

func main(){
	a1 := Animal{Name: "Leo", Type: "Lion"}
	a2 := Animal{Name: "John", Type: "Tiger"}
	a3 := Animal{Name: "Kate", Type: "Squirrel"}
	a4 := Animal{Name: "Po", Type: "Bear"}
	a5 := Animal{Name: "Anna", Type: "Cat"}

	c := Cage{Contain: 0}
	z := Zookeeper{Name: "Christian"}

	fmt.Printf("Cage is empty, %v animal is there. Who is on duty today?\n", c.Contain)
	z.NameZookeper()

	z.CollectAnimal(a1, &c)
	c.CageContain()

	z.CollectAnimal(a2, &c)
	c.CageContain()

	z.CollectAnimal(a3, &c)
	c.CageContain()

	z.CollectAnimal(a4, &c)
	c.CageContain()

	z.CollectAnimal(a5, &c)
	c.CageContain()
}