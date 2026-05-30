package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	string_value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading the string")
	}
	fmt.Println(string_value)
}
