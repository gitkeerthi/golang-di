package user

import "github.com/kpkeerthi/golang-di/db"

type UserDao interface {
	Insert(user User) error
	Get(id int) (User, error)
}

type UserService struct {
	userDao UserDao
}

type User struct {
	Id   int
	Name string
}

// Accept dependency as an interface and return concrete struct
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

// Factory function to create the real dao
func NewUserDao(ds *db.DataSource) *PersistingUserDao {
	return &PersistingUserDao{
		dataSource: ds,
	}
}

// This is our real Dao
type PersistingUserDao struct {
	dataSource *db.DataSource
}

func (dao *PersistingUserDao) Insert(user User) error {
	db := dao.dataSource.Db
	db.Exec("") // implement real interaction with the db
	return nil
}

func (dao *PersistingUserDao) Get(id int) (User, error) {
	db := dao.dataSource.Db
	db.Exec("") // implement real interaction with the db
	return User{}, nil
}

// This is our mock dao. Can be used for unit testing the service
type InMemoryUserDao struct {
	dataSource *db.DataSource
}

func (dao *InMemoryUserDao) Insert(user User) error {
	// emulate db operations in memory
	return nil
}

func (dao *InMemoryUserDao) Get(id int) (User, error) {
	// emulate db operations in memory
	return User{Id: 10, Name: "John Doe"}, nil
}

// Factory function to create the mock dao
func NewMockUserDao() *InMemoryUserDao {
	return &InMemoryUserDao{
		dataSource: nil,
	}
}
