package artist

import (
	"net/http"
	"strconv"

	"final-project-api/internal/helper"

	"github.com/gin-gonic/gin"
)

// It will call the function Get in artist usecase
func (handler *artistHandler) Get(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	artist, err := handler.artistUsecase.Get(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artist)
	context.JSON(http.StatusOK, res)
}

// It will call the function GetAllartist in artist usecase
func (handler *artistHandler) GetAll(context *gin.Context) {
	limit, err := strconv.ParseInt(context.Request.FormValue("limit"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param limit was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	offset, err := strconv.ParseInt(context.Request.FormValue("offset"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param offset was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	// Get all artists from usecase
	artists, err := handler.artistUsecase.GetAll(context, int32(limit), int32(offset))
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artists)
	context.JSON(http.StatusOK, res)
}
