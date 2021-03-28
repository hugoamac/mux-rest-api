package infra

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connection - This interface provides the connection database objection
type Connection interface{
	DB()*mongo.Database
	Close()
}

//conn - This struct type provides the connection from mongo.Database
type conn struct{
	db *mongo.Database
}

// Connect - This method provides the instance of mongo client connect
func NewConnection()Connection{

	mongo_uri:= viper.Get("MONGO_URI").(string)
	mongo_dbname:= viper.Get("MONGO_DBNAME").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil{
		log.Fatal("Error on connected mongo db")
	}

	db:=client.Database(mongo_dbname)

	log.Println("Mongo connected successfully")	
	return &conn{db}
}

//DB - This method provides the instance of mongo client database
func( c *conn)DB()*mongo.Database{
	return c.db
}
//Close -  This method provides the disconnect from mongo database
func( c *conn)Close(){

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Mongo disconnected successfully")	
	c.db.Client().Disconnect(ctx)	
}
