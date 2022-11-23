package response

import "samsul96maarif/github.com/go-api-app/model"

type UserResponse struct {
	model.User
}

type LoginResponse struct {
	Token string `json:"token"`
}

func ConvertUserModelToUserResponse(entity model.User) (res UserResponse) {
	res = UserResponse{
		entity,
	}
	return res
}
