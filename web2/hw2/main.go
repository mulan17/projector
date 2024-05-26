//1. Розробити програму «Зоопарк». Завдання: 5 чи більше звірів повтікали, наглядач повинен їх зібрати.
//Кожну сутність (наглядач, звір, клітка тощо) представляти окремою структурою (zookeeper, animal, cage).
//Користуємось ембдінгом і методами.


// embedding зазвичай використовується, коли треба розширити структуру, в твоєму випадку Cage розширяє тип Animal і це виглядає не зовсім логічним, адже Cage та Animal це сильно різні сутності. 
//Гарним прикладом було б розширити Animal тип типом Monkey або Lion.

// // Також, виглядає не дуже логічним лічильник тварин в клітці, бо завжди в клітці тільки одна тваринка.

// // Враховуючи коментар з лічильником, я б зробив так, щоб функція CollectAnimal не приймала клітку, а створювала нову, складала в неї тваринку та повертала.

// // PS лічильник можна перенести до зукіпера, тоді буде логічно (шось накшталт скільки цей зукіпер зловив тварин взагалі).

package main

import "fmt"

type Zookeeper struct {
	Name        string
	AnimalCount int
}

type Animal struct {
	Name string
	Age  int
}

type AnimalType struct {
	Animal
	Type string
}

type Cage struct {
	Animal AnimalType
}

func (a AnimalType) Print() {
	fmt.Printf("Animal is %v and name is %v\n", a.Type, a.Name)
}

func (c Cage) CageContain() {
	fmt.Printf("Cage contains: %v the %v\n", c.Animal.Name, c.Animal.Type)
}

func (z Zookeeper) NameZookeeper() {
	fmt.Printf("Zookeeper on duty is %v\n", z.Name)
}

func (z *Zookeeper) CollectAnimal(t AnimalType) Cage {
	z.AnimalCount++
	fmt.Printf("%v has collected %v the %v\n", z.Name, t.Name, t.Type)
	return Cage{Animal: t}
}

func main() {
	a1 := Animal{Name: "Leo"}
	a2 := Animal{Name: "Lala"}
	a3 := Animal{Name: "Sisi"}
	a4 := Animal{Name: "Po"}
	a5 := Animal{Name: "Kate"}

	t1 := AnimalType{Animal: a1, Type: "Lion"}
	t2 := AnimalType{Animal: a2, Type: "Tiger"}
	t3 := AnimalType{Animal: a3, Type: "Squirrel"}
	t4 := AnimalType{Animal: a4, Type: "Bear"}
	t5 := AnimalType{Animal: a5, Type: "Cat"}

	z := Zookeeper{Name: "Christian"}

	fmt.Println("Cages are empty, 0 animals are there. Who is on duty today?")
	z.NameZookeeper()

	c1 := z.CollectAnimal(t1)
	c1.CageContain()

	c2 := z.CollectAnimal(t2)
	c2.CageContain()

	c3 := z.CollectAnimal(t3)
	c3.CageContain()

	c4 := z.CollectAnimal(t4)
	c4.CageContain()

	c5 := z.CollectAnimal(t5)
	c5.CageContain()

	fmt.Printf("%v has collected a total of %v animals.\n", z.Name, z.AnimalCount)
}
