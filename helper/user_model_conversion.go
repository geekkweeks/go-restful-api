package helper

import (
	"geekkweeks/go-restful-api/model/domain"
	"geekkweeks/go-restful-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToLoginResponse(username string, token string) web.LoginResponse {
	return web.LoginResponse{
		Username: username,
		Token:    token,
	}
}
