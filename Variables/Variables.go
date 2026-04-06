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
}
