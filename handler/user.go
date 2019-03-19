package handler

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mongodb/mongo-go-driver/bson"
	"log"
	"net/http"
	"petmate/model"
	"time"
)

func (h *Handler) Register(reqCtx echo.Context) (err error) {
	u := new(model.User)
	log.Printf("Register reqCtx: %+v", reqCtx)

	if err1 := reqCtx.Bind(u); err1 != nil {
		return errors.New("参数解析错误")
	}

	if u.Name == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid name"}
	}

	collection := h.DbClient.Database("cowboy").Collection("article")
	defer h.DbClient.Disconnect(nil)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ret, err := collection.InsertOne(ctx, bson.M{"name": u.Name})
	log.Println("Register:", ret.InsertedID)
	return
}
