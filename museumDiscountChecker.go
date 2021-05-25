// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program checks whether the user will get a discount at the museum

package main

import "fmt"

func main() {
	// This function checks whether the user will get a discount at the museum
	var userDay string
	var userAge int

	// input
	fmt.Println("This program checks whether the user will get a discount at the museum")
	fmt.Println()
	fmt.Print("Enter the day you are going: ")
	fmt.Scanln(&userDay)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&userAge)
	fmt.Println()

	// process
	if (userDay == "tuesday" || userDay == "thursday") || (userAge > 12 && userAge < 21) {
		// output
		fmt.Println("You will get a discount")
	} else {
		fmt.Println("You will pay full price")
	}
}
