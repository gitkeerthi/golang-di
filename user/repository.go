package user

import "github.com/gitkeerthi/golang-di/db"

// UserService's dependency, modelled as an interface
type UserRepository interface {
	Insert(user User) error
	Get(id int) (User, error)
}

// This is our real (persistence-aware) repository
type PostgresUserRepository struct {
	dataSource *db.DataSource
}

func (p *PostgresUserRepository) Insert(user User) error {
	db := p.dataSource.Db
	db.Exec("INSERT INTO ...") // implement real interaction with the db
	return nil
}

func (p *PostgresUserRepository) Get(id int) (User, error) {
	db := p.dataSource.Db
	db.Query("SELECT * FROM ...") // implement real interaction with the db
	return User{}, nil
}

// A factory function to create the real repository
func NewPostgresUserRepository(ds *db.DataSource) *PostgresUserRepository {
	return &PostgresUserRepository{
		dataSource: ds,
	}
}

// This is our mock (in-memory) repository. Can be used for unit testing the service.
type InMemoryUserRepository struct {
}

func (m *InMemoryUserRepository) Insert(user User) error {
	// emulate db operations in memory
	return nil
}

func (m *InMemoryUserRepository) Get(id int) (User, error) {
	// emulate db operations in memory
	return User{Id: 10, Name: "John Doe"}, nil
}

// A factory function to create the mock in-memory repository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}
