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

import "fmt"

type Person struct {
	name string
}

type Bag struct {
	contain1 string
	contain2 string
	contain3 string
}

func (p Person) wakeUp(b Bag) {
	fmt.Printf("%v woke up in a cave. He remembers nothing. He has the bag with %v, %v, %v. Help him get home.\n", p.name, b.contain1, b.contain2, b.contain3)
}

func (p Person) choose(question string) bool {
	var answer bool
	for {
		fmt.Println(question)
		_, err := fmt.Scan(&answer)
		if err == nil {
			return answer
		}
		// fmt.Println("Choose: 'true' or 'false'")
	}
}

func (p Person) moveForward() bool {
	if p.choose("Do you want to go out of the cave? (true or false)") {
		if p.choose("Do you need lighter for this?") {
			fmt.Printf("Great. Now %v will try to find his home\n", p.name)
			return true
		}
		fmt.Printf("%v can't find his way in the dark:( Game over:(", p.name)
		return false
	} else {
		fmt.Printf("%v will never get home. Game is over:(\n", p.name)
		return false
	}
}

func (p Person) deadAnimal(b Bag) bool {
	if p.choose("On his road Stiven sees a super strange animal on his road. Do you wanna touch it? (true or false)") {
		if p.choose("Do you need knife for this? (true or false)") {
			fmt.Printf("Animal was poison. %v will never get home. Game is over:(\n", p.name)
			return false
		}
	} else {
		fmt.Println("Greate choose. It was poison animal. Let's go futher")
		return true
	}
	return true
}

func (p Person) resting(b Bag) bool {
	if p.choose("You look tired. I see empty camping. Do you wanna take a rest?(true or false)") {
		fmt.Println("Ok, let's take a rest. Here is super dark and cold...")
		// if p.choose(fmt.Sprintf("Do you wanna take %s to make fire?"), b.contain1) {
		if p.choose("Do you wanna take match to make fire?") {
			fmt.Println("Now it's beter...Look I see something")
			if p.choose("I see some safe. Do you wanna open it? (true or false)") {
				fmt.Printf("Oh no...%v was bitten by a big insect and it was poisonous. Game is over:(", p.name)
				return false
			}
		}
	} else {
		fmt.Printf("It was long road, but %v get home! Congradulation!")
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
			contain1: "match",
			contain2: "knife",
			contain3: "lighter",
		}

		p.wakeUp(b)
		if !p.moveForward() {
			if !p.choose("Do you want to start over? (true or false)") {
				break
			}
			continue
		}
		if !p.deadAnimal(b) {
			if !p.choose("Do you want to start over? (true or false)") {
				break
			}
			continue
		}
		if !p.resting(b) {
			if !p.choose("Do you want to start over? (true or false)") {
				break
			}
			continue
		}
		fmt.Println("Do you want to play again? (true or false)")
		if !p.choose("Do you want to play again? (true or false)") {
			break
		}
	}
}

//comment