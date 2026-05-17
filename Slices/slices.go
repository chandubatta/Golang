package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in the Golang")
	// Creating the Slice
	var fruitlist = []string{"apple", "mango", "gova"}
	fmt.Printf("Type of the fruit list is %T \n", fruitlist)
	//Type of the fruit list is []string

	//Using append() method
	fruitlist = append(fruitlist, "banana", "grapes")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3]) //[mango gova]
	fmt.Println(fruitlist)

	//Using the MAKE key word to create Slice
	highscore := make([]int, 4)
	highscore[0] = 343
	highscore[1] = 345
	highscore[2] = 444
	highscore[3] = 555
	//highscore[4] = 999 //Error Index out of range[4]with length[4]

	highscore = append(highscore, 9090, 8080, 7070) //By using the append we can add
	fmt.Println(highscore)

	fmt.Println(sort.IntsAreSorted(highscore)) //checking sorted are not, taking bool value

	sort.Ints(highscore) //Sorting
	fmt.Println(highscore)

	//how to remove a value from the Slice based on the index

	var courses = []string{"python", "go", "JavaScript", "java", "React", "Angular"}
	fmt.Println(courses)
	var index = 3
	courses = append(courses[:index], courses[index+1:]...) // remove the value
	fmt.Println(courses) //After remove index 3 [python go JavaScript React Angular]
}
