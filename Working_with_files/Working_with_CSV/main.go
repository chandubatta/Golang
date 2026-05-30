package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Working with CSV file")
	file, err := os.OpenFile("students.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fp := csv.NewReader(file)
	data, err := fp.ReadAll()
	if err != nil {
		fmt.Println("Error while reading the file")
	}
	for _, line := range data {
		fmt.Println(line)
	}
	defer file.Close()
}
