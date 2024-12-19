package apihandler

import (
	"bank-app-backend/internal/domain"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const DefaultAccountsLimit int = 100

var Currencies = [...]string{"rub", "usd", "eur"}

type createAccountInput struct {
	Currency string `json:"currency" binding:"required,len=3" enums:"rub"`
}
type createAccountResponse struct {
	Number string `json:"number"`
}

// @Summary		Create bank account
// @Security  UserBearerAuth
// @Accept		json
// @Produce		json
// @Param			input	body		  createAccountInput	true	"Account info"
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

type closeAccountInput struct {
	Number string `form:"number" binding:"required"`
}

// @Summary		Close account
// @Security  UserBearerAuth
// @Produce		json
// @Param			number query		string	  true	"Account number"
// @Success		200		{object}	  response
// @Failure		400		{object}	response
// @Failure   401   {object}  response
// @Failure		403		{object}	response
// @Failure		404		{object}	response
// @Failure		500		{object}	response
// @Router			/account [delete]
func (h* Handler) closeAccount(c *gin.Context) {
	var input closeAccountInput

	if err := c.BindQuery(&input); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input", err)
		return
	}
	number, err := uuid.Parse(input.Number)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid number format")
		return
	}

	userPubId, err := getUserPubId(c)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		return
	}

	if err := h.service.Accounts.Close(
			c.Request.Context(), userPubId, number); err != nil {
		if errors.Is(err, domain.ErrAlreadyClose) || errors.Is(err, domain.ErrClose) {
			newResponse(c, http.StatusForbidden, err.Error())
		} else if errors.Is(err, domain.ErrUnknownAccount) {
			newResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
		}
	} else {
		newResponse(c, http.StatusOK, "success")
	}
}

type userAccountsInput struct {
	Offset int `form:"offset" binding:"gte=0"`
	Limit  int `form:"limit" binding:"gte=0,lte=100"`
}

// @Summary		Get all user's accounts
// @Security  UserBearerAuth
// @Produce		json
// @Param			offset query		int	  false	"Offset" minimum(0)
// @Param			limit  query		int	  false	"Limit"  minimum(0) maximum(100)
// @Success		200		{array}	  domain.Account
// @Failure		400		{object}	response
// @Failure   401   {object}  response
// @Failure   403   {object}  response "User deleted or banned"
// @Failure		404		{object}	response
// @Failure		500		{object}	response
// @Router			/accounts [get]
func (h *Handler) userAccounts(c *gin.Context) {
	var input userAccountsInput

	if err := c.BindQuery(&input); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input", err)
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
