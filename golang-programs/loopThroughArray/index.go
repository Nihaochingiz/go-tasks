package main

import "fmt"


func main() {
	intArray := [5]int{10,20,30,40,50}

	fmt.Println("---------------Example 1--------------------")

	for i := 0; i < len (intArray); i++ {
		fmt.Println(intArray[i])
	}
}