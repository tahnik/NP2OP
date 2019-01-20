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
		ID          int    `db:"id"`
		Created     string `db:"timestamp"`
		Description string `db:"description"`
		Goal        int    `json:"goal"`
		CampaignID  int    `db:"campaignid"`
	}
)
