package main

import (
	"NP2OP/backend/api"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	user          = "root"
	password      = "root"
	db            = "ht"
	connectString = user + ":" + password + "@/" + db
)

func main() {
	db := sqlx.MustConnect("mysql", connectString)

	server := api.ServerState{
		DB: db,
	}
	r := server.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
