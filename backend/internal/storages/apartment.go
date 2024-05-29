package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/internal/services"

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

func (s *apartmentStorage) UpdateApartment(ApartmentID uint64, filter services.ApartmentFilter) error {
	stmt := s.db.Model(&entity.Apartment{})
	var apartment entity.Apartment

	if ApartmentID != 0 {
		stmt = stmt.Where("id = ?", ApartmentID)
	}

	if filter.BuildingID != nil {
		stmt = stmt.Where("building_id = ?", *filter.BuildingID)
	}
	if filter.UserID != nil {
		stmt = stmt.Where("user_id = ?", *filter.UserID)
	}
	if filter.ApartmentNumber != nil {
		apartment.Number = *filter.ApartmentNumber
	}
	if filter.ApartmentArea != nil {
		apartment.Area = *filter.ApartmentArea
	}
	return stmt.Updates(apartment).Error
}
