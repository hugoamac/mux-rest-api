package infra

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect - This method provides the instance of mongo client connect
func Connect()*mongo.Client{

	mongo_uri:= viper.Get("MONGO_URI").(string)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil{
		log.Fatal("Error on connected mongo db")
	}
	log.Println("Mongo connected successfully")
	return client
}