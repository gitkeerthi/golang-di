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

func NewUserDao(ds *db.DataSource) *PersistingUserDao {
	return &PersistingUserDao{
		dataSource: ds,
	}
}

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

func NewMockUserDao() *InMemoryUserDao {
	return &InMemoryUserDao{
		dataSource: nil,
	}
}

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
