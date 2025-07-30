package user

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	FindByID(id int) (*User, error)
}

type UserService struct {
	repo UserRepository
}

// Instead of concrete type, we are passing interface so that we can mock and test it easily.
func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
