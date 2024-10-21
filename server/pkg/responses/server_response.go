package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleOkDataResponse(context *gin.Context, message string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": http.StatusOK,
		"message":    message,
		"data":       data,
	})
}

func HandleOkMessageResponse(context *gin.Context, message string) {
	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": http.StatusOK,
		"message":    message,
	})
}

func HandleCreatedResponse(context *gin.Context, message string, data interface{}) {
	context.JSON(http.StatusCreated, gin.H{
		"statusText": "success",
		"statusCode": http.StatusCreated,
		"message":    message,
		"data":       data,
	})
}
