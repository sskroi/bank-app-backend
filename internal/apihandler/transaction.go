package apihandler

import (
	"bank-app-backend/internal/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const DefaultTransactionsLimit int = 100

type createTransactionInput struct {
	SenderAccNumber		uuid.UUID		`json:"sender_account_number" binding:"required"`
	ReceiverAccNumber 	uuid.UUID 		`json:"receiver_account_number" binding:"required"`
	Amount 				decimal.Decimal `json:"amount" binding:"required"`
}
type createTransactionResponse struct {
	PublicId		  uuid.UUID 	  `json:"public_id"`
	SenderAccNumber   uuid.UUID		  `json:"sender_account_number"`
	ReceiverAccNumber uuid.UUID 	  `json:"receiver_account_number"`
	Sent			  decimal.Decimal `json:"sent"`
	Received		  decimal.Decimal `json:"received"`
	IsConversion	  bool			  `json:"is_conversion"`
	ConversionRate	  decimal.Decimal `json:"conversion_rate"`
}

// @Summary		Create transaction
// @Security  UserBearerAuth
// @Accept		json
// @Produce		json
// @Param			input	body		  createTransactionInput	true	"Transaction info"
// @Success		201		{object}	createTransactionResponse
// @Failure   400   {object}  response
// @Failure   401   {object}  response
// @Failure   403   {object}  response "User deleted or banned"
// @Failure		404		{object}	response "Receiver or sender account not found"
// @Failure		500		{object}	response
// @Router			/transaction [post]
func (h *Handler) createTransaction(c *gin.Context) {
	var input createTransactionInput

	if err := c.BindJSON(&input); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body", err)
		return
	}

	userPubId, err := getUserPubId(c)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		return
	}

	newTransaction, err := h.service.Transactions.Create(
		c.Request.Context(), userPubId, input.SenderAccNumber,
		input.ReceiverAccNumber, input.Amount)
	if err != nil {
		if errors.Is(err, domain.ErrUserDeleted) {
			newResponse(c, http.StatusForbidden, domain.ErrUserDeleted.Error())
			return
		}
		var statusCode int
		if errors.Is(err, domain.ErrSelfAccount) || errors.Is(err, domain.ErrInvalidAmount) {
			statusCode = http.StatusBadRequest
		} else if errors.Is(err, domain.ErrUnknownSender) || errors.Is(err, domain.ErrUnknownReceiver) {
			statusCode = http.StatusNotFound
		} else if errors.Is(err, domain.ErrNegativeSenderBalance) {
			statusCode = http.StatusForbidden
		} else {
			newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
			return
		}
		newErrResponse(c, statusCode, err.Error(), err)
		return
	}

	c.JSON(http.StatusCreated, createTransactionResponse{
		PublicId: newTransaction.PublicId,
		SenderAccNumber: input.SenderAccNumber,
		ReceiverAccNumber: input.ReceiverAccNumber,
		Sent: newTransaction.Sent,
		Received: newTransaction.Received,
		IsConversion: newTransaction.IsConversion,
		ConversionRate: newTransaction.ConversionRate,
	})
}
