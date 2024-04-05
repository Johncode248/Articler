package users

type User struct {
	UserId   string
	Username string
	Password string
}

func GetPassword(user User) string {
	return user.Password
}
