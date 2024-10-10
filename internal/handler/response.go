package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errResponse struct {
	Message string `json:"message"`
}

func newErrResponse(c *gin.Context, statusCode int, err error, msg string) {
	log.Printf("err: %s, msg: %s", err, msg)

	c.AbortWithStatusJSON(statusCode, errResponse{msg})
}
