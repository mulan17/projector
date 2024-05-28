package main

import (
	"fmt"
	"math/rand"
)


func controlflow() {
	i := rand.Intn(100)

	fmt.Println(i)

	switch i{
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("50")
	default:
		fmt.Println("not 10,20,50")
	}

	switch j := i % 2; j {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("could not happen")
	
	}

	if j := i % 2; j == 0 {
		fmt.Println("j % 2 == 0")
	}
}



func ifelse(){
i := rand.Intn(10)
fmt.Println(i)

	 if i > 5 {
		fmt.Println("a")
	 } else {
		fmt.Println("b")
	 }

	 if i >= 5 {
		fmt.Println("a1")
	 } else {
		fmt.Println("b1")
	 }

	 if i/2 == 2 {
		fmt.Println("/ 2 = 2")
	 }

	 if is4(i) || isBigger2(i) {
		fmt.Println("4 and >2")
	}
	
}



func isBigger2 (i int) bool {
	// if i >2 {
	// 	return true 
	// } else {
	// 	return false
	// }
	return i > 2
}

func is4 (i int) bool {
	if i == 4 {
		return true 
	} else {
		return false
	}
}