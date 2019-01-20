package api

import (
	"github.com/jmoiron/sqlx"
)

var ()

//ServerState holds the global variables inside the package such as the db pointer
type ServerState struct {
	DB *sqlx.DB
}
