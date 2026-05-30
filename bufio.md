#What is bufio?

*The bufio package provides buffered I/O (Input/Output) in Go.*

Normally, reading or writing data directly from files, network connections, or standard input can be inefficient because every operation may interact with the operating system.

bufio improves performance by using a memory buffer:

*Buffered Reader (bufio.Reader)* → Reads larger chunks into memory and serves smaller reads from the buffer.
*Buffered Writer (bufio.Writer)* → Collects writes in memory and sends them in larger chunks.
Common Uses
Reading user input from the terminal.
Reading files line by line.
Processing large text files.
Network programming (TCP/HTTP servers).
Efficient writing to files.
2. Simple Example
Reading User Input
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hello,", name)
}
Output
Enter your name: Chandu
Hello, Chandu
Explanation
reader := bufio.NewReader(os.Stdin)

Creates a buffered reader for keyboard input.

reader.ReadString('\n')

Reads characters until a newline (\n) is found.

Useful Methods
ReadString()
text, err := reader.ReadString('\n')

Reads until the delimiter.

ReadBytes()
data, err := reader.ReadBytes('\n')

Returns []byte.

ReadLine()
line, isPrefix, err := reader.ReadLine()

Reads a single line.

Scanner
scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}

Reads input token by token (usually line by line).

3. Common Beginner Mistakes
Mistake 1: Forgetting to Handle Errors

❌ Bad

text, _ := reader.ReadString('\n')

If reading fails, you'll never know.

✅ Good

text, err := reader.ReadString('\n')
if err != nil {
	fmt.Println(err)
}
Mistake 2: Forgetting to Flush a Writer

❌ Bad

writer := bufio.NewWriter(file)
writer.WriteString("Hello")

Data may remain in memory and never reach the file.

✅ Good

writer := bufio.NewWriter(file)
writer.WriteString("Hello")
writer.Flush()
Mistake 3: Using Scanner for Huge Lines

❌ Problem

scanner := bufio.NewScanner(file)

Scanner has a default token size limit (about 64 KB).

For very large lines, use:

reader := bufio.NewReader(file)

or increase scanner buffer size:

scanner.Buffer(make([]byte, 1024), 1024*1024)
4. Real-World Applications
1. Reading Log Files

Imagine a server log:

User Login
Payment Success
Order Created

You can process millions of lines efficiently:

scanner := bufio.NewScanner(file)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}

Used in:

Log analyzers
Monitoring tools
Security auditing systems
2. Building a TCP Chat Server

When a client sends messages:

connReader := bufio.NewReader(conn)

msg, _ := connReader.ReadString('\n')

Used in:

Chat applications
Game servers
Microservices communication
Network protocols