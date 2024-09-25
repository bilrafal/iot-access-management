package api

type UserResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
}
