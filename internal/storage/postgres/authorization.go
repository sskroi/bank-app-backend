package postgres

import "bank-app-backend/internal/storage"

// Authorization
func (self *PgStorage) CreateUser(email, passwordHash, passport, name, surname string, patronymic *string) (storage.User, error) {
	newUser := storage.User{
		Email:        email,
		PasswordHash: passwordHash,
		Passport:     passport,
		Name:         name,
		Surname:      surname,
		Patronymic:   patronymic,
	}

	err := self.db.Omit("id", "public_id").Create(&newUser).Error
	if err != nil {
		return newUser, err
	}

	err = self.db.First(&newUser, newUser.ID).Error
	return newUser, err
}

func (self *PgStorage) GetUser(email, passwordHash string) (storage.User, error) {
	user := storage.User{}

	res := self.db.Where("email = ? AND password_hash = ?", email, passwordHash).Find(&user)

	return user, res.Error
}
