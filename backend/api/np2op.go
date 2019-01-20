package api

import (
	"github.com/jmoiron/sqlx"
)

//ServerState holds the global variables inside the package such as the db pointer
type ServerState struct {
	DB *sqlx.DB
}

//Init initialises the package and should be called at start
func Init() {

}
