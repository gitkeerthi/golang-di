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
	userDao := user.NewInMemoryUserDao()
	userService := user.NewUserService(userDao)

	user, err := userService.GetUser(100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}

func runWithRealDb() {
	config := map[string]interface{}{
		"host":     "localhost",
		"port":     5000,
		"user":     "admin",
		"password": "secret",
	}

	dataSource := db.NewDataSource(config)
	userDao := user.NewPostgresUserDao(dataSource)
	userService := user.NewUserService(userDao)

	user, err := userService.GetUser(100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}
