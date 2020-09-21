package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello dude.")

	db := InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	Route(r, db)

	fmt.Println("Serving on port 8080!")
	http.ListenAndServe(":8080", r)
}
