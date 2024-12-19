package service

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/storage"
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionService struct {
	store storage.Storage
}

func NewTransactionService(store storage.Storage) *TransactionService {
	return &TransactionService{
		store: store,
	}
}

func (s TransactionService) Create(
		ctx context.Context,
		userPubId, senderAccNumber, receiverAccNumber uuid.UUID,
		amount decimal.Decimal) (domain.Transaction, error) {
	newTransaction := domain.Transaction{
		Sent: amount,
	}

	senderId, err := s.store.GetUserId(ctx, userPubId)
	if err != nil {
		return newTransaction, err
	}

	senderAcc, err := s.store.GetAccountByNumber(
		ctx, senderAccNumber, senderId, domain.ErrUnknownSender)
	if err != nil {
		return newTransaction, err
	}
	if senderAcc.IsClose {
		return newTransaction, domain.ErrSenderAccountClose
	}
	newTransaction.SenderAccId = senderAcc.ID

	receiverAcc, err := s.store.GetAccountByNumber(
		ctx, receiverAccNumber, 0, domain.ErrUnknownReceiver)
	if err != nil {
		return newTransaction, err
	}
	if receiverAcc.IsClose {
		return newTransaction, domain.ErrReceiverAccountClose
	}
	newTransaction.ReceiverAccId = receiverAcc.ID

	err = s.store.CreateTransaction(ctx, senderAcc, receiverAcc, &newTransaction)
	return newTransaction, err
}

func (s TransactionService) UserTransactions(
		ctx context.Context, userPubId uuid.UUID,
		accountNumber *uuid.UUID, offset, limit int) ([]domain.TransactionExtended, error) {
	var transactions []domain.TransactionExtended

	userId, err := s.store.GetUserId(ctx, userPubId)
	if err != nil {
		return transactions, err
	}

	return s.store.GetUserTransactions(ctx, userId, accountNumber, offset, limit)
}
