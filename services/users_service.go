package services

import (
	"GolangWorkspace/go-microservices/mvc/domain"
	"GolangWorkspace/go-microservices/mvc/utils"
)

func GetUser(userId int64) (domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
