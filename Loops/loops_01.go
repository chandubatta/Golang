package main

import "fmt"

func main() {
	fmt.Println("Loops in the Golang")

	days := []string{"monday", "sunday", "tuesday", "wednesday", "satursday"}
	fmt.Println(days)

	//Only one for loop in golang
	for d := 0; d < len(days); d++ {
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

	// ,ok syntax type
	//another way of loop
	for _, value := range days {
		fmt.Printf("value is: %v \n", value)
	}

	//Like while loop
	chandu := 1

	for chandu < 10 {

		if chandu == 2 {
			goto div
		}
		if chandu == 5 {
			//break
			chandu++
			continue
		}
		fmt.Printf("value is %v \n", chandu)
		chandu++
	}

	// goto statement
	//label creating, now we can call this label by using the goto statement
div:
	fmt.Println("good person")
}
