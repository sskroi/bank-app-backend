package apihandler

import (
	"bank-app-backend/internal/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userAccountsInput struct {
	Offset	   int     `json:"offset" binding:"omitempty,gte=0"`
	Limit	   int	   `json:"limit" binding:"omitempty,gte=0"`
}

func (h *Handler) userAccounts(c *gin.Context) {
	var input userAccountsInput

	if err := c.BindJSON(&input); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body", err)
		return
	}

	userPubId, err := getUserPubId(c)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		return
	}

	accounts, err := h.service.Accounts.UserAccounts(
		c.Request.Context(), userPubId, input.Offset, input.Limit)

	if err != nil {
		if errors.Is(err, domain.ErrUnknownUserPubId) {
			newResponse(c, http.StatusConflict, domain.ErrUnknownUserPubId.Error())
			return
		}

		newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.JSON(http.StatusOK, accounts)
}
