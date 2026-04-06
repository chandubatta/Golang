package main

import "fmt"

func main() {
	var r rune = 'A'
	fmt.Println(r)
	fmt.Println(string(r))

	str := "GO😊"
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println("\U0001F60A") //for this symbol
	s:= "Hello, Chandu"
	for i, v:= range s{
		fmt.Println("%d %U %c",)
	}
	/*
		length of the str is 6
			✔ Why 6?

		G → 1 byte
		o → 1 byte
		😊 → 4 bytes
	*/
}

// working with RUNE data type
/*
 1. What is a rune?

A rune in Go is an alias for int32 and represents a Unicode code point.

Why it exists:
Go strings are UTF-8 encoded, meaning characters can be 1 to 4 bytes long.
A rune lets you work with actual characters instead of raw bytes.
When to use it:
When handling non-ASCII text (e.g., emojis, accented characters, Asian scripts)
When iterating over characters safely
When performing text processing, validation, or transformation

👉 Think of it like this:

byte → raw data (1 byte)
rune → a full character (Unicode-aware)
🔹 What is rune in Go?

👉 rune is just an alias for:

int32

✔ It stores Unicode characters
✔ Size = 4 bytes (32 bits)

    | Feature  | byte   | rune           |
    | -------- | ------ | -------------- |
    | Type     | uint8  | int32          |
    | Size     | 1 byte | 4 bytes        |
    | Supports | ASCII  | Unicode        |
    | Example  | 'A'    | 'A', '😊', '你' |

	🔹 Why do we need rune?

Because byte is not enough for all characters.

👉 byte (uint8):

1 byte → supports only ASCII (0–255)

👉 rune:

4 bytes → supports Unicode (all languages + emojis)

// select emoji in windows
🔹 1. Windows Shortcut (Best way 💻)

👉 Press:

Windows key + .

OR

Windows key + ;

✔ This opens emoji picker

👉 Then:

Search “smile”
Click 😊
*/
