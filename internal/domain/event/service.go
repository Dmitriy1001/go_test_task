package event

type Service interface {
	Create(eventData map[string]string) error
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Create(eventData map[string]string) error {
	return (*s.repo).Create(eventData)
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return (*s.repo).FindOne(id)
}
