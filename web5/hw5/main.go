//Імплементувати пошук для текстового редактора (аналогічно до завдання в HW4) використовуючи індекс слів у мапі.
//Тобто, для текстового редактора реалізувати методи "проіндексувати текст по словам", та "пошук усіх рядків за словом".

//імплементувати словник англ - укр. Користувач вводить слово, програма - переклад

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

func indexInText(slice []string) map[string][]int {
	index := make(map[string][]int)
	for i, line := range slice {
		words := strings.Fields(line)
		for _, word := range words {
			index[word] = append(index[word], i)
		}
	}
	return index
}

func searchLines(index map[string][]int, slice []string, word string) []string {
	numberLine, found := index[word]
	if !found {
		return nil
	}

	var results []string
	for _, numberLine := range numberLine {
		results = append(results, slice[numberLine])
	}
	return results
}

func main() {
	myfile := flag.String("file", "file1.txt", "File path entering")
	searchWord := flag.String("так", "", "Searching word")

	flag.Parse()

	slice, err := readFromFile(*myfile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	index := indexInText(slice)

	results := searchLines(index, slice, *searchWord)

	if len(results) > 0 {
		fmt.Println("Results of searching:")
		for _, result := range results {
			fmt.Println((result))
		}
	} else {
		fmt.Println("Not found")
	}
}
