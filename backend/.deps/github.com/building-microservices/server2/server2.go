package server2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloResp struct {
	Id      int    `json:"id, string"`
	Message string `json:"msg"`
}

func Run(port int) {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		resp := HelloResp{Message: "This is the response", Id:5}
		anEncoder := json.NewEncoder(w)
		err := anEncoder.Encode(&resp)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("Serving on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
