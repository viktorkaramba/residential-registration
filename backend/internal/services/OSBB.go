package services

import (
	"errors"
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"time"
)

type osbbService struct {
	config          *config.Config
	businessStorage Storages
}

func NewOSBBService(opts *Options) *osbbService {
	return &osbbService{
		config:          opts.Config,
		businessStorage: opts.Storages,
	}
}

func (s *osbbService) AddOSBB(inputOSBB entity.InputOSBB) (*entity.OSBB, error) {
	osbb := &entity.OSBB{
		Building: entity.Building{
			Address: inputOSBB.Address,
		},
		OSBBHead: entity.User{
			FullName:    inputOSBB.FullName,
			Password:    inputOSBB.Password,
			PhoneNumber: inputOSBB.PhoneNumber,
			Role:        entity.UserRoleOSBBHEad,
		},
		Name:   inputOSBB.Name,
		EDRPOU: inputOSBB.EDRPOU,
		Rent:   inputOSBB.Rent,
	}
	err := s.businessStorage.OSBB.CreateOSBB(osbb)
	if err != nil {
		return nil, err
	}

	return osbb, nil
}

func (s *osbbService) AddAnnouncement(UserID, OSBBID uint64, inputAnnouncement entity.InputAnnouncement) (*entity.Announcement, error) {
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		return nil, err
	}
	if user.Role == entity.UserRoleOSBBHEad {
		announcement := &entity.Announcement{
			UserID:    UserID,
			OSBBID:    OSBBID,
			Title:     inputAnnouncement.Title,
			Content:   inputAnnouncement.Content,
			CreatedAt: time.Now().UTC(),
		}
		err := s.businessStorage.OSBB.CreateAnnouncement(announcement)
		if err != nil {
			return nil, err
		}
		return announcement, nil
	}
	return nil, errors.New("user not osbb head")
}
