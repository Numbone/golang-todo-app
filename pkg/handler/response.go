package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.WithFields(logrus.Fields{
		"status": statusCode,
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
		"client": c.ClientIP(),
	}).Error(message)
	c.AbortWithStatusJSON(statusCode, &ErrorResponse{message})
}
