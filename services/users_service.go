package services

import "GolangWorkspace/go-microservices/mvc/domain"

func GetUser(userId uint64) (domain.User, error) {
	return domain.GetUser(userId)
}
