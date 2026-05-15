package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func random() {
	fmt.Println(" Swith case in Golang")
	//rand.seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of the dice is :", diceNumber)

	//rand.Seed is deprecated: As of Go 1.20 there is no reason to call Seed with
	//a random value. Programs that call Seed with a known value to get
	//a specific sequence of results should use New(NewSource(seed)) to
	//obtain a local random generator.deprecateddefault

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}
}
