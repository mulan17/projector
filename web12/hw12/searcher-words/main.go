package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	filePath := "1689007676028_text.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	vowelRegex := regexp.MustCompile(`[аеєиіїоуюяaei]`)

	words := strings.Fields(string(content))

	matchedWords := []string{}

	wordRegex := regexp.MustCompile(`\b\w*ї\w*\b`)

	for _, word := range words {
		if wordRegex.MatchString(word) {
			vowelsCount := len(vowelRegex.FindAllString(word, -1))
			if vowelsCount >= 3 {
				matchedWords = append(matchedWords, word)
			}
		}
	}

	fmt.Println("Words matching the criteria:")
	for _, word := range matchedWords {
		fmt.Println(word)
	}
}
