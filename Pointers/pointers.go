package main

import "fmt"

func main() {
	fmt.Println("Pointers in the Golang")
	var name string

	//Creating Pointer type Variable
	var ptr *int
	fmt.Printf("value of pointer is: %v \n", ptr) //<nil>
	fmt.Println(ptr)                              //<nil>

	var ptr1 *string
	fmt.Printf("value of pointer is: %v \n", ptr1) //<nil>
	fmt.Println(ptr1)                              //<nil>

	mynumber := 23
	//Now creating the Pointer for already created variable using the '&' Symbol
	var ptr2 = &mynumber
	fmt.Println("Value of the ponter is :", ptr2)
	fmt.Println("Value of the pointer is:", *ptr2)
	*ptr2 = *ptr2 * 2
	fmt.Println("New value is: ", mynumber)
}
