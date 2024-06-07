//Імплементувати пошук для текстового редактора (аналогічно до завдання в HW4) використовуючи індекс слів у мапі. 
//Тобто, для текстового редактора реалізувати методи "проіндексувати текст по словам", та "пошук усіх рядків за словом".

//імплементувати словник англ - укр. Користувач вводить слово, програма - переклад

package main

import (
	"bufio"
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

func main() {
	myfile := "file1.txt" // open the file
	slice, err := readFromFile(myfile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// searchWord := "так"
	// results := searchLines(slice, searchWord)

	if len(results) > 0 {
		fmt.Println("Results of searching:")
		for _, result := range results {
			fmt.Println((result))
		}
	} else {
		fmt.Println("Not found")
	}
}