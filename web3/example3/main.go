package main

import "fmt"

type cup struct {
	size string
	isClean bool
}

func(c *cup) clean(){
	c.isClean = true
	fmt.Println("Cleaned cup")
}

type teaCup struct {
	cup
	tea string 
	containsWater bool
}

func (c *teaCup) addTea(tea string) {
	c.tea = tea
	fmt.Printf("Added %s tea\n", tea)
}

func (c *teaCup) addWater() {
	c.containsWater = true
	fmt.Println("Added water")
}

func main() {
	for i := 0; i < 6; i++ {
		if i < 3 {
			continue
		}
		
		fmt.Println(i)
	}
}


func createTeaCups() {
	var teaCounter = 0
	for true {
		if teaCounter ==5 {
			break
		}

		c := teaCup{
			cup: cup{
				size: "big",
			},
			
		}
	
		c.clean()
	
		tea := "black"
	
		c.addTea(tea)
	
		c.addWater()
	
		fmt.Printf("Got cup of tea %d: %+v\n", teaCounter, c)

		teaCounter++

	}

	fmt.Println("Finished")


}