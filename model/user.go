package model

import "log"

type User struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:password`
}

func (user *User) Create() error {
	log.Printf("create user: %+v \n", user)
	key := "u_" + user.Username

	fields := make(map[string]interface{})
	fields["username"] = user.Username
	fields["password"] = user.Password
	fields["userid"] = user.Userid

	if ret, err := RedisCli.HMSet(key, fields).Result(); err != nil {
		log.Printf("create user ret: %+v \n", ret)

	}

	return nil
}
