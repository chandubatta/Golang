package main

import "fmt"

func main() {
	fmt.Println("If Else Condition Controll Statements in Go")
	logincount := 23
	var result string
	if logincount < 10 {
		result = "Regular User"
	} else if logincount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly login count is 10"
	}
	fmt.Println(result)

	//another way
	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is odd")
	}

	//web request handling // variable creating and condition at a time in if condition
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is graten than 10")
	}

	//error handling time if condition usage

	//if err != nil {

	//}

}
