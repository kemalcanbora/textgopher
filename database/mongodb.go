package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	c "textgopher/config"
	"textgopher/models"
	"time"
)

type MongoClient struct {
	Client  *mongo.Client
	Context context.Context
	Cancel  func()
}

func Connection() *MongoClient {
	var err error

	credential := options.Credential{
		Username: c.Configure().MongoUserName,
		Password: c.Configure().MongoPassword,
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(c.Configure().MongoURL).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	mongoClient := &MongoClient{
		Client:  client,
		Context: ctx,
		Cancel:  cancel,
	}
	return mongoClient
}

func (m *MongoClient) Insert(data interface{}, collectionName string) (*mongo.InsertOneResult, error) {
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (m *MongoClient) FindWithEmail(email string) (models.User, error) {
	var dbUser models.User
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(c.Configure().MongoUserCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&dbUser)
	if err != nil {
		log.Println(err)
	}
	return dbUser, err
}
