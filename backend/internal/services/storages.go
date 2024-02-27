package services

import (
	"residential-registration/backend/internal/entity"
)

type Storages struct {
	User     UserStorage
	Building BuildingStorage
	OSBB     OSBBStorage
	Token    TokenStorage
}

type UserStorage interface {
	CreateUser(User *entity.User) error
	GetUser(UserID uint64) (*entity.User, error)
	GetUserByPhoneNumber(phoneNumber entity.PhoneNumber) (*entity.User, error)
	ListUsers(filter ListUserFilter) ([]entity.User, error)
	UpdateUser(UserID, OSBBID uint64, user *entity.EventUserUpdatePayload) error
}

type BuildingStorage interface {
	GetByOSBBID(OSBBID uint64) (*entity.Building, error)
}

type OSBBStorage interface {
	CreateOSBB(OSBB *entity.OSBB) error
	GetOSBB(OSBBID uint64) (*entity.OSBB, error)
	CreateAnnouncement(announcement *entity.Announcement) error
	ListAnnouncements(filter ListAnnouncementFilter) ([]entity.Announcement, error)
	CreatePoll(poll *entity.Poll) error
	GetPoll(PollID uint64) (*entity.Poll, error)
	CreatAnswer(answer *entity.Answer) error
	CreatePayment(payment *entity.Payment) error
	CreateUserPayment(userPayment *entity.Purchase) error
}

type TokenStorage interface {
	CreateToken(token *entity.Token) error
	GetByToken(token string) (*entity.Token, error)
	Update(token *entity.Token) error
	UpdateByUser(token *entity.Token) error
}

type ListUserFilter struct {
	OSBBID *uint64
}

type ListAnnouncementFilter struct {
	OSBBID *uint64
}
