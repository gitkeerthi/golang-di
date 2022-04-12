package user


// Consumer of a dependency - a service
type UserService struct {
	userDao UserDao
}

func (s *UserService) CreateUser(user User) error {
	return s.userDao.Insert(user)
}

func (s *UserService) GetUser(id int) (User, error) {
	return s.userDao.Get(id)
}

// A factory function to create the service
// The service accepts its dependency as an interface
func NewUserService(dao UserDao) *UserService {
	return &UserService{
		userDao: dao,
	}
}
