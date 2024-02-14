package services

import (
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
)

type userService struct {
	config          *config.Config
	businessStorage Storages
}

func NewUserService(opts *Options) *userService {
	return &userService{
		config:          opts.Config,
		businessStorage: opts.Storages,
	}
}

func (s *userService) AddUser(OSBBID uint64, inputUser entity.InputUser) (*entity.User, error) {

	building, err := s.businessStorage.Building.GetByOSBBID(OSBBID)

	if err != nil {
		return nil, err
	}

	User := &entity.User{
		Apartment: entity.Apartment{
			BuildingID: building.ID,
			Number:     inputUser.ApartmentNumber,
			Area:       inputUser.ApartmentArea,
		},
		OSBBID:   OSBBID,
		FullName: inputUser.FullName,
		Password: inputUser.Password,
		Role:     entity.UserRoleInhabitant,
	}
	err = s.businessStorage.User.CreateUser(User)
	if err != nil {
		return nil, err
	}

	return User, nil
}
