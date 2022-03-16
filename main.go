package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting go-world server...")
	http.HandleFunc("/", goServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func goServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "405")
}
