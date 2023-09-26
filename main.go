package main

import (
	"errors"
	"fmt"
	f "fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		f.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	f.Fprintf(w, "Post request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	f.Fprintf(w, "Name= %s\n", name)
	f.Fprintf(w, "Address= %s\n", address)
}  

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/main" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	f.Fprintf(w, "Hello!!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/main", mainHandler)

	f.Printf("Starting server on port 8000\n")
	// if err := http.ListenAndServe(":8000", nil); err != nil {
	// 	log.Fatal(err)
	// }
	err := http.ListenAndServe(":8000", nil)
	if errors.Is(err, http.ErrServerClosed){ //checks if the server was shutdown
		fmt.Printf("Server has been shutdown")
	}else if err != nil {
		fmt.Printf("Error starting server")
		log.Fatal(err)
	}
}
