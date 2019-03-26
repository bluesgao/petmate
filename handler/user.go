package handler

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

type CreateResponse struct {
	Username string `json:"username"`
}

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		log.Printf("user create bind err:%+v \n", err)
		common.SendResult(c, common.BizErrBind, nil)
		return
	}

	user := model.User{
		Username: r.Username,
		Password: r.Password,
	}

	if err := user.Create(); err != nil {
		log.Printf("user create err:%+v \n", err)
		common.SendResult(c, common.BizErrDB, nil)
		return
	}
	common.SendResult(c, common.BizOk, nil)
}
