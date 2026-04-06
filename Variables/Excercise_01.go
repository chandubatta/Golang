/*
Exercise 1 — Unit Converter
Write a Go program that stores a temperature in Celsius as a variable and converts it to Fahrenheit,
 storing the result in a second variable. Print both values with clear labels.
*/

package main

import "fmt"

func main() {
	// Step 1: Store the temperature in Celsius
	var celsius float64 = 100.0

	// Step 2: Convert Celsius to Fahrenheit using the formula
	fahrenheit := (celsius * 9 / 5) + 32

	// Step 3: Print both values
	fmt.Printf("Celsius:    %.2f°C\n", celsius)
	fmt.Printf("Fahrenheit: %.2f°F\n", fahrenheit)
}
