package api

import (
	"encoding/json"
)

type User struct {
	Id         int    `db:"id" json:"id"`
	Username   string `db:"username" json:"username"`
	Password   string `db:"password" json:"password"`
	Name       string `db:"name" json:"name"`
	Email      string `db:"email" json:"email"`
	Phone      string `db:"phone" json:"phone"`
	Country    string `db:"country" json:"country"`
	UserTypeId int    `db:"usertype_id" json:"usertype_id"`
}

func GetUserFromByes(b []byte) (u *User, err error) {
	if err = json.Unmarshal(b, &u); err != nil {
		return nil, err
	}
	return u, nil
}
