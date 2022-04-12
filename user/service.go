package user

// Consumer of a dependency - a service
type UserService struct {
	userRepository UserRepository
}

func (s *UserService) CreateUser(user User) error {
	return s.userRepository.Insert(user)
}

func (s *UserService) GetUser(id int) (User, error) {
	return s.userRepository.Get(id)
}

// A factory function to create the service
// The service accepts its dependency as an interface
func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}
