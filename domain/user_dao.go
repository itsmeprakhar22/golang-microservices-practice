package domain

func GetUser(userId uint64) (User, error) {
	return User{
		123,
		"Prakhar",
		"Sh",
		"ab@cd.com",
	}, nil
}
