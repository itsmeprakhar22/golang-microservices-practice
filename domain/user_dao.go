package domain

import "GolangWorkspace/go-microservices/mvc/utils"

func GetUser(userId int64) (User, *utils.ApplicationError) {
	return User{
		123,
		"Prakhar",
		"Sh",
		"ab@cd.com",
	}, nil
}
