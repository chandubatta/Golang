package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in the Golang")
	language := make(map[string]string)
	language["js"] = "javascript"
	language["rj"] = "reactjs"
	language["dsa"] = "Data Structures and Algorithms"
	language["py"] = "python"
	fmt.Println("list of all lamguages", language)
}
