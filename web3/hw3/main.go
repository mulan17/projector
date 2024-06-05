// Розробити гру-текстовий квест «Новий світ».

// Ваш персонаж прокидається в невідомому місці з деякими речами. Він нічого не памʼятає.
// У нього є можливість піти одним з кількох шляхів (усі перелічені сутності — структури).
// Ситуація розвивається залежно від обраного рішення.

// Ігровий режим: текстом пишеться ситуація і пропонуються текстові варіанти, які може обрати гравець.
// Гравець пише один з варіантів і читає, як у цьому випадку розвивається ситуація.

// Можливий сценарій: Стівен прокинувся біля входу в печеру. Він лише памʼятає своє імʼя.
// Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж. У печері темно, тому Стівен іде стежкою,
// яка веде від печери в ліс. У лісі Стівен натикається на мертве тіло дивної тварини.
// Він обирає нічого з цим не робити й іти далі. Через деякий час Стівен приходить до безлюдного табору.
// Він вже втомлений і вирішує відпочити, а не йти далі. У найближчому наметі він знаходить сейф з кодовим замком з двох чисел.
// Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає.
// Стівен непритомніє. А все могло бути зовсім інакше.

// Треба використати Scan, type struct, if/else, switch/case, for loop.

package main

import (
	"fmt"
)

type Person struct {
	name string
}

type Bag struct {
	items map[string]bool
}

func (b *Bag) hasItem(item string) bool {
	return b.items[item]
}

func (b *Bag) useItem(item string) bool {
	if b.hasItem(item) {
		b.items[item] = false
		return true
	}
	return false
}

func (p Person) wakeUp(b Bag) {
	fmt.Printf("%v woke up in a cave. He remembers nothing. He has a bag with some items.\n", p.name)
}

func (p Person) choose(question string) bool {
	var answer bool
	for {
		fmt.Println(question)
		_, err := fmt.Scan(&answer)
		if err == nil {
			return answer
		}
	}
}

func (p Person) moveForward(b *Bag) bool {
	if p.choose("Do you want to go out of the cave? (true or false)") {
		if p.choose("Do you need a lighter for this? (true or false)") {
			if b.useItem("lighter") {
				fmt.Printf("Great. Now %v will try to find his home\n", p.name)
				return true
			} else {
				fmt.Printf("%v doesn't have a lighter. Game over:(\n", p.name)
				return false
			}
		}
		fmt.Printf("%v can't find his way in the dark:( Game over:(\n", p.name)
		return false
	} else {
		fmt.Printf("%v will never get home. Game is over:(\n", p.name)
		return false
	}
}

func (p Person) deadAnimal(b *Bag) bool {
	if p.choose("On his road Steven sees a super strange animal. Do you wanna touch it? (true or false)") {
		if p.choose("Do you need a knife for this? (true or false)") {
			if b.useItem("knife") {
				fmt.Printf("The animal was poisonous. %v will never get home. Game is over:(\n", p.name)
				return false
			} else {
				fmt.Printf("%v doesn't have a knife. He avoids touching the animal and moves on.\n", p.name)
				return true
			}
		}
	} else {
		fmt.Println("Great choice. It was a poisonous animal. Let's go further")
		return true
	}
	return true
}

func (p Person) resting(b *Bag) bool {
	if p.choose("You look tired. I see an empty camping. Do you wanna take a rest? (true or false)") {
		fmt.Println("Ok, let's take a rest. Here is super dark and cold...")
		if p.choose("Do you wanna take matches to make fire? (true or false)") {
			if b.useItem("match") {
				fmt.Println("Now it's better... Look, I see something")
				if p.choose("I see a safe. Do you wanna open it? (true or false)") {
					fmt.Printf("Oh no... %v was bitten by a big insect and it was poisonous. Game is over:(\n", p.name)
					return false
				}
			} else {
				fmt.Printf("%v doesn't have matches. It's too cold to rest here. Game over:(\n", p.name)
				return false
			}
		}
	} else {
		fmt.Printf("It was a long road, but %v gets home! Congratulations!\n", p.name)
		return true
	}
	return true
}

func main() {
	for {
		fmt.Println("Game: New World!")

		p := Person{
			name: "Steven",
		}

		b := Bag{
			items: map[string]bool{
				"match":   true,
				"knife":   true,
				"lighter": true,
			},
		}

		p.wakeUp(b)
		if !p.moveForward(&b) {
			if !p.choose("Do you want to start over? (true or false)") {
				break
			}
			continue
		}
		if !p.deadAnimal(&b) {
			if !p.choose("Do you want to start over? (true or false)") {
				break
			}
			continue
		}
		if !p.resting(&b) {
			if !p.choose("Do you want to start over? (true or false)") {
				break
			}
			continue
		}
		if !p.choose("Do you want to play again? (true or false)") {
			break
		}
	}
}
