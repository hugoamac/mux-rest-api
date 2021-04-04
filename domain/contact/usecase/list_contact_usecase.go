package usecase

import (
	"mux-rest-api/domain/contact"
)

//ListContactUseCase - This interface provides an object with the responsibility to retrieve a list of contacts.
type ListContactUseCase interface {
	Execute() (*ListContactDTOResponse, error)
}

//listContactUseCase - This struct provides the implementation of ListContactUseCase interface.
type listContactUseCase struct {
	service contact.ContactService
}

//NewListContactUseCase - This method provides the instance of ListContactUseCase interface.
func NewListContactUseCase(s contact.ContactService) ListContactUseCase {
	return &listContactUseCase{s}
}

// Execute - This method provides to perform the action of retrieving a list of contacts.
func (o *listContactUseCase) Execute() (*ListContactDTOResponse, error) {

	contacts, err := o.service.List()

	if err != nil {
		return nil, err
	}

	list := &ListContactDTOResponse{Retcode: contact.RETCODE_NOT_FOUND}

	if len(contacts) > 0 {
		list = &ListContactDTOResponse{Retcode: contact.RETCODE_OK, Contacts: contacts}
	}

	return list, nil
}
