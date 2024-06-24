package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numChannel := make(chan int)
	minMaxChannel := make(chan [2]int)

	go generateNumbersInRange(numChannel)
	go findMinMax(numChannel, minMaxChannel)

	time.Sleep(5 * time.Second)
}

func generateNumbersInRange(ch chan<- int) {
	for {
		num := rand.Intn(100)
		fmt.Println("Generated number:", num)
		ch <- num
		time.Sleep(500 * time.Millisecond)
	}
}

func findMinMax(numCh <-chan int, minMaxCh chan<- [2]int) {
	var min, max int
	first := true

	for num := range numCh {
		if first {
			min, max = num, num
			first = false
		} else {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		minMaxCh <- [2]int{min, max}
	}
}
