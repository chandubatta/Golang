package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Time study in the Golang")
	Present_time := time.Now()
	fmt.Println("Present Time is :", Present_time)
	fmt.Println(Present_time.Format("01-02-2006 15:04:05 Monday")) // Tihs is the official format of the time in the Golang

	create_date := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC) //create a date with own format
	fmt.Println(create_date)
	fmt.Println(create_date.Format("01-02-2006 Monday"))
}
