package main

import "fmt"

type Hi struct {
	Name string
	Age int
}

func (p *Hi) Birthday() {
	p.Age++
	fmt.Printf("%v is now %d years old\n", p.Name, p.Age)
}

func (p *Hi) Birthday2() {
	p.Age++
	fmt.Printf("%v is now %d years old\n", p.Name, p.Age)
}

func (p *Hi) Birthday3() {
	p.Age = p.Age + 3
	fmt.Printf("%v is now %d years old\n", p.Name, p.Age)
}


func main(){
p := Hi{Name: "Nastya", Age: 20,}
	p.Birthday()

	fmt.Printf("%v age with pointer\n", p.Age)

	p.Birthday2()

	fmt.Printf("%v age without pointer\n", p.Age)

	p.Birthday3()

	fmt.Printf("%v age without pointer\n", p.Age)

}