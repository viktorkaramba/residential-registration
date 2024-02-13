package services

import (
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
)

type Services struct {
	Inhabitant InhabitantService
	OSBB       OSBBService
	Token      TokenService
}

type Options struct {
	Config   *config.Config
	Storages Storages
}

type InhabitantService interface {
	AddInhabitant(OSBBID uint64, inputInhabitant entity.InputInhabitant) (*entity.Inhabitant, error)
}

type OSBBService interface {
}

type TokenService interface {
	GenerateToken(inhabitantID uint64) (entity.TokenValue, error)
	GetByToken(token string) (*entity.Token, error)
	ParseToken(token string) (uint64, error)
	RefreshToken(inhabitantID uint64) (string, error)
}
