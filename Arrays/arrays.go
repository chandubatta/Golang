package main

import "fmt"

func main() {
	fmt.Println("Arrays in the Golang")

	// This is the Basic Declaration of the array
	var fruitlist [4]string
	// Add some data into the Array
	fruitlist[0] = "apple"
	fruitlist[1] = "mango"
	fruitlist[3] = "grapes"

	fmt.Println("The fruit list is:", fruitlist)
	fmt.Printf("Type of the fruit list is %T \n", fruitlist)
	//Type of the fruit list is [4]string

	//Another way of declaration of the array
	var veg_list = [3]string{"tomato", "beens"}
	fmt.Println("vegtable list is:", veg_list)
	fmt.Printf("Type of the veg list is %T \n", veg_list)
	//Type of the veg list is [4]string

	//Array declaration using the Make() function
	var size int
	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&size)
	arr := make([]string, size)
	fmt.Println("Enter the string values into the array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
}
