//Написати дві програми:
// 1. Яка створює 3 горутини. Перша горутина генерує випадкові числа й надсилає їх через канал у другу горутину. 
// Друга горутина отримує випадкові числа та знаходить їх середнє значення, після чого надсилає його в третю горутину через канал. 
// Третя горутина виводить середнє значення на екран.

// 2. Яка створює 2 горутини. Перша горутина генерує випадкові числа в заданому діапазоні й надсилає їх через канал у другу горутину.
//  Друга горутина отримує випадкові числа і знаходить найбільше й найменше число, після чого надсилає їх назад у першу горутину 
// через канал. Перша горутина виводить найбільше й найменше числа на екран.


package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numChannel := make(chan int)
	avgChannel := make(chan float64)
	done := make(chan struct{})

	var s sync.WaitGroup
	s.Add(3)

	go generateNumbers(&s, numChannel, done)
	go calculateAverage(&s, numChannel, avgChannel, done)
	go printAverage(&s, avgChannel)

	time.Sleep(5 * time.Second)
	close(done)
	s.Wait()
}

func generateNumbers(s *sync.WaitGroup, ch chan<- int, done <-chan struct{}) {
	defer s.Done()
	for {
		select {
		case <-done:
			close(ch)
			return
		default:
			num := rand.Intn(100)
			fmt.Println("Generated number:", num)
			ch <- num
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func calculateAverage(s *sync.WaitGroup, numCh <-chan int, avgCh chan<- float64, done <-chan struct{}) {
	defer s.Done()
	var sum, count int
	for {
		select {
		case <-done:
			close(avgCh)
			return
		case num, ok := <-numCh:
			if !ok {
				close(avgCh)
				return
			}
			sum += num
			count++
			avg := float64(sum) / float64(count)
			avgCh <- avg
		}
	}
}

func printAverage(s *sync.WaitGroup, ch <-chan float64) {
	defer s.Done()
	for avg := range ch {
		fmt.Println("Average:", avg)
	}
}
