//вибачаюсь, що мало різноманітності функцій:)

package main

import (
	"fmt"
)

func minus(a, b int) int {
	return a - b

}

func main() {
	println("Біографія видатного українця Степана Бандери:))")

	//string
	name := "Степан Андрійович Бандера"
	birth := "1 січня 1909"
	yearBirth := 1909
	birthPlace := "селі cтарий Яричів, Австро-Угорщина (нині Україна)"

	fmt.Printf("%s народився %s в %s.\n", name, birth, birthPlace)

	//const
	const shortName, age = "Степан", 25
	university := "юридичний факультет Львівського університету"
	fmt.Printf("%s закінчив %s у віці %d років.\n", shortName, university, age)

	//string
	intro := "Після закінчення університету Степан Бандера активно включився y політичну діяльність, приєднуючись до українського національного руху"
	fmt.Printf("%s.\n", intro)

	//function minus for 2 events
	println("ОСНОВНІ ПОДІЇ")
	event1Number := 1
	event1 := "Заснування та очолення Української Повстанської Армії (УПА)"
	event1Description := "Бандера був одним з ключових організаторів та лідерів УПА, в якій він відіграв важливу роль у боротьбі за незалежність України"
	event1DateStart := 1942
	event1DateEnd := 1950

	fmt.Printf("%d.%s:\n%s.\n", event1Number, event1, event1Description)
	fmt.Printf("Роки: %d-%d\n", event1DateStart, event1DateEnd)
	fmt.Printf("Вік: %d-%d\n", minus(event1DateStart, yearBirth), minus(event1DateEnd, yearBirth))

	event2Number := 2
	event2 := "Заснування та очолення Української Повстанської Армії (УПА)"
	event2Description := "Бандера був одним з ключових організаторів та лідерів УПА, в якій він відіграв важливу роль у боротьбі за незалежність України"
	event2DateStart := 1942
	event2DateEnd := 1950

	fmt.Printf("%d.%s:\n%s.\n", event2Number, event2, event2Description)
	fmt.Printf("Роки: %d-%d\n", event2DateStart, event2DateEnd)
	fmt.Printf("Вік: %d-%d\n", minus(event2DateStart, yearBirth), minus(event2DateEnd, yearBirth))

	//function minus for death
	yearDeath := 1959
	fmt.Printf("%s помер у віці %d років.\n", name, minus(yearDeath, yearBirth))

}
