package handler

import (
	"encoding/json"
	"mux-rest-api/domain/contact"
	"mux-rest-api/domain/contact/usecase"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

//CreateContactHandler - This interface provides the object with the responsibility to handle the request to create a new contact.
type CreateContactHandler interface {
	HandlerRoute
}

//createContactHandler - This struct provides the implementation of CreateContactHandler interface.
type createContactHandler struct {
	db *mongo.Database
}

//NewCreateContactHandler - This method provides the instance of CreateContactHandler interface.
func NewCreateContactHandler(db *mongo.Database) CreateContactHandler {

	return &createContactHandler{db}
}

//Props - This method provides the path and http verb of handler CreateContactHandler.
func (h *createContactHandler) Props() (path string, method string) {
	path = "/api/v1/contact"
	method = "POST"

	return path, method
}

//Handler - This method provides the handling of the request to create a new contact.
func (h *createContactHandler) Handler(w http.ResponseWriter, r *http.Request) {

	contactRepository := contact.NewContactRepository(h.db)
	contactService := contact.NewContactService(contactRepository)
	createContactUseCase := usecase.NewCreateContactUseCase(contactService)

	var dataRequest *usecase.CreateContactDTORequest

	err := json.NewDecoder(r.Body).Decode(&dataRequest)

	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	dataResponse, err := createContactUseCase.Execute(dataRequest)

	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	ResponseJSON(w, http.StatusOK, dataResponse)
}
