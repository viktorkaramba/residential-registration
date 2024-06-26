package services

import (
	"crypto/sha1"
	"fmt"
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
	GetOSBB(UserID uint64) (*entity.OSBB, error)
	ListOSBBS() ([]entity.OSBB, error)
	UpdateOSBB(UserID uint64, input entity.EventOSBBUpdatePayload) error
	AddApartment(UserID, OSBBID uint64, inputApartment entity.EventApartmentPayload) (*entity.Apartment, error)
	AddAnnouncement(UserID, OSBBID uint64, inputAnnouncement entity.EventAnnouncementPayload) (*entity.Announcement, error)
	ListAnnouncements(UserID, OSBBID uint64) ([]entity.Announcement, error)
	UpdateAnnouncement(UserID, OSBBID, AnnouncementID uint64, input entity.EventAnnouncementUpdatePayload) error
	DeleteAnnouncement(UserID, OSBBID, AnnouncementID uint64) error
	AddPoll(UserID, OSBBID uint64, inputPoll entity.EventPollPayload) (*entity.Poll, error)
	AddPollTest(UserID, OSBBID uint64, inputPollTest entity.EventPollTestPayload) (*entity.Poll, error)
	ListPolls(UserID, OSBBID uint64) ([]entity.Poll, error)
	UpdatePoll(UserID, OSBBID, PollID uint64, poll entity.EventPollUpdatePayload) error
	DeletePoll(UserID, OSBBID, PollID uint64) error
	AddPollAnswer(UserID, PollID, OSBBID uint64, inputPollAnswer entity.EventPollAnswerPayload) (*entity.Answer, error)
	AddPollAnswerTest(UserID, PollID, OSBBID uint64, inputPollAnswerTest entity.EventPollAnswerTestPayload) (*entity.Answer, error)
	UpdateTestAnswer(UserID, OSBBID, PollID, TestAnswerID uint64, testAnswer entity.EventTestAnswerUpdatePayload) error
	DeleteTestAnswer(UserID, OSBBID, PollID, TestAnswerID uint64) error
	GetUserAnswer(UserID, OSBBID, PollID uint64) (*entity.Answer, error)
	UpdateAnswer(UserID, OSBBID, PollID uint64, answer *entity.EventUserAnswerUpdatePayload) error
	DeleteAnswer(UserID, OSBBID, PollID uint64) error
	GetPollResult(UserID, OSBBID, PollID uint64) (*entity.PollResult, error)
	AddPayment(UserID, OSBBID uint64, inputPayment entity.EventPaymentPayload) (*entity.Payment, error)
	ListPayments(UserID, OSBBID uint64) ([]entity.Payment, error)
	UpdatePayment(UserID, OSBBID, PaymentID uint64, inputPayment entity.EventPaymentUpdatePayload) error
	DeletePayment(UserID, OSBBID, PaymentID uint64) error
	UserUpdatePurchase(UserID, OSBBID, PaymentID uint64, inputPurchase entity.EventUserPurchaseUpdatePayload) error
	ListPurchasesByUser(UserID, OSBBID uint64, filter entity.EventPurchaseFilterPayload) ([]entity.EventUserPurchasesResponse, error)
	ListPurchasesByOSBBHead(UserID, OSBBID uint64, filter entity.EventPurchaseOSBBHeadFilterPayload) ([]entity.EventUserPurchasesResponse, error)
	UpdatePurchase(UserID, OSBBID, PaymentID uint64, inputPurchase entity.EventUserPurchaseUpdatePayload) error
	GetInhabitant(UserID uint64) (*entity.User, error)
	ListInhabitants(UserID, OSBBID uint64, filter UserFilter) ([]entity.User, error)
	UpdateInhabitant(UserID, OSBBID uint64, inhabitant entity.EventUserUpdatePayload) error
	ApproveInhabitant(UserID, OSBBID uint64, inhabitant entity.EventApproveUser) error
}

type TokenService interface {
	GenerateToken(UserID uint64) (entity.TokenValue, error)
	GetByToken(token string) (*entity.Token, error)
	ParseToken(token string) (uint64, error)
	RefreshToken(UserID uint64) (string, error)
}

func GeneratePasswordHash(salt, password string) entity.Password {
	hash := sha1.New()
	hash.Write([]byte(password))
	return entity.Password(fmt.Sprintf("%x", hash.Sum([]byte(salt))))
}
