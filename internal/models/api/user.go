package api

type UserId string

type UserResponse struct {
	Id   UserId `json:"id"`
	Name string `json:"name"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
}
