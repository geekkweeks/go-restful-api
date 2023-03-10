package main

import (
	"geekkweeks/go-restful-api/app"
	"geekkweeks/go-restful-api/controller"
	"geekkweeks/go-restful-api/exception"
	"geekkweeks/go-restful-api/helper"
	"geekkweeks/go-restful-api/middleware"
	"geekkweeks/go-restful-api/repository"
	"geekkweeks/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()

	router.POST("/api/users/login", userController.Login)

	router.GET("/api/users", userController.GetAll)
	router.GET("/api/users/:userId", userController.FindById)
	router.POST("/api/users", userController.Add)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

	// error handler
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router), // every request will check by middleware first before continuing to next handler
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
