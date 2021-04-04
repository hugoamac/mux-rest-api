package contact

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ContactEntity - This entity provides the data model for the contact.
type ContactEntity struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	LastName string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
}

//String - This method provides the string format for ContactEntity.
func (c *ContactEntity) String() string {
	return fmt.Sprintf("ID:%s,Name:%s,LastName:%s,Email:%s,Phone:%s", c.ID, c.Name, c.LastName, c.Email, c.Phone)
}

//FullName - This method provides the fullname of contact.
func (c *ContactEntity) FullName() string {
	return fmt.Sprintf("%s %s", c.Name, c.LastName)
}
