package main

import "fmt"

func main() {
	fmt.Println("Loops in the Golang")

	days := []string{"monday", "sunday", "tuesday", "wednesday", "satursday"}
	fmt.Println(days)

	//Only one for loop in golang
	for d := 0; d <= len(days); d++ {
		fmt.Println(days[d])
	}

	//another way of loop
	for i := range days {
		fmt.Println(days[i])
	}

	//another way of loop
	for indexx, value := range days {
		fmt.Printf("index is: %v and value is: %v \n", indexx, value)
	}
}
