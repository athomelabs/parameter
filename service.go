package parameter

type Servicer interface {
	Create(model *Parameter) error
	Update(model *Parameter) error
	UpdateByName(string, string) error
	Delete(id uint) error
	GetByID(id uint) (*Parameter, error)
	GetAll() (Parameters, error)
	GetByName(name string) (*Parameter, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(model *Parameter) error {
	return s.repo.create(model)
}

func (s *Service) Update(model *Parameter) error {
	return s.repo.update(model)
}

func (s *Service) UpdateByName(name string, value string) error {
	return s.repo.updateByName(name, value)
}

func (s *Service) Delete(id uint) error {
	return s.repo.delete(id)
}

func (s *Service) GetByID(id uint) (*Parameter, error) {
	return s.repo.getByID(id)
}

func (s *Service) GetAll() (Parameters, error) {
	return s.repo.getAll()
}

func (s *Service) GetByName(name string) (*Parameter, error) {
	return s.repo.getByName(name)
}
