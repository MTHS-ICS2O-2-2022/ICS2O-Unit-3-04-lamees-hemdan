// Created By: Lamees Hemdan
// Created On: April 2023
//
// This file contains the main function for the go_app application.

// This function converts temperature from Fahrenheit to Celsius
package main

import (
	"fmt"
)

func main() {

	var fahrenheit float64
	// input
	fmt.Print("Enter the temperature in Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	// process
	celsius := (fahrenheit - 32) * 5 / 9

	// output
	fmt.Println("The temperature in Celsius is: ", celsius)

	fmt.Println("\nDone.")
}
