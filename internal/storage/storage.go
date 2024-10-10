package storage

type Storage interface {
	// Authorization
	CreateUser(email, passwordHash, passport, name, surname string, patronymic *string) (User, error)
	GetUser(email, passwordHash string) (User, error)

	// Account
}

