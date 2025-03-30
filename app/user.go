package app

type User struct {
	UserId int
	Name   string
	Email  string
}

func NewUser(userId int, name, email string) *User {
	return &User{
		UserId: userId,
		Name:   name,
		Email:  email,
	}
}
