package main

import "fmt"
func main(){
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(75.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(23 * 365 / 687)
	fmt.Print(" years old.")
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}