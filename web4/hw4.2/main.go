// 2. Унікальні структури. Створити тип структура, що містить одне поле (наприклад `ID`). 
// Написати функцію, яка на вхід приймає слайс з елементами створеного типу, а повертає слайс того ж типу лише з унікальними значеннями 
// (структури з дублікатами значення поля відкидаються). Результ функції має бути відсортований у порядку зростання значень поля структури. 
// Додаткові умови: Не використовувати бібліотеки для пошуку унікальних значень. 
// Використати можливості стандартної бібліотеки `sort` для сортування. Приклад: [{3}, {2}, {1}, {2}] -> [{1}, {2}, {3}]


package main

import (
	"fmt"
	"sort"
)

type Struct struct {
	ID int
}


func sortID(slice []Struct) {
	sort.Slice(slice, func(i, j int) bool{
		return slice[i].ID < slice[j].ID
	})
}

func unique(slice []Struct) []Struct {
	if len(slice) == 0 {
		return slice
	}

	var uniqueSlice []Struct
	uniqueSlice = append(uniqueSlice, slice[0])

	for i := 1; i < len(slice); i++ {
		if slice[i].ID != slice[i-1].ID {
			uniqueSlice = append(uniqueSlice, slice[i])
		}
	}

	return uniqueSlice
}


func main() {
	sliceID := []Struct{
		{ID: 12345},
		{ID: 12445},
		{ID: 12345},
		{ID: 12345},
		{ID: 12645},
		{ID: 12385},
		{ID: 12345},

	}
	
	sortID(sliceID)
	result := unique(sliceID)

	fmt.Println("Sorted and Unique IDs:")
	for _, number := range result {
		fmt.Printf("ID: %d\n", number.ID)
	}


}