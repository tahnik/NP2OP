package server1

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	port := 8081

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hey\nI'm a function\nlol")
	})

	log.Printf("Serving on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
