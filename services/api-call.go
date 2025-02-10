package services

import (
	"errors"
	utils "main/utils"
	"time"
)

type User struct {
	name string
	age  int
}

func GetUserData(user chan<- User, err chan<- error) {
	time.Sleep(time.Millisecond * 750)

	isSuccess := utils.RandBool()

	userData := User{
		name: "John",
		age:  25,
	}

	if isSuccess {
		user <- userData
		err <- nil
	} else {
		err <- errors.New("Failed to fetch user data")
	}
}
