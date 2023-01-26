package service

import (
	"context"
	"geekkweeks/go-restful-api/model/web"
)

type UserService interface {
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	Add(ctx context.Context, request web.UserAddRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.UserResponse
	GetAll(ctx context.Context) []web.UserResponse
}
