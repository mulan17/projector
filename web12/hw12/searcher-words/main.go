package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

func main() {
	const filename = "/Users/anastasiya/Documents/projector/web12/hw12/texts/1689007676028_text.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read file")
	}

	wordRegex := regexp.MustCompile(`\p{L}+`)

	words := wordRegex.FindAllString(string(content), -1)

	fmt.Println("Words:")
	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}
}
