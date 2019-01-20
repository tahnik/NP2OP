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
		ID                    int    `db:"id"`
		Username              string `db:"username"`
		Description           string `db:"description"`
		Timestamp             string `db:"timestamp"`
		TotalCampaignSnapshot string `db:"total_campaign_snapshot"`
		CourseName            string `db:"name"`
		Cost                  int    `db:"cost"`
	}
)
