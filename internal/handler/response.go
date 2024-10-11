package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type errResponse struct {
	Message string `json:"message"`
}

func newErrResponse(c *gin.Context, statusCode int, err error, msg string) {
	slog.Warn("new error response", "err", err, "resp msg", msg)

	c.AbortWithStatusJSON(statusCode, errResponse{msg})
}
