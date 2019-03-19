package handler

import "go.mongodb.org/mongo-driver/mongo"

type Handler struct {
	DbClient *mongo.Client
}
