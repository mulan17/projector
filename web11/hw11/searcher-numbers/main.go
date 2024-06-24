package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	filePath := "1689007675141_numbers.txt"

	// Define the regular expression for phone numbers
	phoneRegex := regexp.MustCompile(`\b(?:\d{10}|\(\d{3}\)\s*\d{3}[\s-]?\d{4}|\d{3}[\s-]?\d{3}[\s-]?\d{4})\b`)

	// Read the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Find all phone numbers in the content
	phoneNumbers := phoneRegex.FindAllString(string(content), -1)

	// Print each phone number found
	for _, number := range phoneNumbers {
		fmt.Println(number)
	}
}
