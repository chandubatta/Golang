#Defer

- In the reverse order they were deferred.

***LIFO (Last In FIrst Out)***

- First execute with out defer functions
- After complete the with out Defer functions execution
- Start the execution of the with defer functions, but working in order of LIFO (Last In   First Out)
- Defer keyword used functions are the working like a **STACK**

***CODE***
```func main() {
	//fmt.Println("Hello")            //Hello
	//defer fmt.Println("World")       //World

	defer fmt.Println("World") //Hello
	fmt.Println("Hello")       //World

	fmt.Println("World") //World
	fmt.Println("Hello") //Hello

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("hello")

	//Hello
	//World
	//Hello
	//hello
	//three
	//two
	//one
	//World

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}


output
--------

Hello
World
Hello
hello
4
3
2
1
0
three
two
one
World