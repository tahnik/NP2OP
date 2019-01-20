package api

import (
	"encoding/json"
)

type User struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	Phone      string `db:"phone"`
	Country    string `db:"country"`
	// UserTypeId int    `json:usertype_id`
}

func GetUserFromByes(b []byte) (u *User, err error) {
	if err = json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}