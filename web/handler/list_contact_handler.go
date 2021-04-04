package handler

import (
	"mux-rest-api/domain/contact"
	"mux-rest-api/domain/contact/usecase"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

//ListContactHandler - This interface provides the object with the responsibility to handle the request to retrieve the contact list.
type ListContactHandler interface {
	HandlerRoute
}

//listContactHandler - This struct provides the implementation of ListContactHandler interface.
type listContactHandler struct {
	db *mongo.Database
}

//NewListContactHandler - This method provides the instance of ListContactHandler interface.
func NewListContactHandler(db *mongo.Database) ListContactHandler {

	return &listContactHandler{db}
}

//Props - This method provides the path and http verb of ListContactHandler.
func (h *listContactHandler) Props() (path string, method string) {
	path = "/api/v1/contact"
	method = "GET"

	return path, method
}

//Handler - This method provides the handling of the request to retrieve the contact list.
func (h *listContactHandler) Handler(w http.ResponseWriter, r *http.Request) {

	contactRepository := contact.NewContactRepository(h.db)
	contactService := contact.NewContactService(contactRepository)
	listContactUseCase := usecase.NewListContactUseCase(contactService)

	dataResponse, err := listContactUseCase.Execute()

	if err != nil {
		ResponseError(w, http.StatusNotFound, err.Error())
	}

	ResponseJSON(w, http.StatusOK, dataResponse)
}
