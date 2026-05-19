package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello")            //Hello
	//defer fmt.Println("World")       //World

	defer fmt.Println("World") //Hello
	fmt.Println("Hello")       //World

	fmt.Println("World") //World
	fmt.Println("Hello") //Hello

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("hello")

	//Hello
	//World
	//Hello
	//hello
	//three
	//two
	//one
	//World

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}

}

// STACK [world,one,two,three,0,1,2,3,4]
// When ever we used the DEFER that time all DEFER values are stored into the STACK
// After stored into the STACK, Execute like a LIFO
//Print like a reverse order
