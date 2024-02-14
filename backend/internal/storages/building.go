package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"

	"gorm.io/gorm"
)

type buildingStorage struct {
	db *gorm.DB
}

func NewBuildingStorage(db *gorm.DB) *buildingStorage {
	return &buildingStorage{
		db: db,
	}
}

func (s *buildingStorage) GetByOSBBID(OSBBID uint64) (*entity.Building, error) {
	building := &entity.Building{}
	err := s.db.Model(&entity.Building{}).Where(entity.Building{OSBBID: OSBBID}).First(building).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return building, err
}
