package main

import "fmt"

//Creating Structure
type User struct { // Use Capital letters first letter
	Name   string
	Age    int
	Email  string
	Status bool
	Oneage int
}

//Method or function creation
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

//
func (u User) NewMail() {
	u.Email = "div@gmail.com"
	fmt.Println("Email of this user is : ", u.Email)
}

func main() {
	fmt.Println("Structs in the golang")
	//No inheritance in golang; No Super or parent
	chandu := User{"Chandu", 26, "chandu@gmail.com", true, 16}
	fmt.Println(chandu)
	//{Chandu 26 chandu@gmail.com true}

	fmt.Printf("Chandu details %+v \n", chandu)
	//Chandu details {Name:Chandu Age:26 Email:chandu@gmail.com Status:true}

	fmt.Printf("Name is :%v, age is: %v \n", chandu.Name, chandu.Email)
	//Name is :Chandu, age is: chandu@gmail.com

	chandu.GetStatus()
	chandu.NewMail()
	//when I run this function the new email is printing
	//Email of this user is :  div@gmail.com

	// After call the NewMail() I am printing again the actual values
	//Name is :Chandu, age is: chandu@gmail.com //old Email only printing not printing new Email
	fmt.Printf("Name is :%v, age is: %v \n", chandu.Name, chandu.Email)
	//when ever we pass this like means 'func (u User) NewMail() {' we pass as a copy
	// when ever you Pass this objects or structs 'u', it actually passes on a copy.
	// This is the reason to why pointers are designed.

	// If you want to pass on the original object you should be passing up reference of it or a pointer to that
}
