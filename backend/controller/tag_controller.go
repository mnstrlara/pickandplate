package controller

import (
	"net/http"
	"pickandplate/backend/data/request"
	"pickandplate/backend/data/response"
	"pickandplate/backend/helper"
	"pickandplate/backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService service.TagsService
}
// @title Pick and Plate API
// @version 1.0
// @description This is the API for Pick and Plate
// @host localhost:8000
// @BasePath /api/v1

// @tag.name Tags
// @tag.description Tag management

// FindAll godoc
// @Summary Get all tags
// @Description Get a list of all tags
// @Tags Tags
// @Accept json
// @Produce json
// @Success 200 {array} model.Tags
// @Router /tag [get]

func NewTagController(service service.TagsService) *TagController {
	return &TagController{tagService: service}
}

func (controller *TagController) Create(ctx *gin.Context) {
	createTagRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.tagService.Create(createTagRequest)

	webResponse := response.Response{
		Code: 200,
		Status: "Ok",
		Data: nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateTagRequest.Id = id

	controller.tagService.Update(updateTagRequest)

	webResponse := response.Response{
		Code: 200,
		Status: "Ok",
		Data: nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagService.Delete(id)

	webResponse := response.Response{
		Code: 200,
		Status: "Ok",
		Data: nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagService.FindById(id)

	webResponse := response.Response{
		Code: 200,
		Status: "Ok",
		Data: tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagService.FindAll()

	webResponse := response.Response{
		Code: 200,
		Status: "Ok",
		Data: tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}