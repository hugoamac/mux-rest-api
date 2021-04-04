package routes

import (
	"log"
	"mux-rest-api/web/handler"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

//MyRoutes - This interface provides the object with the responsibility to configure the routes of the application.
type MyRoutes interface {
	Install(h handler.HandlerRoute)
}

//myRoutes - This struct provides the implementation for MyRoutes interface.
type myRoutes struct {
	db     *mongo.Database
	router *mux.Router
}

//NewMyRoutes - This method provides the instance of MyRoutes interface.
func NewMyRoutes(db *mongo.Database, router *mux.Router) MyRoutes {
	return &myRoutes{db, router}
}

//Install - This method provides the means to install the handlers.
func (r *myRoutes) Install(h handler.HandlerRoute) {

	path, method := h.Props()
	r.router.HandleFunc(path, h.Handler).Methods(method)
	log.Printf("%s %s", method, path)
}
