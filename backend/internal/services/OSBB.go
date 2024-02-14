package services

import (
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
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
