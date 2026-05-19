package main

import "fmt"

//This is the main function, This function not receive any values and not return also.
//This function is the entry point to execution.
func main() {
	fmt.Println("Functions in the Golang")
	greeter() //Calling the created functin inside the main function

	//This is not allow the function inside the another function
	// func greetertwo(){
	// 	fmt.Println("I am another greeter two function")
	// }

	greetertwo()

	result := addition(2, 4)
	fmt.Println("Result is:", result)

	add, sub, mul, div := arthimetic(3, 3)
	fmt.Println("addition is", add, "subtract is", sub)
	fmt.Println("Multiplication is", mul, "Divition is", div)

	proresult, message := proAdder(2, 3, 4, 5, 6)
	fmt.Println(proresult, message)

}

//This is creation of the function
func greeter() {
	fmt.Println("Hello from the golang")
}

//Now same function declaring the out side the function
func greetertwo() {
	fmt.Println("I am another greeter two function")
}

//This function returning the value
func addition(a, b int) int {
	return a + b
}

//The function returning the so many values
//Smilarly it will return the different data types also
func arthimetic(a int, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// This function receive more value as a parameter
// This function is called Variadic Function (Three tripple ... dot's).
func proAdder(nums ...int) (int, string) {
	total := 0
	for _, val := range nums {
		total = total + val
	}
	return total, "hi, this is VARIADIC function"
}
