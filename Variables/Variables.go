package main

import "fmt"

func main() {
	// Explicit declaration with var
	var name string = "Gopher"

	// Type inferred by compiler
	var age = 5

	// Short variable declaration (most common inside functions)
	language := "Go"

	// Multiple variables at once
	var x, y int = 10, 20

	// Zero values — Go initializes variables automatically
	var isReady bool    // false
	var score int       // 0
	var username string // ""

	fmt.Printf("%s has been around for %d years\n", language, age)
	fmt.Printf("Hello, %s! x=%d, y=%d\n", name, x, y)
	fmt.Printf("Zero values — isReady: %v, score: %d, username: %q\n", isReady, score, username)

	// Default value for the variables
	var chandu int
	fmt.Printf("User iput value %v and type %T \n", chandu, chandu)
	var c1 float32
	fmt.Printf("User iput value %v and type %T \n", c1, c1)
	var c2 bool
	fmt.Printf("User iput value %v and type %T \n", c2, c2)
	var c3 string
	fmt.Printf("User iput value %v and type %T \n", c3, c3)
	var c4 byte
	fmt.Printf("User iput value %v and type %T \n", c4, c4)
	var c5 rune
	fmt.Printf("User iput value %v and type %T \n", c5, c5)
	var c6 uintptr
	fmt.Printf("User iput value %v and type %T \n", c6,c6)

}
