package services

import (
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/database"
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

func (s *osbbService) ListOSBBS() ([]entity.OSBB, error) {
	logger := s.logger.Named("ListOSBBS")
	osbbs, err := s.businessStorage.OSBB.ListOSBBS(OSBBFilter{
		WithBuilding: true,
		WithOSBBHead: true,
	})
	if err != nil {
		logger.Error("failed to get list osbbs", "error", err)
		return nil, errs.Err(err).Code("Failed to get list osbb").Kind(errs.Database)
	}

	return osbbs, nil
}

func (s *osbbService) GetOSBB(UserID uint64) (*entity.OSBB, error) {
	logger := s.logger.Named("GetOSBB").With("user_id", UserID)
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{
		OSBBID: &user.OSBBID, WithOSBBHead: true, WithBuilding: true,
	})
	if err != nil {
		logger.Error("failed to get list osbbs", "error", err)
		return nil, errs.Err(err).Code("Failed to get list osbb").Kind(errs.Database)
	}

	return osbb, nil
}

func (s *osbbService) AddAnnouncement(UserID, OSBBID uint64, inputAnnouncement entity.EventAnnouncementPayload) (*entity.Announcement, error) {
	logger := s.logger.Named("AddAnnouncement").
		With("user_id", UserID).With("osbb_id", OSBBID).With("input_announcement", inputAnnouncement)
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create an announcement", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create an announcement").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
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

func (s *osbbService) ListAnnouncements(UserID, OSBBID uint64) ([]entity.Announcement, error) {
	logger := s.logger.Named("ListAnnouncements").
		With("user_id", UserID).With("osbb_id", OSBBID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	announcements, err := s.businessStorage.OSBB.ListAnnouncements(AnnouncementFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get list announcements", "error", err)
		return nil, errs.Err(err).Code("Failed to get list announcements").Kind(errs.Database)
	}
	return announcements, nil
}

func (s *osbbService) AddPoll(UserID, OSBBID uint64, inputPoll entity.EventPollPayload) (*entity.Poll, error) {
	logger := s.logger.Named("AddPoll").
		With("user_id", UserID).With("osbb_id", OSBBID).With("input_poll", inputPoll)
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
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
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll test", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll test").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	poll := &entity.Poll{
		UserID:      UserID,
		OSBBID:      OSBBID,
		Question:    inputPollTest.Question,
		TestAnswers: inputPollTest.TestAnswer,
		CreatedAt:   time.Now().UTC(),
		FinishedAt:  inputPollTest.FinishedAt,
		Type:        entity.PollTypeTest,
	}
	err = s.businessStorage.OSBB.CreatePoll(poll)
	if err != nil {
		logger.Error("failed to сreate poll test", "error", err)
		return nil, errs.Err(err).Code("Failed to сreate poll test").Kind(errs.Database)
	}
	return poll, nil
}

func (s *osbbService) ListPolls(UserID, OSBBID uint64) ([]entity.Poll, error) {
	logger := s.logger.Named("ListPolls").
		With("user_id", UserID).With("osbb_id", OSBBID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err == nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user != nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	polls, err := s.businessStorage.OSBB.ListPolls(PollFilter{OSBBID: &OSBBID, WithTestAnswers: true})
	if err != nil {
		logger.Error("failed to get list polls", "error", err)
		return nil, errs.Err(err).Code("Failed to get list polls").Kind(errs.Database)
	}
	return polls, nil
}

func (s *osbbService) AddPollAnswer(UserID, PollID, OSBBID uint64, inputPollAnswer entity.EventPollAnswerPayload) (*entity.Answer, error) {
	logger := s.logger.Named("AddPollAnswer").
		With("user_id", UserID).With("poll_id", PollID).With("osbb_id", OSBBID).
		With("input_poll_answer", inputPollAnswer)
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID, WithTestAnswers: false})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return nil, errs.M("failed to get poll").Code("Failed to get poll").Kind(errs.Database)
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

