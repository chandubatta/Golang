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
	fmt.Printf("%T \n", language)

	//Printing Particular value
	fmt.Println("dsa short for:", language["dsa"])
	//Delete some values
	delete(language, "rj")
	fmt.Println(language)

	//Usage of the loop in MAP
	for key, val := range language {
		fmt.Println("key value is:", key, "value is:", val)
		//Another way is of printing
		fmt.Printf("For key %v, Value is %v \n", key, val)
	}

}
