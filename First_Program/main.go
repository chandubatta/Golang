package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	//"log"
)

func main() {
	fmt.Print("welcome to golang")
	gretter()
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8080",r))
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Println("error starting the server: ", err)
	}

	/*
		->In the above code previos I use dthe port number 4000, Its not worked terminal stucked. because of 4000 already using in the our local system
		that means on the 4000 server already another program runnung.
		-> better to use error handling other than the log package

		-> 4000,3000,5000 sometimes used other apps
		-> use 8080, 8081, 9000 Usually safe to use

		-> To know which program running in the 4000 port to run this command
		   (netstat -ano | findstr :4000)

	*/
}
func gretter() {
	fmt.Println("Hello I am the gretter function")
}
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Golang</h1>"))
}
