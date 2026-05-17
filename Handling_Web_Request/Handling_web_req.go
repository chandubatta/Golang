package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("welcome to golang web req handling")
	const url = "google.com"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data_bytes, _ := io.ReadAll(response.Body)
	content := string(data_bytes)
	fmt.Println(content)
}
