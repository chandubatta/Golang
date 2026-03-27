package main

import (
	"fmt"
	"os"
)

/*
1 bit = 0 or 1
8 bit = 10101100  = 1 byte
1024 bytes = 1 KB
1024 KB    = 1 MB
1024 MB    = 1 GB
1024 GB    = 1 TB
*/

func main() {
	var b byte = 65
	var k byte = 10
	fmt.Println(k)
	fmt.Println(string(k))
	fmt.Println(b)
	fmt.Println(string(b))
	/*
			Key points:
		32 → Space (first visible)
		48–57 → Digits (0–9)
		65–90 → Uppercase letters (A–Z)
		97–122 → Lowercase letters (a–z)
		Others → symbols (! @ # $ % ^ & * etc.)

	*/

	// This loop shows the All ASCII values from 0 to 127

	for i := 0; i <= 127; i++ {
		fmt.Printf("%d -> %q\n", i, i)
		// here we print i two times one time %d another time %q .
	}
	FileHandling()
	data_01, err := os.ReadFile("byte_data_type_01.txt")
	//after file read stored into the data_01
	if err != nil {
		fmt.Println(err)
	}
	Write_File(data_01)

}
func FileHandling() {
	data, err := os.ReadFile("byte_data_type.txt")
	if err != nil {
		fmt.Println("error while reading the file", err)
	} else {
		fmt.Println("file reading completed")
	}
	if data == nil {
		fmt.Println("Nothing to read inside the file")
	} else {
		fmt.Println("Content of the inside this file in binary format :", data)
		fmt.Println("Content of the inside this file in string format :", string(data))
	}

}
func Write_File(a []byte) {
	file, err := os.OpenFile("byte_data_type.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	/*
										0644
										││││
										│││└── Others
										││└─── Group
										│└──── Owner
										└───── (prefix)

								| Number | Meaning     |
								| ------ | ----------- |
								| 4      | Read (r)    |
								| 2      | Write (w)   |
								| 1      | Execute (x) |

								)
						🔹 Now decode 0644

						🔹 Why it starts with 0?
				       👉 The 0 means:
				            Number is in octal (base 8) format
				            Used in Unix/Linux permissions
						👉 Owner = 6
						6 = 4 + 2 = Read + Write
						👉 Group = 4
						4 = Read only
						👉 Others = 4
						4 = Read only

						🔹 Final Meaning
						Owner  → Read + Write
						Group  → Read only
						Others → Read only

						🔹 In Linux style
						-rw-r--r--

						| Value | Meaning                                    |
		                | ----- | ------------------------------------------ |
		                | 0777  | Everyone can read/write/execute (not safe) |
		                  0755  | Owner full, others read + execute          |
		                | 0600  | Only owner can read/write (secure 🔐)      |


	*/
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write([]byte("\n"))
	if err != nil {
		fmt.Println("error while writing", err)
	}

	_, err = file.Write(a)
	if err != nil {
		fmt.Println("error while writing", err)
	}

	_, err = file.Write([]byte("Hi chandu how are you"))
	if err != nil {
		fmt.Println("error while writing", err)
	}

}
