package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"

	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *userStorage {
	return &userStorage{
		db: db,
	}
}

func (s *userStorage) CreateUser(User *entity.User) error {
	return s.db.Create(User).Error
}

func (s *userStorage) GetUser(UserID uint64) (*entity.User, error) {
	User := &entity.User{}
	err := s.db.Model(&entity.User{}).Where(entity.User{ID: UserID}).First(User).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return User, err
}

func (s *userStorage) GetUserByPhoneNumber(phoneNumber entity.PhoneNumber) (*entity.User, error) {
	User := &entity.User{}
	err := s.db.Model(&entity.User{}).Where(entity.User{PhoneNumber: phoneNumber}).First(User).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return User, err
}
