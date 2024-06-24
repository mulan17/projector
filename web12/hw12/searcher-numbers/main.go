package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	filePath := "1689007675141_numbers.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	phoneRegex := regexp.MustCompile(`\b\d{3}[-.\s]?\d{3}[-.\s]?\d{4}\b`)
	phoneNumbers := phoneRegex.FindAllString(string(content), -1)

	for _, number := range phoneNumbers {
		fmt.Println(number)
	}
}
