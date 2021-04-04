package infra

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connection - This interface provides the object with the responsibility to create a connection to the mongodb database.
type Connection interface {
	DB() *mongo.Database
	Close()
}

//conn - This struct provides the implementation for Connection interface.
type conn struct {
	db *mongo.Database
}

// Connect - This method provides the instance of Connection interface.
func NewConnection(config Config) Connection {

	mongo_uri := config.MongoUri
	mongo_dbname := config.MongoDbName

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal("Error on connected mongo db")
	}

	db := client.Database(mongo_dbname)

	log.Println("Mongo connected successfully")
	return &conn{db}
}

//DB - This method provides the instance of connection to the mongodb database.
func (c *conn) DB() *mongo.Database {
	return c.db
}

//Close - This method provides to close the connection to the mongodb database.
func (c *conn) Close() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Mongo disconnected successfully")
	c.db.Client().Disconnect(ctx)
}
