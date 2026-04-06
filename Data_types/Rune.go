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
