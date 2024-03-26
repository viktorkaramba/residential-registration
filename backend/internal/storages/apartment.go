package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"

	"gorm.io/gorm"
)

type apartmentStorage struct {
	db *gorm.DB
}

func NewApartmentStorage(db *gorm.DB) *apartmentStorage {
	return &apartmentStorage{
		db: db,
	}
}

func (s *apartmentStorage) GetByUserID(UserID uint64) (*entity.Apartment, error) {
	apartment := &entity.Apartment{}
	err := s.db.Model(&entity.Apartment{}).Where(entity.Apartment{UserID: UserID}).First(apartment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return apartment, err
}
