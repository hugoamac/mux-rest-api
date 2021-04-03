package contact

//ContactService - This type provides the contact service interface
type ContactService interface {
	Create(c *ContactEntity) error
	List() ([]*ContactEntity, error)
	FindById(id string) (*ContactEntity, error)
	Update(c *ContactEntity) (*ContactEntity, error)
	Remove(id string) error
}

//contactService -This struct provides the contact service object
type contactService struct {
	repository ContactRepository
}

//NewContactService - This method provides the instance of ContactService
func NewContactService(repository ContactRepository) ContactService {
	return &contactService{repository}
}

//Create - This method provides the created a new contact
func (s *contactService) Create(c *ContactEntity) error {
	return s.repository.Create(c)
}

//ListAll - This method provides the list of all contacts
func (s *contactService) List() ([]*ContactEntity, error) {
	return s.repository.FetchAll()
}

//FindById - This method provides get the contact by id
func (s *contactService) FindById(id string) (*ContactEntity, error) {
	return s.repository.Find(id)
}

//Update - This method provides update the contact
func (s *contactService) Update(c *ContactEntity) (*ContactEntity, error) {
	return s.repository.Update(c)
}

//Remove - This method provides removed the contract
func (s *contactService) Remove(id string) error {
	return s.repository.Delete(id)
}
