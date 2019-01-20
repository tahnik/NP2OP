package api

import (
	"github.com/jmoiron/sqlx"
)

var ()

type (
	//ServerState holds the global variables inside the package such as the db pointer
	ServerState struct {
		DB *sqlx.DB
	}

	//Post is an instance of a user's Post
	Post struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		Created     string `json:"created"`
		Description string `json:"description"`
		Goal        int    `json:"goal"`
	}
)
