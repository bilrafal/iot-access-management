package repo

type User struct {
	Id   string
	Name string
}

func NewUser(id string, name string) *User {
	return &User{Id: id, Name: name}
}
