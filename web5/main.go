//хещування це процес де ми викликаемо якусь функцію, яка може нам залежності від вхідних данних дати стале число,
//тобто при один вхідних данних завжди будуть одні вихідні дані. Ми не можемо з вихідного числа отримати вхідне. І це є як раз хеш.
//Тобто хеш функція видає нам хеш.
//Ці хеши є побудовою хеш таблиц. У функції є ключ і з цієї вибірки ключів ми можемо за допомогою хеш функції
//отримати хеш котрий буде індексом в хеш таблиці. Ми можемо зрозуміти в яку ділянку памʼяти покласти яке значення, який
//ми взяли ключ отримати за допомогою хеш функції хеш, він же індекс та поклали в певний елемент списку
//колізія - це коли ключ дає один і теж самий хещ, тому ділянки розшириють через лінк ліст
//мапи - швидкий доступ по ключу, порядок елементів невідомий, просте видалення елементу.
//слайси - доступ по індексу, чіткий визначенний список, може містити дублікати

// type Person sctuct {
// 	Name string
// 	Age int
// }

// var m = map[string]Person{
// 	"John": Person{
// 		Name: "John",
// 		Age: 22,
// 	},
// 	"Stasy": Person{
// 		Name: "Stasy",
// 		Age: 25,
// 	},
// }

package main

import (
	"fmt"
	"hashtable/hashtable"
)

func main() {
	var h map[string]string

	fmt.Println(h["Marry"])

	h = make(map[string]string)

	h["John"] = "Software Engineer"

	h["James"] = "Plumber"

	h["Stasy"] = "Manager"

	fmt.Println(h["John"])
	fmt.Println(h["James"])
	fmt.Println(h["Stasy"])

	jv := h["John"]
	fmt.Println(jv)

	sv, ok := h["Stasy"]
	if ok {
		fmt.Println("Stacy is", sv)
	} else {
		fmt.Println("No info for Stasy")
	}

	delete(h, "Stacy")

	sv, ok = h["Stacy"]
	if ok {
		fmt.Println("Stacy is", sv)
	} else {
		fmt.Println("No info for Stasy")
	}

	fmt.Println("ITERATION")

	for k, v := range h {
		fmt.Println(k,v)
	}

	h["a"] = "b"

	for k, v := range h {
		fmt.Println(k,v)
	}

}

func exampleHashTable() {
	h := hashtable.New()

	h.Set("John", "Software Engineer")

	h.Set("James", "Plumber")

	h.Set("Stasy", "Manager")

	fmt.Println(h.Get("John"))
	fmt.Println(h.Get("James"))
	fmt.Println(h.Get("Stasy"))

	fmt.Println(h.Get("Marry"))

	fmt.Println("AFTER DELETE")

	h.Delete("Marry")

	h.Delete("John")

	if v, ok := h.Get("John"); ok {
		fmt.Println("John is", v)
	} else {
		fmt.Println("No info for John")
	}
	
}