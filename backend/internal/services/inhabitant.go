package services

import (
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
)

type inhabitantService struct {
	config          *config.Config
	businessStorage Storages
}

func NewInhabitantService(opts *Options) *inhabitantService {
	return &inhabitantService{
		config:          opts.Config,
		businessStorage: opts.Storages,
	}
}

func (s *inhabitantService) AddInhabitant(BuildingID uint64, inputInhabitant entity.InputInhabitant) (*entity.Inhabitant, error) {
	inhabitant := &entity.Inhabitant{
		Apartment: entity.Apartment{
			BuildingID: BuildingID,
			Number:     inputInhabitant.ApartmentNumber,
			Area:       inputInhabitant.ApartmentArea,
		},
		FullName: inputInhabitant.FullName,
		Password: inputInhabitant.Password,
	}
	err := s.businessStorage.Inhabitant.CreateInhabitant(inhabitant)
	if err != nil {
		return nil, err
	}

	return inhabitant, nil
}
