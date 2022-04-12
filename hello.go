package main

import (
	"fmt"

	"github.com/kpkeerthi/golang-ci/db"
	"github.com/kpkeerthi/golang-ci/user"
)

func main() {
	runWithMockDb()
	runWithRealDb()
}

func runWithMockDb() {
	userDao := user.NewMockUserDao()
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
	userDao := user.NewUserDao(dataSource)
	userService := user.NewUserService(userDao)

	user, err := userService.GetUser(100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}
