package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	filePath := "1689007676028_text.txt"

	wordRegex := regexp.MustCompile(`\b[a-zA-Z]*[aeiou][a-zA-Z]*[bcdfghjklmnpqrstvwxyz]|[a-zA-Z]*([a-zA-Z])\1[a-zA-Z]*\b`)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	text := string(content)

	words := strings.Fields(text)

	var matchingWords []string
	for _, word := range words {
		if wordRegex.MatchString(word) {
			matchingWords = append(matchingWords, word)
		}
	}

	for _, word := range matchingWords {
		fmt.Println(word)
	}
}
