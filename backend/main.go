package main

import (
	"net/http"
	"pickandplate/backend/config"
	"pickandplate/backend/controller"
	"pickandplate/backend/helper"
	"pickandplate/backend/repository"
	"pickandplate/backend/router"
	"pickandplate/backend/service"
	"pickandplate/backend/model"
	"time"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagRepository := repository.NewTagsRepositoryImpl(db)

	tagService := service.NewTagServiceImpl(tagRepository, validate)

	tagController := controller.NewTagController(tagService)

	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr: ":8888",
		Handler: routes,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
