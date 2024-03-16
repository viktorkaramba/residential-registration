package services

import (
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/errs"
	"time"
)

type Storages struct {
	User     UserStorage
	Building BuildingStorage
	OSBB     OSBBStorage
	Token    TokenStorage
}

type UserStorage interface {
	CreateUser(User *entity.User) error
	GetUser(UserID uint64, filter UserFilter) (*entity.User, error)
	ListUsers(filter UserFilter) ([]entity.User, error)
	UpdateUser(UserID, OSBBID uint64, user *entity.EventUserUpdatePayload) error
}

type BuildingStorage interface {
	GetByOSBBID(OSBBID uint64) (*entity.Building, error)
}

type OSBBStorage interface {
	CreateOSBB(OSBB *entity.OSBB) error
	ListOSBBS(filter OSBBFilter) ([]entity.OSBB, error)
	GetOSBB(filter OSBBFilter) (*entity.OSBB, error)
	CreateAnnouncement(announcement *entity.Announcement) error
	GetAnnouncement(AnnouncementID uint64, filter AnnouncementFilter) (*entity.Announcement, error)
	ListAnnouncements(filter AnnouncementFilter) ([]entity.Announcement, error)
	UpdateAnnouncement(AnnouncementID uint64, announcement *entity.EventAnnouncementUpdatePayload) error
	DeleteAnnouncement(AnnouncementID uint64, filter AnnouncementFilter) error
	CreatePoll(poll *entity.Poll) error
	ListPolls(filter PollFilter) ([]entity.Poll, error)
	UpdatePoll(PollID uint64, opts *entity.EventPollUpdatePayload) error
	DeletePoll(PollID uint64, filter PollFilter) error
	UpdateTestAnswer(TestAnswerID uint64, poll *entity.EventTestAnswerUpdatePayload) error
	DeleteTestAnswer(TestAnswerID uint64, filter TestAnswerFilter) error
	GetPoll(PollID uint64, filter PollFilter) (*entity.Poll, error)
	GetPollResult(PollID uint64, filter PollFilter) (*entity.PollResult, error)
	CreatAnswer(answer *entity.Answer) error
	ListAnswers(filter AnswerFilter) ([]entity.Answer, error)
	GetAnswer(AnswerID uint64, filter AnswerFilter) (*entity.Answer, error)
	UpdateAnswer(AnswerID, PollID uint64, answer *entity.EventUserAnswerUpdatePayload) error
	DeleteAnswer(AnswerID uint64, filter AnswerFilter) error
	CreatePayment(payment *entity.Payment) error
	CreateUserPayment(userPayment *entity.Purchase) error
}

type TokenStorage interface {
	CreateToken(token *entity.Token) error
	GetByToken(token string) (*entity.Token, error)
	Update(token *entity.Token) error
	UpdateByUser(token *entity.Token) error
}

type UserFilter struct {
	OSBBID *uint64
	*entity.PhoneNumber
	*entity.UserRole
}

type AnnouncementFilter struct {
	OSBBID *uint64
}

type PollFilter struct {
	OSBBID          *uint64
	WithTestAnswers bool
}

type TestAnswerFilter struct {
	PollID  *uint64
	Content *entity.Text
}

type AnswerFilter struct {
	PollID       *uint64
	UserID       *uint64
	TestAnswerID *uint64
	Content      *entity.Text
	CreatedAt    *time.Time
	UpdateAt     *time.Time
}

type OSBBFilter struct {
	OSBBID            *uint64
	WithBuilding      bool
	WithAnnouncements bool
	WithOSBBHead      bool
}

var (
	ErrPhoneNumberDuplicate = errs.M("user with this number already exist").Code("duplicate_phone_number")
	ErrEDRPOUDuplicate      = errs.M("osbb with this edrpou already exist").Code("duplicate_edrpou")
)
