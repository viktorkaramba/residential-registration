package services

import "residential-registration/backend/config"

type Services struct {
	User UserService
}

type Options struct {
	Config   *config.Config
	Storages Storages
}

type UserService interface {
}
