package main

import (
	"net/http"
	"pickandplate/backend/config"
	"pickandplate/backend/controller"
	"pickandplate/backend/helper"
	"pickandplate/backend/model"
	"pickandplate/backend/repository"
	"pickandplate/backend/router"
	"pickandplate/backend/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

import _ "pickandplate/backend/docs"

func main() {
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagRepository := repository.NewTagsRepositoryImpl(db)
	tagService := service.NewTagServiceImpl(tagRepository, validate)
	tagController := controller.NewTagController(tagService)

	r := router.NewRouter(tagController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
