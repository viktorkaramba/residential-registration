package services

import (
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/logging"
)

type Services struct {
	Auth  AuthService
	OSBB  OSBBService
	Token TokenService
}

type Options struct {
	Logger   logging.Logger
	Config   *config.Config
	Storages Storages
}

type AuthService interface {
	AddUser(OSBBID uint64, inputUser entity.EventUserPayload) (*entity.User, error)
	Login(inputLogin entity.EventLoginPayload) (*entity.User, error)
	Logout(token entity.TokenValue) error
}

type OSBBService interface {
	AddOSBB(inputOSSB entity.EventOSBBPayload) (*entity.OSBB, error)
	AddAnnouncement(UserID, OSBBID uint64, inputAnnouncement entity.EventAnnouncementPayload) (*entity.Announcement, error)
	ListAnnouncements(UserID, OSBBID uint64) ([]entity.Announcement, error)
	AddPoll(UserID, OSBBID uint64, inputPoll entity.EventPollPayload) (*entity.Poll, error)
	AddPollTest(UserID, OSBBID uint64, inputPollTest entity.EventPollTestPayload) (*entity.Poll, error)
	ListPolls(UserID, OSBBID uint64) ([]entity.Poll, error)
	AddPollAnswer(UserID, PollID uint64, inputPollAnswer entity.EventPollAnswerPayload) (*entity.Answer, error)
	AddPollAnswerTest(UserID, PollID uint64, inputPollAnswerTest entity.EventPollAnswerTestPayload) (*entity.Answer, error)
	GetPollResult(UserID, OSBBID, PollID uint64) (*entity.PollResult, error)
	AddPayment(UserID, OSBBID uint64, inputPayment entity.EventPaymentPayload) (*entity.Payment, error)
	AddPurchase(UserID, PaymentID uint64) (*entity.Purchase, error)
	GetInhabitant(UserID uint64) (*entity.User, error)
	ListInhabitants(UserID, OSBBID uint64) ([]entity.User, error)
	UpdateInhabitant(UserID, OSBBID uint64, inhabitant entity.EventUserUpdatePayload) error
}

type TokenService interface {
	GenerateToken(UserID uint64) (entity.TokenValue, error)
	GetByToken(token string) (*entity.Token, error)
	ParseToken(token string) (uint64, error)
	RefreshToken(UserID uint64) (string, error)
}
