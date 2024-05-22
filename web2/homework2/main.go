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


func main(){
	a := Animal{Name: "Leo", Type: "Lion"}
	c := Cage{Contain: 1}
	z := Zookeeper{Name: "Christian"}

	a.Print()
	z.NameZookeper()
	c.CageContain()

}