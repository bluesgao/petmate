package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"petmate/common"
	"petmate/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		log.Fatalf("user create bind err:%+v \n", err)
		common.SendResponse(c, common.RequstBindError, nil)
		return
	}

	user := model.User{
		Username: r.Username,
		Password: r.Password,
	}

	if err := user.Create(); err != nil {
		log.Printf("user create err:%+v \n", err)
		common.SendResponse(c, common.RedisError, nil)
		return
	}
	common.SendResponse(c, common.Ok, user.Username)
}
