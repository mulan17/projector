// Створити програму для симуляції групи людей, які одночасно грають в ігри на великому екрані (моделюємо гру схожу на 
// kahoot https://www.youtube.com/watch?v=az1xm2Ij7rA). Програма має використовувати горутину-генератор, який кожні 10 секунд 
// генерує новий ігровий раунд (питання та варіанти відповідей) та відправляє його до горутин-гравців через канал. 
// Гравці отримують новий ігровий раунд та вводять свої відповіді через окремий канал. Далі горутина-лічильник перевіряє 
// правильність відповідей та повертає результат (кількість відповідей по варіантах та/або загальний результат гри по гравцях) у 
// головну горутину через окремий канал, яка виводить результат раунду на екран. Якщо користувач перериває програму, то програма має 
// коректно завершувати роботу з використанням контексту.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type PlayerAnswer struct {
	PlayerID int
	Answer   int
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signals to gracefully shut down
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalCh
		cancel()
	}()

	questionCh := make(chan Question)
	answerCh := make(chan PlayerAnswer)
	resultCh := make(chan map[int]int)

	// Start goroutines
	go gameRoundGenerator(ctx, questionCh)
	go gameResultCounter(ctx, answerCh, resultCh)

	// Simulate players
	numPlayers := 5
	for i := 0; i < numPlayers; i++ {
		go player(ctx, i, questionCh, answerCh)
	}

	// Display results
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case result := <-resultCh:
				fmt.Println("Round results:", result)
			}
		}
	}()

	// Keep the main function running
	<-ctx.Done()
	fmt.Println("Shutting down...")
}

func gameRoundGenerator(ctx context.Context, questionCh chan<- Question) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	questions := []Question{
		{"What is 10+2?", []string{"10", "12", "14", "16"}, 2},
		{"What is the capital of Ukraine?", []string{"Donetsk", "Kyiv", "Paris", "Lviv"}, 2},
		{"What is the largest planet?", []string{"Earth", "Jupiter", "Mars", "Saturn"}, 1},
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			q := questions[rand.Intn(len(questions))]
			questionCh <- q
			fmt.Println("New question:", q.Text)
		}
	}
}

func player(ctx context.Context, id int, questionCh <-chan Question, answerCh chan<- PlayerAnswer) {
	for {
		select {
		case <-ctx.Done():
			return
		case q := <-questionCh:
			answer := rand.Intn(len(q.Options))
			fmt.Printf("Player %d answered %d\n", id, answer)
			answerCh <- PlayerAnswer{PlayerID: id, Answer: answer}
		}
	}
}

func gameResultCounter(ctx context.Context, answerCh <-chan PlayerAnswer, resultCh chan<- map[int]int) {
	results := make(map[int]int)

	for {
		select {
		case <-ctx.Done():
			return
		case answer := <-answerCh:
			results[answer.Answer]++
			if len(results) == 5 { // Assuming we have 5 players
				resultCh <- results
				results = make(map[int]int)
			}
		}
	}
}



