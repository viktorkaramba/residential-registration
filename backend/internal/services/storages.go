package services

import (
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/internal/storages"

	"gorm.io/gorm"
)

type Storages struct {
	User     UserStorage
	Building BuildingStorage
	OSBB     OSBBStorage
	Token    TokenStorage
}

func NewStorages(db *gorm.DB) *Storages {
	return &Storages{
		User:     storages.NewUserStorage(db),
		Building: storages.NewBuildingStorage(db),
		OSBB:     storages.NewOSBBStorage(db),
		Token:    storages.NewTokenStorage(db),
	}
}

type UserStorage interface {
	CreateUser(User *entity.User) error
	GetUser(UserID uint64) (*entity.User, error)
	GetUserByPhoneNumber(phoneNumber entity.PhoneNumber) (*entity.User, error)
}

type BuildingStorage interface {
	GetByOSBBID(OSBBID uint64) (*entity.Building, error)
}

type OSBBStorage interface {
	CreateOSBB(OSBB *entity.OSBB) error
	GetOSBB(OSBBID uint64) (*entity.OSBB, error)
	CreateAnnouncement(announcement *entity.Announcement) error
	CreatePoll(poll *entity.Poll) error
	GetPoll(PollID uint64) (*entity.Poll, error)
	CreatAnswer(answer *entity.Answer) error
}

type TokenStorage interface {
	CreateToken(token *entity.Token) error
	GetByToken(token string) (*entity.Token, error)
	Update(token *entity.Token) error
	UpdateByUser(token *entity.Token) error
}
