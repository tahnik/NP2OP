package main

import (
	"fmt"

	"NP2OP/backend/api"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	server   = "db"
	port     = "3306"
	user     = "root"
	password = "root"
	db       = "ht"

	testDB = user + ":" + password + "@tcp(localhost:3306)/" + db
)

var (
	prodDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, server, port, db)
)

func main() {
	db := sqlx.MustConnect("mysql", prodDB)

	server := api.ServerState{
		DB: db,
	}
	r := server.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
