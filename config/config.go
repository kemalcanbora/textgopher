package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AWSAccessKeyId      string `json:"aws_access_key_id"`
	AWSSecretAccessKey  string `json:"aws_secret_access_key"`
	AWSRegion           string `json:"aws_region"`
	MongoUserName       string `json:"mongo_username"`
	MongoPassword       string `json:"mongo_password"`
	MongoDatabase       string `json:"mongo_database"`
	MongoURL            string `json:"mongo_url"`
	MongoUserCollection string `json:"mongo_user_collection"`
}

func Configure() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var c Config
	c.AWSAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	c.AWSSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	c.AWSRegion = os.Getenv("AWS_REGION")
	c.MongoUserName = os.Getenv("MONGO_ROOT_USERNAME")
	c.MongoPassword = os.Getenv("MONGO_ROOT_PASSWORD")
	c.MongoDatabase = os.Getenv("MONGO_DATABASE")
	c.MongoURL = os.Getenv("MONGO_URL")
	c.MongoUserCollection = os.Getenv("MONGO_USER_COLLECTION")

	return c
}
