package repo

type User struct {
	UserId string
	Name   string
}

func NewUser(userId string, name string) *User {
	return &User{UserId: userId, Name: name}
}