func (s *osbbService) AddPollAnswerTest(UserID, PollID, OSBBID uint64, inputPollAnswerTest entity.EventPollAnswerTestPayload) (*entity.Answer, error) {
	logger := s.logger.Named("AddPollAnswerTest").
		With("user_id", UserID).With("poll_id", PollID).With("osbb_id", PollID).
		With("input_poll_answer_test", inputPollAnswerTest)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID, WithTestAnswers: true})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return nil, errs.M("failed to get poll").Code("Failed to get poll").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist", "error", err)
		return nil, errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}
	var isExist bool
	for _, answer := range poll.TestAnswers {
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

func (s *osbbService) GetPollResult(UserID, OSBBID, PollID uint64) (*entity.PollResult, error) {
	logger := s.logger.Named("ListPollsAnswers").
		With("user_id", UserID).With("osbb_id", OSBBID).With("poll_id", PollID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	pollResult, err := s.businessStorage.OSBB.GetPollResult(PollID)
	if err != nil {
		logger.Error("failed to get polls", "error", err)
		return nil, errs.Err(err).Code("Failed to get list polls").Kind(errs.Database)
	}
	return pollResult, nil
}

func (s *osbbService) AddPayment(UserID, OSBBID uint64, inputPayment entity.EventPaymentPayload) (*entity.Payment, error) {
	logger := s.logger.Named("AddPayment").
		With("user_id", UserID).With("osbb_id", OSBBID).With("input_payment", inputPayment)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll answer", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll answer").Kind(errs.Private)
	}
	payment := &entity.Payment{
		OSBBID:      OSBBID,
		CreatedAt:   time.Now().UTC(),
		Deadline:    inputPayment.Deadline,
		Amount:      inputPayment.Amount,
		Appointment: inputPayment.Appointment,
	}
	err = s.businessStorage.OSBB.CreatePayment(payment)
	if err != nil {
		logger.Error("failed to create payment", "error", err)
		return nil, errs.Err(err).Code("Failed to create payment").Kind(errs.Database)
	}
	return payment, nil
}

func (s *osbbService) AddPurchase(UserID, PaymentID uint64) (*entity.Purchase, error) {
	logger := s.logger.Named("AddPayment").
		With("user_id", UserID).With("payment_id", PaymentID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll answer", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll answer").Kind(errs.Private)
	}
	userPayment := &entity.Purchase{
		PaymentID:       PaymentID,
		UserID:          UserID,
		PaymentStatus:   entity.Paid,
		PostgreSQLModel: database.PostgreSQLModel{},
	}
	err = s.businessStorage.OSBB.CreateUserPayment(userPayment)
	if err != nil {
		logger.Error("failed to create payment", "error", err)
		return nil, errs.Err(err).Code("Failed to create payment").Kind(errs.Database)
	}
	return userPayment, nil
}

func (s *osbbService) GetInhabitant(UserID uint64) (*entity.User, error) {
	logger := s.logger.Named("GetInhabitant").With("user_id", UserID)

	inhabitant, err := s.businessStorage.User.GetUser(UserID, UserFilter{})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}

	return inhabitant, nil
}

func (s *osbbService) ListInhabitants(UserID, OSBBID uint64) ([]entity.User, error) {
	logger := s.logger.Named("ListInhabitans").
		With("user_id", UserID).With("osbb_id", OSBBID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHEad {
		logger.Error("User can not create a poll answer", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll answer").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	inhabitants, err := s.businessStorage.User.ListUsers(UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get list users", "error", err)
		return nil, errs.Err(err).Code("Failed to get list users").Kind(errs.Database)
	}
	return inhabitants, nil
}

func (s *osbbService) UpdateInhabitant(UserID, OSBBID uint64, inhabitant entity.EventUserUpdatePayload) error {
	logger := s.logger.Named("UpdateInhabitant").
		With("user_id", UserID).With("inhabitant", inhabitant)
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if err := inhabitant.Validate(); err != nil {
		logger.Error("failed to validate update inhabitant data", "error", err)
		return errs.Err(err).Code("failed to validate update inhabitant data").Kind(errs.Validation)
	}
	err = s.businessStorage.User.UpdateUser(UserID, OSBBID, &inhabitant)
	if err != nil {
		logger.Error("failed to update inhabitant", "error", err)
		return errs.Err(err).Code("Failed to update inhabitant").Kind(errs.Database)
	}

	return nil
}
