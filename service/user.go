package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"textgopher/models"
)

type User interface {
	Insert(data interface{}, collectionName string) (*mongo.InsertOneResult, error)
	FindWithEmail(email string) (models.User, error)
}
