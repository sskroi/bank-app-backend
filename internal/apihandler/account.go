package apihandler

import (
	"bank-app-backend/internal/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const DefaultAccountsLimit int = 100

type userAccountsInput struct {
	Offset	   int     `binding:"gte=0"`
	Limit	   int	   `binding:"gte=0,lte=100"`
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

	if input.Limit == 0 {
		input.Limit = DefaultAccountsLimit
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
