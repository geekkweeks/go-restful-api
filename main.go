package main

import (
	"geekkweeks/go-restful-api/app"
	"geekkweeks/go-restful-api/controller"
	"geekkweeks/go-restful-api/repository"
	"geekkweeks/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()

	router.GET("/api/users", userController.GetAll)
	router.GET("/api/users/:userId", userController.FindById)
	router.POST("/api/users/", userController.Add)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

}
