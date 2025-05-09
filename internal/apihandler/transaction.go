package apihandler

import (
	"bank-app-backend/internal/domain"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const DefaultTransactionsLimit int = 100

type createTransactionInput struct {
	SenderAccNumber		uuid.UUID		`json:"senderAccountNumber" binding:"required"`
	ReceiverAccNumber 	uuid.UUID 		`json:"receiverAccountNumber" binding:"required"`
	Amount 				decimal.Decimal `json:"amount" binding:"required"`
}
type createTransactionResponse struct {
	PublicId		  uuid.UUID 	  `json:"publicId"`
	SenderAccNumber   uuid.UUID		  `json:"senderAccountNumber"`
	ReceiverAccNumber uuid.UUID 	  `json:"receiverAccountNumber"`
	Sent			  decimal.Decimal `json:"sent"`
	Received		  decimal.Decimal `json:"received"`
	IsConversion	  bool			  `json:"isConversion"`
	ConversionRate	  decimal.Decimal `json:"conversionRate"`
	Timestamp		  time.Time		  `json:"timestamp"`
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
		} else if errors.Is(err, domain.ErrNegativeSenderBalance) ||
				errors.Is(err, domain.ErrSenderAccountClose) ||
				errors.Is(err, domain.ErrReceiverAccountClose) ||
				strings.Contains(err.Error(), "no conversion from") {
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
		Timestamp: newTransaction.Timestamp,
	})
}

type transactionsInput struct {
	AccNumber	string	`form:"accountNumber"`
	Offset		int		`form:"offset" binding:"gte=0"`
	Limit		int		`form:"limit" binding:"gte=0,lte=100"`
}

// @Summary		Get all user's transactions
// @Security  UserBearerAuth
// @Produce		json
// @Param			accountNumber query		string	  false	"Account number"
// @Param			offset query		int	  false	"Offset" minimum(0)
// @Param			limit  query		int	  false	"Limit"  minimum(0) maximum(100)
// @Success		200		{array}	  domain.TransactionExtended
// @Failure		400		{object}	response
// @Failure   401   {object}  response
// @Failure   403   {object}  response "User deleted or banned"
// @Failure		404		{object}	response
// @Failure		500		{object}	response
// @Router			/transactions [get]
func (h *Handler) userTransactions(c *gin.Context) {
	var input transactionsInput

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

	var accNumber *uuid.UUID
	if _uuid, err := uuid.Parse(input.AccNumber); err == nil {
		accNumber = &_uuid
	}
	transactions, err := h.service.Transactions.UserTransactions(
		c.Request.Context(), userPubId, accNumber, input.Offset, input.Limit)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
	} else {
		c.JSON(http.StatusOK, transactions)	
	}
}
