package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, ዋኤል")
	})

	fmt.Println("Server running on port 8282")
	if err := http.ListenAndServe(":8282", nil); err != nil {
		log.Fatal(err)
	}

}
