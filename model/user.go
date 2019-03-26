package model

import "log"

type User struct {
	Name string `json:"name"`
}

func (u *User) create(user User) (id int, err error) {
	log.Printf("create user: %+v \n", user)

}
