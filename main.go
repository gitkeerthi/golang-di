package main

import (
	"fmt"

	"github.com/gitkeerthi/golang-di/db"
	"github.com/gitkeerthi/golang-di/user"
)

func main() {
	runWithMockDb()
	if false {
		runWithRealDb()
	}
}

func runWithMockDb() {
	userRepository := user.NewInMemoryUserRepository()
	userService := user.NewUserService(userRepository)

	user, err := userService.GetUser(100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}

func runWithRealDb() {
	dataSource := db.NewDataSource("user=postgres dbname=test password=secret host=localhost sslmode=disable")
	defer dataSource.Db.Close()

	userRepository := user.NewPostgresUserRepository(dataSource)
	userService := user.NewUserService(userRepository)

	user, err := userService.GetUser(100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}
