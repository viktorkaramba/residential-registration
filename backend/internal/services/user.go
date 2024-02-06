package services

import "residential-registration/backend/config"

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
