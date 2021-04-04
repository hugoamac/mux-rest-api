package usecase

import (
	"mux-rest-api/domain/contact"
)

//CreateContactUseCase - This interface provides an object with the responsibility to create a new contact.
type CreateContactUseCase interface {
	Execute(request *CreateContactDTORequest) (*CreateContactDTOResponse, error)
}

//createContactUseCase - This struct provides the implementation of CreateContactUseCase interface.
type createContactUseCase struct {
	service contact.ContactService
}

//NewCreateContactUseCase - This method provides the instance of CreateContactUseCase interface.
func NewCreateContactUseCase(s contact.ContactService) CreateContactUseCase {
	return &createContactUseCase{s}
}

// Execute - This method provides to perform the action of creating a new contact.
func (o *createContactUseCase) Execute(request *CreateContactDTORequest) (*CreateContactDTOResponse, error) {

	var c contact.ContactEntity
	request.Decode(&c)

	err := o.service.Create(&c)

	if err != nil {
		return nil, err
	}

	reponse := &CreateContactDTOResponse{Retcode: 0}

	return reponse, nil
}
