package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(writer, "POST request accepted successfuly\n")
	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(writer, "Name is %s\n", name)
	fmt.Fprintf(writer, "Address is %s\n", address)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 Not found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, request.Method+" method is not supported for this route", http.StatusNotFound)
		return
	}

	fmt.Fprintf(writer, "This is hello page")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
