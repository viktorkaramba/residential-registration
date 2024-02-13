package services

import (
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/internal/storages"

	"gorm.io/gorm"
)

type Storages struct {
	Inhabitant InhabitantStorage
	OSBB       OSBBStorage
	Token      TokenStorage
}

func NewStorages(db *gorm.DB) *Storages {
	return &Storages{
		Inhabitant: storages.NewInhabitantStorage(db),
		OSBB:       storages.NewOSBBStorage(db),
		Token:      storages.NewTokenStorage(db),
	}
}

type InhabitantStorage interface {
	CreateInhabitant(inhabitant *entity.Inhabitant) error
	GetInhabitant(inhabitantID uint64) (*entity.Inhabitant, error)
}

type OSBBStorage interface {
	CreateOSBB(OSBB *entity.OSBB) error
}

type TokenStorage interface {
	CreateToken(token *entity.Token) error
	GetByToken(token string) (*entity.Token, error)
	Update(token *entity.Token) error
	UpdateByInhabitant(token *entity.Token) error
}
