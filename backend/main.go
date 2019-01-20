package main

import (
	"NP2OP/backend/api"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	server   = "db"
	port     = "3306"
	user     = "root"
	password = "root"
	db       = "ht"
)

func main() {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, server, port, db)
	db := sqlx.MustConnect("mysql", connectString)

	server := api.ServerState{
		DB: db,
	}
	r := server.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
