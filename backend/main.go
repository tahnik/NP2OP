package main

import (
	"fmt"

	"NP2OP/backend/api"

	"github.com/jmoiron/sqlx"
)

func main() {
	fmt.Println("we in bois")
	db := sqlx.MustConnect("mysql", "user=root dbname=ht password=root sslmode=disable")

	server := api.ServerState{
		DB: db,
	}
	r := server.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
