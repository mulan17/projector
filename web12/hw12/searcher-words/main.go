package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	filePath := "1689007676028_text.txt"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	wordRegex := regexp.MustCompile(`\b[a-zA-Z]*[aeiou][a-zA-Z]*\b`)
	words := wordRegex.FindAllString(string(content), -1)

	for _, word := range words {
		fmt.Println(word)
	}
}
