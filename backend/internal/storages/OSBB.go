package storages

import (
	"residential-registration/backend/internal/entity"

	"gorm.io/gorm"
)

type OSBBStorage struct {
	db *gorm.DB
}

func NewOSBBStorage(db *gorm.DB) *OSBBStorage {
	return &OSBBStorage{
		db: db,
	}
}

func (s *OSBBStorage) CreateOSBB(OSBB *entity.OSBB) error {
	return s.db.Create(OSBB).Error
}
