package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to User input in the golang")
	var (
		intvalue     int
		uintvalue    uint
		floatvalue   float32
		stringvalue  string
		boolvalue    bool
		complexvalue complex64
	)
	fmt.Println("Enter interger value")
	if _, err := fmt.Scan(&intvalue); err != nil {
		fmt.Println("Error while scaning int value", err)
		return
	}
	if _, err := fmt.Scan(&uintvalue); err != nil {
		fmt.Println("Error while scaning the uint vale", err)
		return
	}
	if _, err := fmt.Scan(&floatvalue); err != nil {
		fmt.Println("error while scaning the float value", err)
		return
	}
	if _, err := fmt.Scan(&stringvalue); err != nil {
		fmt.Println("Error while scaning the string value")
	}
}
