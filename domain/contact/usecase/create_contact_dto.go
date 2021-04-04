package usecase

import (
	"mux-rest-api/domain/contact"
)

//CreateContactDTORequest - This struct provides the DTO for Request of CreateContactUseCase.
type CreateContactDTORequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

//Decode - This method provides the parse CreateContactDTORequest to ContactEntity.
func (d *CreateContactDTORequest) Decode(c *contact.ContactEntity) {

	c.Name = d.Name
	c.LastName = d.LastName
	c.Email = d.Email
	c.Phone = d.Phone
}

//CreateContactDTOResponse - This method provides the DTO for Response of CreateContactUseCase.
type CreateContactDTOResponse struct {
	Retcode int `json:"retcode"`
}
