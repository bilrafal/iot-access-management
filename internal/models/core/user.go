package core

import "github.com/google/uuid"

type UserId string

type User struct {
	Id   UserId
	Name string
}

func NewUser(name string) *User {
	id := uuid.New().String()
	return &User{Id: UserId(id), Name: name}
}
