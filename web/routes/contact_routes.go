package routes

import (
	"mux-rest-api/web/handler"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

//ContactRoutes - This method provides the means to install all handlers for the contact domain.
func ContactRoutes(db *mongo.Database, router *mux.Router) {

	contactRoutes := NewMyRoutes(db, router)
	createContactHandler := handler.NewCreateContactHandler(db)
	listContactHandler := handler.NewListContactHandler(db)
	//handler for the contact domain
	contactRoutes.Install(createContactHandler)
	contactRoutes.Install(listContactHandler)
}
