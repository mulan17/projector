package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

func main() {
	const filename = "/Users/anastasiya/Downloads/1689007675141_numbers.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read file")
	}

	phoneRegex := regexp.MustCompile(`\b(\d{10}|\(\d{3}\)\s?\d{3}-\d{4}|\(\d{3}\)\d{3}-\d{4}|\d{3}-\d{3}-\d{4}|\d{3}\.\d{3}\.\d{4}|\d{3}\s\d{3}\s\d{4})\b`)

	phoneNumbers := phoneRegex.FindAllString(string(content), -1)

	fmt.Println("Phone numbers:")
	for i, number := range phoneNumbers {
		fmt.Printf("%d: %s\n", i+1, number)
	}
}
