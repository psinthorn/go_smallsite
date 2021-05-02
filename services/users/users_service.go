package services

var (
	UsersService usersServiceInterface = &userServices{}
)

type usersServiceInterface interface {
	Create() error
	Get() error
	Update() error
	Delete() error
}

type userServices struct {
}

func (s *userServices) Create() error {
	return nil
}

func (s *userServices) Get() error {
	return nil
}

func (s *userServices) Update() error {
	return nil
}

func (s *userServices) Delete() error {
	return nil
}
