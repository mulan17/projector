//1. Пошук для текстового редактора. Створити slice string з текстом, який користувач ввів у текстовий редактор.
//Написати функцію, яка приймає на вхід рядок для пошуку та знаходить у текстовому редакторі всі рядки, які містять рядок пошуку.
//Використовуючи цю функцію, додати можливість пошуку тексту в текстовому редакторі та вивести на екран усі відповідні результати.
//Розширена задача: ініціалізувати текс в редакторі не через код програми, а зчитавши рядки тексту з файлу

package main

import (
	"fmt"
	"strings"
)

func searchLines(Slice []string, searchWord string) []string {
	var results []string
	for _, word :=range Slice {
		if strings.Contains(word, searchWord) {
			results = append(results, word)
		}
	}
	return results
}

func main() {
	Slice := []string{
		"one",
		"two",
		"three",
	}

	fmt.Println(Slice)

	searchWord := "11"
	results := searchLines(Slice, searchWord)

	if len(results) >0 {
		fmt.Println("Results of searching:")
		for _, result := range results {
			fmt.Println((result))
		}
	} else {
		fmt.Println("Not found")
	}
}