// Унікальні структури. Створити тип структура, що містить одне поле (наприклад `ID`).
// Написати функцію, яка на вхід приймає слайс з елементами створеного типу, а повертає слайс того ж типу лише з унікальними значеннями
// (структури з дублікатами значення поля відкидаються). Результ функції має бути відсортований у порядку зростання значень поля структури.
// Додаткові умови: Не використовувати бібліотеки для пошуку унікальних значень.
// Використати можливості стандартної бібліотеки `sort` для сортування. Приклад: [{3}, {2}, {1}, {2}] -> [{1}, {2}, {3}]
// створити структуру з полем ід
// функція що приймає слайс з елементами створеного типу

package main

import "fmt"

func main(){
	var firstSlice = make([]string, 4)

	for i := range firstSlice{
		firstSlice[i] = "banana"
	}

	fmt.Println("1. firstSlice", firstSlice)

// 0:len(firstSlice)
	var secondSlice = firstSlice[:]

	secondSlice[1] = "apple"

	fmt.Println("2. firstSlice", firstSlice)
	fmt.Println("2. secondSlice", secondSlice)

	secondSlice = secondSlice[1:]
	fmt.Println("3. firstSlice", firstSlice)
	fmt.Println("3. secondSlice", secondSlice)

	// // a[low:high:max]
	// The indices low and high select which elements of operand a appear in the result.
	// Max controls the resulting slice's capacity by setting it to max - low.
	// Language specs: https://go.dev/ref/spec#Slice_expressions (Full slice expressions)
	thirdSlice := firstSlice[2:4:4] 
	fmt.Println("4. thirdSlice", thirdSlice, cap(thirdSlice))

}

func exampleSlice() {

	var sliceLiteral []int

	sliceLiteral = []int{101, 2020, 303}

	fmt.Println(sliceLiteral, len(sliceLiteral), cap(sliceLiteral))

	//створює масив з 3 елементами, з нього створює слайс, вказуе слайс на початок масиву та засовує в слайс літерал
	sliceMake := make([]int, 1, 3)

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	sliceMake[0] = 101

	poinetToFirstElem := &sliceMake[0]

	fmt.Println("By pointer:", *poinetToFirstElem)

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	sliceMake = append(sliceMake, 202, 303, 404)

	*poinetToFirstElem = 55

	if sliceMake[0] == 55 {
		sliceMake[0] = 505
	}

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))
	
}

func exampleArray() {
	var arr [3]int //3 is lenght

	arr[0] = 1

	arr[0] = 101

	setArrToFives(arr)

	printArr(arr)

	fmt.Println(len(arr), cap(arr))
}

// функція приймає масив з розміром три типом інт та вказує кожний елемент масиву на 5
func setArrToFives(arr [3]int) {
	arr[0] = 5
	arr[1] = 5
	arr[2] = 5

}

// для і ініціюалізувати ранже масиву, тобто ми хочемо щоб і дорівнювалось індексу масива і пройшлось по всім індексам масиву
func printArr(arr [3]int) {
	for i, v := range arr {
		//ми беремо массив і буде змінюватись елемент масиву який ми отримуемо
		v = 202 //змінна містить елемент масиву, а чи зміниться елемент масиву якщо ми з
		arr[i] = 202
		fmt.Println(i, arr[i], v)
	}

}
