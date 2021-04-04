package contact

//ContactService - This interface provides the object with the responsibility to perform contact domain services.
type ContactService interface {
	Create(c *ContactEntity) error
	List() ([]*ContactEntity, error)
	FindById(id string) (*ContactEntity, error)
	Update(c *ContactEntity) (*ContactEntity, error)
	Remove(id string) error
}

//contactService - This struct provides the implemtation for ContactService interface.
type contactService struct {
	repository ContactRepository
}

//NewContactService - This method provides the instance of ContactService interface.
func NewContactService(repository ContactRepository) ContactService {
	return &contactService{repository}
}

//Create - This method provides the creation of a new contact.
func (s *contactService) Create(c *ContactEntity) error {
	return s.repository.Create(c)
}

//ListAll - This method provides to retrieve the contact list.
func (s *contactService) List() ([]*ContactEntity, error) {
	return s.repository.FetchAll()
}

//FindById - This method provides to retrieve a contact by the identifier.
func (s *contactService) FindById(id string) (*ContactEntity, error) {
	return s.repository.Find(id)
}

//Update - This method provides for updating contact data.
func (s *contactService) Update(c *ContactEntity) (*ContactEntity, error) {
	return s.repository.Update(c)
}

//Remove - This method provides to remove a contact by the identifier.
func (s *contactService) Remove(id string) error {
	return s.repository.Delete(id)
}
