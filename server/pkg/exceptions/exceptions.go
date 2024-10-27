package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Exception struct {
	Error      error
	StatusCode int
}

func NewException(statusCode int, err error) *Exception {
	return &Exception{
		Error:      err,
		StatusCode: statusCode,
	}
}

func HandleBadRequestException(context *gin.Context, err error) {
	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"statusText": "failure",
		"statusCode": 400,
		"errorType":  "BadRequestException",
		"error":      err.Error(),
	})
}

func HandleConflictException(context *gin.Context, errText string) {
	context.AbortWithStatusJSON(http.StatusConflict, gin.H{
		"statusText": "failed",
		"statusCode": http.StatusConflict,
		"errorType":  "ConflictException",
		"error":      errText,
	})
}

func HandleInternalServerException(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"statusText": "failed",
		"statusCode": http.StatusInternalServerError,
		"errorType":  "InternalServerErrorException",
		"error":      "something went wrong",
	})
}

func HandleNotFoundException(context *gin.Context, err error) {
	context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"statusText": "failure",
		"statusCode": http.StatusBadRequest,
		"errorType":  "NotFoundException",
		"error":      err,
	})
}

func HandleValidationException(context *gin.Context, err error) {
	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"statusText": "failure",
		"statusCode": http.StatusBadRequest,
		"errorType":  "ValidationException",
		"error":      err.Error(),
	})
}
