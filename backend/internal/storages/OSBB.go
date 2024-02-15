package storages

import (
	"errors"
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

func (s *OSBBStorage) GetOSBB(OSBBID uint64) (*entity.OSBB, error) {
	osbb := &entity.OSBB{}
	err := s.db.Model(&entity.OSBB{}).Where(entity.OSBB{ID: OSBBID}).First(osbb).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return osbb, err
}

func (s *OSBBStorage) CreateAnnouncement(announcement *entity.Announcement) error {
	return s.db.Create(announcement).Error
}

func (s *OSBBStorage) CreatePoll(poll *entity.Poll) error {
	return s.db.Create(poll).Error
}

func (s *OSBBStorage) GetPoll(PollID uint64) (*entity.Poll, error) {
	poll := &entity.Poll{}
	err := s.db.Model(&entity.Poll{}).Where(entity.Poll{ID: PollID}).First(poll).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return poll, err
}

func (s *OSBBStorage) CreatAnswer(answer *entity.Answer) error {
	return s.db.Create(answer).Error
}
