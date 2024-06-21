//1. Пошук для текстового редактора. Створити slice string з текстом, який користувач ввів у текстовий редактор.
//Написати функцію, яка приймає на вхід рядок для пошуку та знаходить у текстовому редакторі всі рядки, які містять рядок пошуку.
//Використовуючи цю функцію, додати можливість пошуку тексту в текстовому редакторі та вивести на екран усі відповідні результати.
//Розширена задача: ініціалізувати текс в редакторі не через код програми, а зчитавши рядки тексту з файлу

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func readFromFile(myfile string) ([]string, error) {
	file, err := os.Open(myfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var slice []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	return slice, scanner.Err()
}

func searchLines(slice []string, searchWord string) []string {
	var results []string
	for _, word := range slice { // пройтись по кожному рядку слайса
		if strings.Contains(word, searchWord) {
			results = append(results, word)
		}
	}
	return results
}

func main() {
	// myfile := "file1.txt" // open the file
	myfile := flag.String("file", "file1.txt", "File path entering")
	searchWord := flag.String("search", "", "Searching word")

	flag.Parse()

	slice, err := readFromFile(*myfile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// defer myfile.Close()

	// slice := []string{
	// 	"one",
	// 	"two",
	// 	"three",
	// }
	//fmt.Println(Slice)

	// searchWord := "так"
	results := searchLines(slice, *searchWord)

	if len(results) > 0 {
		fmt.Println("Results of searching:")
		for _, result := range results {
			fmt.Println((result))
		}
	} else {
		fmt.Println("Not found")
	}
}
