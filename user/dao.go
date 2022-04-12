package user

import "github.com/gitkeerthi/golang-di/db"

// The dependency, modelled as an interface
type UserDao interface {
	Insert(user User) error
	Get(id int) (User, error)
}

// This is our real (persistence-aware) Dao
type PostgresUserDao struct {
	dataSource *db.DataSource
}

func (p *PostgresUserDao) Insert(user User) error {
	db := p.dataSource.Db
	db.Exec("INSERT INTO ...") // implement real interaction with the db
	return nil
}

func (p *PostgresUserDao) Get(id int) (User, error) {
	db := p.dataSource.Db
	db.Query("SELECT * FROM ...") // implement real interaction with the db
	return User{}, nil
}

// A factory function to create the real dao
func NewPostgresUserDao(ds *db.DataSource) *PostgresUserDao {
	return &PostgresUserDao{
		dataSource: ds,
	}
}

// This is our mock (in-memory) dao. Can be used for unit testing the service.
type InMemoryUserDao struct {
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
	return &InMemoryUserDao{}
}
