package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/internal/services"

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

func (s *buildingStorage) UpdateBuilding(BuildingID uint64, filter services.BuildingFilter) error {
	stmt := s.db.Model(&entity.Building{})
	var building entity.Building

	if BuildingID != 0 {
		stmt = stmt.Where("id = ?", BuildingID)
	}

	if filter.OSBBID != nil {
		stmt = stmt.Where("osbb_id = ?", *filter.OSBBID)
	}
	if filter.Address != nil {
		building.Address = *filter.Address
	}
	return stmt.Updates(building).Error
}
