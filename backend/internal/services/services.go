package services

import (
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/logging"
)

type Services struct {
	User  UserService
	OSBB  OSBBService
	Token TokenService
}

type Options struct {
	Logger   logging.Logger
	Config   *config.Config
	Storages Storages
}

type UserService interface {
	AddUser(OSBBID uint64, inputUser entity.InputUser) (*entity.User, error)
}

type OSBBService interface {
	AddOSBB(inputOSSB entity.InputOSBB) (*entity.OSBB, error)
	AddAnnouncement(UserID, OSBBID uint64, inputAnnouncement entity.InputAnnouncement) (*entity.Announcement, error)
	AddPoll(UserID, OSBBID uint64, inputPoll entity.InputPoll) (*entity.Poll, error)
	AddPollTest(UserID, OSBBID uint64, inputPollTest entity.InputPollTest) (*entity.Poll, error)
	AddPollAnswer(UserID, PollID uint64, inputPollAnswer entity.InputPollAnswer) (*entity.Answer, error)
	AddPollAnswerTest(UserID, PollID uint64, inputPollAnswerTest entity.InputPollAnswerTest) (*entity.Answer, error)
}

type TokenService interface {
	GenerateToken(UserID uint64) (entity.TokenValue, error)
	GetByToken(token string) (*entity.Token, error)
	ParseToken(token string) (uint64, error)
	RefreshToken(UserID uint64) (string, error)
}
