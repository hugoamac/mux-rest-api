package contact

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CollectionName - This constant provides the collection name
const (
	CollectionName = "contacts"
)

// ContactRepository - This interface provides the object with the responsibility to carry out the database transactions of the contact entity.
type ContactRepository interface {
	Create(c *ContactEntity) error
	FetchAll() ([]*ContactEntity, error)
	Find(id string) (*ContactEntity, error)
	Update(c *ContactEntity) (*ContactEntity, error)
	Delete(id string) error
}

// contactRepository - This struct provides the implementation for ContactRepository interface.
type contactRepository struct {
	collection *mongo.Collection
}

//NewContactRepository - This method provides the instance of ContactRepository interface.
func NewContactRepository(db *mongo.Database) ContactRepository {
	return &contactRepository{db.Collection(CollectionName)}
}

//Create -  This method provides the persistence of the contact entity in the database.
func (r *contactRepository) Create(c *ContactEntity) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.collection.InsertOne(ctx, &c)

	if err != nil {
		return err
	}

	c.ID = res.InsertedID.(primitive.ObjectID)

	return nil
}

//FetchAll - This method provides to retrieve the list of contacts in the database.
func (r *contactRepository) FetchAll() ([]*ContactEntity, error) {

	var contacts []*ContactEntity

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var c *ContactEntity
		cursor.Decode(&c)
		contacts = append(contacts, c)
	}

	return contacts, nil
}

//Find - This method provides to retrieve a contact in the database by the identifier.
func (r *contactRepository) Find(id string) (*ContactEntity, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var c *ContactEntity

	err = r.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

// Update - This method provides for updating a contact in the database.
func (r *contactRepository) Update(c *ContactEntity) (*ContactEntity, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{"_id", c.ID}}
	data := bson.D{{"$set", bson.D{
		{"name", c.Name},
		{"last_name", c.LastName},
		{"email", c.Email},
		{"phone", c.Phone},
	}}}

	res, err := r.collection.UpdateOne(ctx, filter, data)

	if err != nil {
		return nil, err
	}

	if res.ModifiedCount > 0 {
		return c, nil
	}

	return nil, errors.New("document not found")
}

//Delete - This method provides to remove a contact from the database by the identifier.
func (r *contactRepository) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}

	res, err := r.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount > 0 {
		return nil
	}

	return errors.New("document not found")
}
