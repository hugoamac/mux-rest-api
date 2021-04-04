package usecase

import (
	"mux-rest-api/domain/contact"
)

//ListContactDTOResponse - This method provides the DTO for Response of ListContactUseCase.
type ListContactDTOResponse struct {
	Retcode  int                      `json:"retcode"`
	Contacts []*contact.ContactEntity `json:"contacts,omitempty"`
}
