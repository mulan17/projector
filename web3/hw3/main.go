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
	match bool
	knife bool
	ligther bool
}

func (p Person) wakeUp() {
	fmt.Printf("%v woke up in a cave. He remember nothing. Help him get home\n", p.name )
}

func (p Person) outFromCave(){
	var answear bool
	fmt.Println("Do you wanna go out of cave? ('true' or 'false')")

	_, err := fmt.Scan(&answear)
	if err != nil {
		fmt.Println("Choose: 'true' or 'false'")
		return
	}

	if answear == true {
		fmt.Printf("Great. Now %v will try to find his home\n", p.name)
	} else {
		fmt.Println("Game is over")
	}
}

func main(){
	println("Game: New Wordl!")

	p := Person{
		name: "Stive",
	}

	p.wakeUp()
	p.outFromCave()

}