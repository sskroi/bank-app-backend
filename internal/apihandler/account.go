package apihandler

import (
	"bank-app-backend/internal/domain"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const DefaultAccountsLimit int = 100

var Currencies = [...]string{"rub", "usd", "eur", "cny"}

type createAccountInput struct {
	Currency string `binding:"required,len=3"`
}
type createAccountResponse struct {
	Number string `json:"number"`
}

// @Summary		Create bank account
// @Description	Creates bank account
// @Security  UserBearerAuth
// @Accept			json
// @Produce		json
// @Param			input	body		createAccountInput	true	"Account info"
// @Success		201		{object}	createAccountResponse
// @Failure		400		{object}	response
// @Failure   401   {object}  response
// @Failure   403   {object}  response "User deleted or banned"
// @Failure		404		{object}	response
// @Failure		500		{object}	response
// @Router			/account [post]
func (h *Handler) createAccount(c *gin.Context) {
	var input createAccountInput

	if err := c.BindJSON(&input); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body", err)
		return
	}

	userPubId, err := getUserPubId(c)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		return
	}

	currency := strings.ToLower(input.Currency)
	found := false
	for _, c := range Currencies {
		if currency == c {
			found = true
			break
		}
	}
	if !found {
		newResponse(c, http.StatusBadRequest, domain.ErrUnknownCurrency.Error())
		return
	}

	accountNumber, err := h.service.Accounts.Create(
		c.Request.Context(), userPubId, currency)
	if err != nil {
		if errors.Is(err, domain.ErrUserDeleted) {
			newResponse(c, http.StatusForbidden, domain.ErrUserDeleted.Error())
			return
		}

		newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.JSON(http.StatusCreated, createAccountResponse{accountNumber.String()})
}

type userAccountsInput struct {
	Offset int `binding:"gte=0"`
	Limit  int `binding:"gte=0,lte=100"`
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
		if errors.Is(err, domain.ErrUserDeleted) {
			newResponse(c, http.StatusConflict, domain.ErrUserDeleted.Error())
			return
		}

		newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	c.JSON(http.StatusOK, accounts)
}
