package main

import (
	"fmt"

	"NP2OP/backend/api"
)

func main() {
	fmt.Println("we in bois")
	r := api.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
