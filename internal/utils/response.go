package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func JSONError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}
