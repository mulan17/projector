package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numChannel := make(chan int)
	resultChannel := make(chan [2]int)

	go numberGenerator(numChannel)
	go findMinAndMax(numChannel, resultChannel)

	result := <-resultChannel
	fmt.Printf("Min: %d, Max: %d\n", result[0], result[1])
}

func numberGenerator(numChannel chan<- int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Local random generator
	for i := 0; i < 10; i++ {
		numChannel <- rng.Intn(100)
	}
	close(numChannel)
}

func findMinAndMax(numChannel <-chan int, resultChannel chan<- [2]int) {
	min := int(^uint(0) >> 1)
	max := -min - 1

	for num := range numChannel {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	resultChannel <- [2]int{min, max}
	close(resultChannel)
}
