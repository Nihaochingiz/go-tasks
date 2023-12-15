package main 

import "fmt"

func main() {
	var command = "go inside"
	if command == "go east" {
		fmt.Println("You head further up the mountain. ")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didnt quite get that")
	}
}