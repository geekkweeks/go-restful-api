package service

import (
	"context"
	"database/sql"
	"geekkweeks/go-restful-api/config"
	"geekkweeks/go-restful-api/exception"
	"geekkweeks/go-restful-api/helper"
	"geekkweeks/go-restful-api/model"
	"geekkweeks/go-restful-api/model/domain"
	"geekkweeks/go-restful-api/model/web"
	"geekkweeks/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	// check if user by username and password is found
	user, err := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	// if data is not found, will be given the error not found
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// check the password is match or not
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicIfError(err)

	// Generate token
	expTime := time.Now().Add(time.Minute * 2)
	claims := &model.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// implementation of algorithm in token
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	helper.PanicIfError(err)

	return helper.ToLoginResponse(user.Username, token)
}

func (service *UserServiceImpl) Add(ctx context.Context, request web.UserAddRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	// hash user password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	request.Password = string(hashPassword)
	userEntity := domain.User{
		Username:  request.Username,
		Password:  request.Password,
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
	// if data is not found, will be given the error not found
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

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
	// if data is not found, will be given the error not found
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, userEntity)
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	userEntity, err := service.UserRepository.FindById(ctx, tx, id)
	// if data is not found, will be given the error not found
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(userEntity)
}

func (service *UserServiceImpl) GetAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.TxCommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)

	return helper.ToUserResponses(users)
}
