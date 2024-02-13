package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"

	"gorm.io/gorm"
)

type inhabitantStorage struct {
	db *gorm.DB
}

func NewInhabitantStorage(db *gorm.DB) *inhabitantStorage {
	return &inhabitantStorage{
		db: db,
	}
}

func (s *inhabitantStorage) CreateInhabitant(inhabitant *entity.Inhabitant) error {
	return s.db.Create(inhabitant).Error
}

func (s *inhabitantStorage) GetInhabitant(inhabitantID uint64) (*entity.Inhabitant, error) {
	inhabitant := &entity.Inhabitant{}
	err := s.db.Model(&entity.Inhabitant{}).Where(entity.Inhabitant{ID: inhabitantID}).First(inhabitant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return inhabitant, err
}
