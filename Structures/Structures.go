package main

import "fmt"

//Creating Structure
type User struct { // Use Capital letters first letter
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	fmt.Println("Structs in the golang")
	//No inheritance in golang; No Super or parent
	chandu := User{"Chandu", 26, "chandu@gmail.com", true}
	fmt.Println(chandu)
	//{Chandu 26 chandu@gmail.com true}

	fmt.Printf("Chandu details %+v \n", chandu)
	//Chandu details {Name:Chandu Age:26 Email:chandu@gmail.com Status:true}

	fmt.Printf("Name is :%v, age is: %v", chandu.Name, chandu.Email)
	//Name is :Chandu, age is: chandu@gmail.com
}
