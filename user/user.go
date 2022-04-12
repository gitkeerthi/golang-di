package user

import "github.com/kpkeerthi/golang-di/db"

// The dependency, modelled as an interface
type UserDao interface {
	Insert(user User) error
	Get(id int) (User, error)
}

// Consumer of a dependency - a service
type UserService struct {
	userDao UserDao
}

// Domain model
type User struct {
	Id   int
	Name string
}

// A factory function to create the service
// The service accepts its dependency as an interface
func NewUserService(dao UserDao) *UserService {
	return &UserService{
		userDao: dao,
	}
}

func (s *UserService) CreateUser(user User) error {
	return s.userDao.Insert(user)
}

func (s *UserService) GetUser(id int) (User, error) {
	return s.userDao.Get(id)
}

// This is our real Dao
type PostgresUserDao struct {
	dataSource *db.DataSource
}

func (p *PostgresUserDao) Insert(user User) error {
	db := p.dataSource.Db
	db.Exec("") // implement real interaction with the db
	return nil
}

func (p *PostgresUserDao) Get(id int) (User, error) {
	db := p.dataSource.Db
	db.Exec("") // implement real interaction with the db
	return User{}, nil
}

// A factory function to create the real dao
func NewPostgresUserDao(ds *db.DataSource) *PostgresUserDao {
	return &PostgresUserDao{
		dataSource: ds,
	}
}

// This is our mock dao. Can be used for unit testing the service.
type InMemoryUserDao struct {
	dataSource *db.DataSource
}

func (m *InMemoryUserDao) Insert(user User) error {
	// emulate db operations in memory
	return nil
}

func (m *InMemoryUserDao) Get(id int) (User, error) {
	// emulate db operations in memory
	return User{Id: 10, Name: "John Doe"}, nil
}

// A factory function to create the mock in-memory dao
func NewInMemoryUserDao() *InMemoryUserDao {
	return &InMemoryUserDao{
		dataSource: nil,
	}
}
