package service

import "bank-app-backend/internal/storage"

type AuthService struct {
	store storage.Storage
}

func NewAuthService(store storage.Storage) *AuthService {
	return &AuthService{store}
}


