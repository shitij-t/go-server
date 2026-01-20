package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// create fileserver
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting servere on port: 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() Error:", err)
		return
	}

	fmt.Fprint(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}
