package service

import (
	"context"
	"database/sql"
	"geekkweeks/go-restful-api/helper"
	"geekkweeks/go-restful-api/model/domain"
	"geekkweeks/go-restful-api/model/web"
	"geekkweeks/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB             //need pointer(*) because DB is struct
	Validate       *validator.Validate //need pointer(*) because Validate is struct
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Add(ctx context.Context, request web.UserAddRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	userEntity := domain.User{
		Username:  request.Username,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Phone:     request.Phone,
	}

	userEntity = service.UserRepository.Add(ctx, tx, userEntity)

	return helper.ToUserResponse(userEntity)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	// check if user by id is exist
	userEntity, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// if exist
	userEntity.Username = request.Username
	userEntity.FirstName = request.FirstName
	userEntity.LastName = request.LastName
	userEntity.Phone = request.Phone

	userEntity = service.UserRepository.Update(ctx, tx, userEntity)

	return helper.ToUserResponse(userEntity)
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	// check if user by id is exist
	userEntity, err := service.UserRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	service.UserRepository.Delete(ctx, tx, userEntity)
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	userEntity, err := service.UserRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToUserResponse(userEntity)
}

func (service *UserServiceImpl) GetAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)

	return helper.ToUserResponses(users)
}
