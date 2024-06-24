package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numChannel := make(chan int)
	avgChannel := make(chan float64)

	go generateNumbers(numChannel)
	go calculateAverage(numChannel, avgChannel)
	go printAverage(avgChannel)

	time.Sleep(5 * time.Second)
}

func generateNumbers(ch chan<- int) {
	for {
		num := rand.Intn(100)
		fmt.Println("Generated number:", num)
		ch <- num
		time.Sleep(500 * time.Millisecond)
	}
}

func calculateAverage(numCh <-chan int, avgCh chan<- float64) {
	var sum, count int
	for num := range numCh {
		sum += num
		count++
		avg := float64(sum) / float64(count)
		avgCh <- avg
	}
}

func printAverage(ch <-chan float64) {
	for avg := range ch {
		fmt.Println("Average:", avg)
	}
}
