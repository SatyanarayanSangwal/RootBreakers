package config

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func LoadConfig() {
	viper.SetConfigFile(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("MONGO_URI")))
	if err != nil {
		log.Fatal("Mongo Client Error: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Mongo Connection Error: ", err)
	}

	DB = client.Database(viper.GetString("DB_NAME"))
	log.Println("Connected to MongoDB 🚀")
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
