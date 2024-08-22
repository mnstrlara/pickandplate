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