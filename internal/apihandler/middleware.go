package apihandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// verifyAuth verifies token, sets userPubIdKey for gin.Context or aborts chain if token is invalid
func (h *Handler) verifyAuth(c *gin.Context) {
	accessToken, err := parseAuthBearerHeader(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, err.Error(), err)
		c.Abort()
		return
	}

	userPublicId, err := h.service.Users.VerifyAccessToken(c.Request.Context(), accessToken)
	if err != nil {
		newErrResponse(c, http.StatusUnauthorized, "invalid access token", err)
		c.Abort()
		return
	}

	c.Set(userPubIdKey, userPublicId)
	c.Next()
}

func parseAuthBearerHeader(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")

	if len(header) == 0 {
		return "", fmt.Errorf("empty Authorization header")
	}

	splitAuthHeader := strings.Split(header, " ")
	if len(splitAuthHeader) != 2 || splitAuthHeader[0] != "Bearer" {
		return "", fmt.Errorf("invalid Authorization header")
	}

	if len(splitAuthHeader[1]) == 0 {
		return "", fmt.Errorf("empty Bearer token")
	}

	return splitAuthHeader[1], nil
}
