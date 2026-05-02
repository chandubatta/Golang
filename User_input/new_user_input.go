package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our piza app")
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating :")

	//comma ok or // err ok syntax
	ratting, _ := r.ReadString('\n')
	fmt.Printf("Thanks for rating %v and type %T \n", ratting, ratting)
	//fmt.Printf("Type %T", ratting)

	//Type Conversion from one data type to another
	// As of now we read the String
	Convert_numrating, _ := strconv.ParseFloat(strings.TrimSpace(ratting), 64)
	fmt.Printf("Thanks for rating %v and type %T \n", Convert_numrating, Convert_numrating)

}
