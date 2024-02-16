package services

import (
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/errs"
	"residential-registration/backend/pkg/logging"
	"time"
)

type osbbService struct {
	logger          logging.Logger
	config          *config.Config
	businessStorage Storages
}

func NewOSBBService(opts *Options) *osbbService {
	return &osbbService{
		logger:          opts.Logger.Named("OSBBService"),
		config:          opts.Config,
		businessStorage: opts.Storages,
	}
}

func (s *osbbService) AddOSBB(inputOSBB entity.EventOSBBPayload) (*entity.OSBB, error) {
	logger := s.logger.Named("AddOSBB").
		With("input_osbb", inputOSBB)
	osbb := &entity.OSBB{
		Building: entity.Building{
			Address: inputOSBB.Address,
		},
		OSBBHead: entity.User{
			FullName: entity.FullName{
				FirstName:  inputOSBB.FirstName,
				Surname:    inputOSBB.Surname,
				Patronymic: inputOSBB.Patronymic,
			},
			Password:    inputOSBB.Password,
			PhoneNumber: inputOSBB.PhoneNumber,
			Role:        entity.UserRoleOSBBHEad,
		},
		Name:   inputOSBB.Name,
		EDRPOU: inputOSBB.EDRPOU,
		Rent:   inputOSBB.Rent,
	}
	err := s.businessStorage.OSBB.CreateOSBB(osbb)
	if err != nil {
		logger.Error("failed to create osbb", "error", err)
		return nil, errs.Err(err).Code("Failed to create osbb").Kind(errs.Database)
	}

	return osbb, nil
}

func (s *osbbService) AddAnnouncement(UserID, OSBBID uint64, inputAnnouncement entity.EventAnnouncementPayload) (*entity.Announcement, error) {
	logger := s.logger.Named("AddAnnouncement").
		With("user_id", UserID).With("osbb_id", OSBBID).With("input_announcement", inputAnnouncement)
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create an announcement", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create an announcement").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBID)
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	announcement := &entity.Announcement{
		UserID:    UserID,
		OSBBID:    OSBBID,
		Title:     inputAnnouncement.Title,
		Content:   inputAnnouncement.Content,
		CreatedAt: time.Now().UTC(),
	}
	err = s.businessStorage.OSBB.CreateAnnouncement(announcement)
	if err != nil {
		logger.Error("failed to сreate announcement", "error", err)
		return nil, errs.Err(err).Code("Failed to сreate announcement").Kind(errs.Database)
	}
	return announcement, nil
}

func (s *osbbService) AddPoll(UserID, OSBBID uint64, inputPoll entity.EventPollPayload) (*entity.Poll, error) {
	logger := s.logger.Named("AddPoll").
		With("user_id", UserID).With("osbb_id", OSBBID).With("input_poll", inputPoll)
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBID)
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	poll := &entity.Poll{
		UserID:     UserID,
		OSBBID:     OSBBID,
		Question:   inputPoll.Question,
		CreatedAt:  time.Now().UTC(),
		FinishedAt: inputPoll.FinishedAt,
		Type:       entity.PollTypeOpenAnswer,
	}
	err = s.businessStorage.OSBB.CreatePoll(poll)
	if err != nil {
		logger.Error("failed to сreate poll", "error", err)
		return nil, errs.Err(err).Code("Failed to сreate poll").Kind(errs.Database)
	}
	return poll, nil
}

func (s *osbbService) AddPollTest(UserID, OSBBID uint64, inputPollTest entity.EventPollTestPayload) (*entity.Poll, error) {
	logger := s.logger.Named("AddPollTest").
		With("user_id", UserID).With("osbb_id", OSBBID).With("input_poll_test", inputPollTest)
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll test", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll test").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBID)
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	poll := &entity.Poll{
		UserID:     UserID,
		OSBBID:     OSBBID,
		Question:   inputPollTest.Question,
		TestAnswer: inputPollTest.TestAnswer,
		CreatedAt:  time.Now().UTC(),
		FinishedAt: inputPollTest.FinishedAt,
		Type:       entity.PollTypeTest,
	}
	err = s.businessStorage.OSBB.CreatePoll(poll)
	if err != nil {
		logger.Error("failed to сreate poll test", "error", err)
		return nil, errs.Err(err).Code("Failed to сreate poll test").Kind(errs.Database)
	}
	return poll, nil
}

func (s *osbbService) AddPollAnswer(UserID, PollID uint64, inputPollAnswer entity.EventPollAnswerPayload) (*entity.Answer, error) {
	logger := s.logger.Named("AddPollAnswer").
		With("user_id", UserID).With("poll_id", PollID).With("input_poll_answer", inputPollAnswer)
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll answer", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll answer").Kind(errs.Private)
	}
	poll, err := s.businessStorage.OSBB.GetPoll(PollID)
	if err != nil {
		logger.Error("poll do not exist", "error", err)
		return nil, errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist", "error", err)
		return nil, errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}
	answer := &entity.Answer{
		PollID:  PollID,
		UserID:  UserID,
		Content: inputPollAnswer.Content,
	}
	err = s.businessStorage.OSBB.CreatAnswer(answer)
	if err != nil {
		logger.Error("failed to сreate answer", "error", err)
		return nil, errs.Err(err).Code("Failed to сreate answer").Kind(errs.Database)
	}
	return answer, nil
}

func (s *osbbService) AddPollAnswerTest(UserID, PollID uint64, inputPollAnswerTest entity.EventPollAnswerTestPayload) (*entity.Answer, error) {
	logger := s.logger.Named("AddPollAnswerTest").
		With("user_id", UserID).With("poll_id", PollID).With("input_poll_answer_test", inputPollAnswerTest)

	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll answer", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll answer").Kind(errs.Private)
	}
	poll, err := s.businessStorage.OSBB.GetPoll(PollID)
	if err != nil {
		logger.Error("poll do not exist", "error", err)
		return nil, errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist", "error", err)
		return nil, errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}
	var isExist bool
	for _, answer := range poll.TestAnswer {
		if answer.PollID == PollID && answer.ID == inputPollAnswerTest.TestAnswerID {
			isExist = true
		}
	}
	if !isExist {
		logger.Error("test answer do not exist", "error", err)
		return nil, errs.M("test answer not found").Code("Test answer do not exist").Kind(errs.Database)
	}
	answer := &entity.Answer{
		PollID:       PollID,
		UserID:       UserID,
		TestAnswerID: inputPollAnswerTest.TestAnswerID,
	}
	err = s.businessStorage.OSBB.CreatAnswer(answer)
	if err != nil {
		logger.Error("failed to сreate test answer", "error", err)
		return nil, errs.Err(err).Code("Failed to сreate test answer").Kind(errs.Database)
	}
	return answer, nil
}
