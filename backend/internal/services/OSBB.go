package services

import (
	"errors"
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/database"
	"residential-registration/backend/pkg/errs"
	"residential-registration/backend/pkg/logging"
	"residential-registration/backend/pkg/typecast"
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
			Password:    GeneratePasswordHash(s.config.Salt, string(inputOSBB.Password)),
			PhoneNumber: inputOSBB.PhoneNumber,
			Role:        entity.UserRoleOSBBHead,
			IsApproved:  typecast.ToPtr(true),
		},
		Name:   inputOSBB.Name,
		Photo:  inputOSBB.Photo,
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
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{
		OSBBID: &user.OSBBID, WithOSBBHead: true, WithBuilding: true,
	})
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	if err != nil {
		logger.Error("failed to get list osbbs", "error", err)
		return nil, errs.Err(err).Code("Failed to get list osbb").Kind(errs.Database)
	}

	osbbHead, err := s.businessStorage.User.GetUser(0, UserFilter{
		OSBBID:   &osbb.ID,
		UserRole: typecast.ToPtr(entity.UserRoleOSBBHead),
	})

	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	osbb.OSBBHead = *osbbHead
	return osbb, nil
}

func (s *osbbService) UpdateOSBB(UserID uint64, input entity.EventOSBBUpdatePayload) error {
	logger := s.logger.Named("UpdateOSBB").
		With("user_id", UserID).With("input", input)
	if err := input.Validate(); err != nil {
		logger.Error("failed to validate update announcement data", "error", err)
		return errs.Err(err).Code("failed to validate update announcement data").Kind(errs.Validation)
	}
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not update an announcement", "error", err)
		return errs.M("user not osbb head").Code("User can not update an announcement").Kind(errs.Private)
	}

	err = s.businessStorage.OSBB.UpdateOSBB(user.OSBBID, &input)
	if err != nil {
		logger.Error("failed to update osbb", "error", err)
		return errs.Err(err).Code("Failed to update osbb").Kind(errs.Database)
	}

	if input.Address != nil {
		err = s.businessStorage.Building.UpdateBuilding(0, BuildingFilter{
			OSBBID:  &user.OSBBID,
			Address: input.Address,
		})
		if err != nil {
			logger.Error("failed to update building", "error", err)
			return errs.Err(err).Code("Failed to update building").Kind(errs.Database)
		}
	}
	return nil
}

func (s *osbbService) AddApartment(UserID, OSBBID uint64, inputApartment entity.EventApartmentPayload) (*entity.Apartment, error) {
	logger := s.logger.Named("AddApartment").
		With("user_id", UserID).With("osbb_id", OSBBID).With("input_apartment", inputApartment)
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not create an announcement", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create an announcement").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	apartment := &entity.Apartment{
		BuildingID: OSBBID,
		UserID:     UserID,
		Number:     inputApartment.Number,
		Area:       inputApartment.Area,
	}
	err = s.businessStorage.OSBB.CreateApartment(apartment)
	if err != nil {
		logger.Error("failed to сreate apartment", "error", err)
		return nil, errs.Err(err).Code("Failed to сreate apartment").Kind(errs.Database)
	}
	return apartment, nil
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
	if user.Role != entity.UserRoleOSBBHead {
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

func (s *osbbService) UpdateAnnouncement(UserID, OSBBID, AnnouncementID uint64, input entity.EventAnnouncementUpdatePayload) error {
	logger := s.logger.Named("UpdateAnnouncement").
		With("user_id", UserID).With("osbb_id", OSBBID).With("announcement_id", AnnouncementID).
		With("announcement", input)
	if err := input.Validate(); err != nil {
		logger.Error("failed to validate update announcement data", "error", err)
		return errs.Err(err).Code("failed to validate update announcement data").Kind(errs.Validation)
	}
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not update an announcement", "error", err)
		return errs.M("user not osbb head").Code("User can not update an announcement").Kind(errs.Private)
	}

	announcement, err := s.businessStorage.OSBB.GetAnnouncement(AnnouncementID, AnnouncementFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get announcement", "error", err)
		return errs.Err(err).Code("Failed to get announcement").Kind(errs.Database)
	}
	if announcement == nil {
		logger.Error("announcement do not exist in current osbb", "error", err)
		return errs.M("announcement not found in current osbb").Code("announcement do not exist").Kind(errs.Database)
	}

	err = s.businessStorage.OSBB.UpdateAnnouncement(AnnouncementID, &input)
	if err != nil {
		logger.Error("failed to update announcement", "error", err)
		return errs.Err(err).Code("Failed to update announcement").Kind(errs.Database)
	}

	return nil
}

func (s *osbbService) DeleteAnnouncement(UserID, OSBBID, AnnouncementID uint64) error {
	logger := s.logger.Named("DeleteAnnouncement").
		With("user_id", UserID).With("osbb_id", OSBBID).With("announcement_id", AnnouncementID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not delete a announcement", "error", err)
		return errs.M("user not osbb head").Code("User can not delete a announcement").Kind(errs.Private)
	}

	announcement, err := s.businessStorage.OSBB.GetAnnouncement(AnnouncementID, AnnouncementFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get announcement", "error", err)
		return errs.Err(err).Code("Failed to get announcement").Kind(errs.Database)
	}
	if announcement == nil {
		logger.Error("announcement do not exist in current osbb", "error", err)
		return errs.M("announcement not found in current osbb").Code("announcements do not exist").Kind(errs.Database)
	}
	err = s.businessStorage.OSBB.DeleteAnnouncement(AnnouncementID, AnnouncementFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to delete announcement", "error", err)
		return errs.Err(err).Code("Failed to delete announcement").Kind(errs.Database)
	}

	return nil
}

func (s *osbbService) AddPoll(UserID, OSBBID uint64, inputPoll entity.EventPollPayload) (*entity.Poll, error) {
	logger := s.logger.Named("AddPoll").
		With("user_id", UserID).With("osbb_id", OSBBID).With("input_poll", inputPoll)
	if inputPoll.FinishedAt.Compare(time.Now()) == -1 {
		logger.Error("failed to update poll", "error", errors.New("finished at must be after current time"))
		return nil, errs.M("finished at must be after current time").Code("Failed to update poll ").Kind(errs.Database)
	}
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
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
		IsClosed:   false,
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
		With("user_id", UserID).With("osbb_id", OSBBID).With("poll_id").
		With("input_poll_test", inputPollTest)
	if len(inputPollTest.TestAnswer) < 2 {
		logger.Error("failed to add poll test", "error", errors.New("count of test answers must be greater than 2"))
		return nil, errs.M("count of test answers must be greater than 1").Code("Failed to add poll test").Kind(errs.Database)
	}
	if inputPollTest.FinishedAt.Compare(time.Now()) == -1 {
		logger.Error("failed to update poll", "error", errors.New("finished at must be after current time"))
		return nil, errs.M("finished at must be after current time").Code("Failed to update poll ").Kind(errs.Database)
	}
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
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
	polls, err := s.businessStorage.OSBB.ListPolls(PollFilter{OSBBID: &OSBBID, WithTestAnswers: true})
	if err != nil {
		logger.Error("failed to get list polls", "error", err)
		return nil, errs.Err(err).Code("Failed to get list polls").Kind(errs.Database)
	}
	return polls, nil
}

func (s *osbbService) UpdatePoll(UserID, OSBBID, PollID uint64, input entity.EventPollUpdatePayload) error {
	logger := s.logger.Named("UpdatePoll").
		With("user_id", UserID).With("osbb_id", OSBBID).With("poll_id", PollID).
		With("poll", input)
	if err := input.Validate(); err != nil {
		logger.Error("failed to validate update poll data", "error", err)
		return errs.Err(err).Code("failed to validate update poll data").Kind(errs.Validation)
	}
	if input.FinishedAt.Compare(time.Now()) == -1 {
		logger.Error("failed to update poll", "error", errors.New("finished at must be after current time"))
		return errs.M("finished at must be after current time").Code("Failed to update poll ").Kind(errs.Database)
	}
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not update a poll", "error", err)
		return errs.M("user not osbb head").Code("User can not update a poll").Kind(errs.Private)
	}
	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return errs.Err(err).Code("Failed to get poll").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist in current osbb", "error", err)
		return errs.M("poll not found in current osbb").Code("poll do not exist").Kind(errs.Database)
	}
	err = s.businessStorage.OSBB.UpdatePoll(PollID, &input)
	if err != nil {
		logger.Error("failed to update poll", "error", err)
		return errs.Err(err).Code("Failed to update poll").Kind(errs.Database)
	}

	return nil
}

func (s *osbbService) DeletePoll(UserID, OSBBID, PollID uint64) error {
	logger := s.logger.Named("DeletePoll").
		With("user_id", UserID).With("osbb_id", OSBBID).With("poll_id", PollID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not delete a poll", "error", err)
		return errs.M("user not osbb head").Code("User can not delete a poll").Kind(errs.Private)
	}

	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return errs.Err(err).Code("Failed to get poll").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist in current osbb", "error", err)
		return errs.M("poll not found in current osbb").Code("poll do not exist").Kind(errs.Database)
	}
	err = s.businessStorage.OSBB.DeletePoll(PollID, PollFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to delete poll", "error", err)
		return errs.Err(err).Code("Failed to delete poll").Kind(errs.Database)
	}

	return nil
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
	isAnswerAlreadyExist, err := s.businessStorage.OSBB.GetAnswer(0, AnswerFilter{
		PollID: &PollID,
		UserID: &UserID,
	})
	if err != nil {
		logger.Error("failed to get answer", "error", err)
		return nil, errs.Err(err).Code("Failed to get answer").Kind(errs.Database)
	}
	if isAnswerAlreadyExist != nil {
		logger.Error("answer already exist", "error", err)
		err = s.businessStorage.OSBB.UpdateAnswer(isAnswerAlreadyExist.ID, 0, &entity.EventUserAnswerUpdatePayload{
			TestAnswerID: &inputPollAnswerTest.TestAnswerID,
		})
		if err != nil {
			logger.Error("failed to update answer", "error", err)
			return nil, errs.Err(err).Code("Failed to update answer").Kind(errs.Database)
		}
		return isAnswerAlreadyExist, nil
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

func (s *osbbService) UpdateTestAnswer(UserID, OSBBID, PollID, TestAnswerID uint64, testAnswer entity.EventTestAnswerUpdatePayload) error {
	logger := s.logger.Named("UpdateTestAnswer").
		With("user_id", UserID).With("osbb_id", OSBBID).With("test_answer_id", TestAnswerID).
		With("test_answer", testAnswer)
	if err := testAnswer.Validate(); err != nil {
		logger.Error("failed to validate update poll data", "error", err)
		return errs.Err(err).Code("failed to validate update poll data").Kind(errs.Validation)
	}
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not update a test answer", "error", err)
		return errs.M("user not osbb head").Code("User can not update a test answer").Kind(errs.Private)
	}
	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID, WithTestAnswers: true})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return errs.M("failed to get poll").Code("Failed to get poll").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist", "error", err)
		return errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}

	var isExist bool
	for _, answer := range poll.TestAnswers {
		if answer.PollID == PollID && answer.ID == TestAnswerID {
			isExist = true
		}
	}
	if isExist {
		err = s.businessStorage.OSBB.UpdateTestAnswer(TestAnswerID, &testAnswer)
		if err != nil {
			logger.Error("failed to update test answer", "error", err)
			return errs.Err(err).Code("Failed to update test answer").Kind(errs.Database)
		}
	} else {
		logger.Error("test answer do not exist", "error", err)
		return errs.M("test answer  not found").Code("Test answer  do not exist").Kind(errs.Database)
	}

	return nil
}

func (s *osbbService) DeleteTestAnswer(UserID, OSBBID, PollID, TestAnswerID uint64) error {
	logger := s.logger.Named("DeleteTestAnswer").
		With("user_id", UserID).With("osbb_id", OSBBID).With("test_answer_id", TestAnswerID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not delete an test answer", "error", err)
		return errs.M("user not osbb head").Code("User can not delete a test answer").Kind(errs.Private)
	}
	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID, WithTestAnswers: true})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return errs.M("failed to get poll").Code("Failed to get poll").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist", "error", err)
		return errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}
	if len(poll.TestAnswers) < 3 {
		logger.Error("failed to add poll test", "error", errors.New("count of test answers must be greater than 2"))
		return errs.M("count of test answers must be greater than 1").Code("Failed to add poll test").Kind(errs.Database)
	}
	var isExist bool
	for _, answer := range poll.TestAnswers {
		if answer.PollID == PollID && answer.ID == TestAnswerID {
			isExist = true
			break
		}
	}

	if isExist {
		err = s.businessStorage.OSBB.DeleteTestAnswer(TestAnswerID, TestAnswerFilter{})
		if err != nil {
			logger.Error("failed to delete test answer", "error", err)
			return errs.Err(err).Code("Failed to delete test answer").Kind(errs.Database)
		}
	} else {
		logger.Error("test answer do not exist", "error", err)
		return errs.M("test answer  not found").Code("Test answer  do not exist").Kind(errs.Database)
	}
	return nil
}

func (s *osbbService) GetUserAnswer(UserID, OSBBID, PollID uint64) (*entity.Answer, error) {
	logger := s.logger.Named("GetUserAnswer").
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
	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID, WithTestAnswers: false})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return nil, errs.M("failed to get poll").Code("Failed to get poll").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist", "error", err)
		return nil, errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}

	userAnswers, err := s.businessStorage.OSBB.GetAnswer(0, AnswerFilter{
		PollID: &PollID,
		UserID: &UserID,
	})
	if err != nil {
		logger.Error("failed to get list answers", "error", err)
		return nil, errs.Err(err).Code("Failed to get list answers").Kind(errs.Database)
	}
	return userAnswers, nil
}

func (s *osbbService) UpdateAnswer(UserID, OSBBID, PollID uint64, answer *entity.EventUserAnswerUpdatePayload) error {
	logger := s.logger.Named("UpdateAnswer").
		With("user_id", UserID).With("osbb_id", OSBBID).With("poll_id", PollID).
		With("answer", answer)
	if err := answer.Validate(); err != nil {
		logger.Error("failed to validate update poll data", "error", err)
		return errs.Err(err).Code("failed to validate update poll data").Kind(errs.Validation)
	}
	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}

	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not update a test answer", "error", err)
		return errs.M("user not osbb head").Code("User can not update a test answer").Kind(errs.Private)
	}

	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID, WithTestAnswers: true})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return errs.M("failed to get poll").Code("Failed to get poll").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist", "error", err)
		return errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}
	err = s.businessStorage.OSBB.UpdateAnswer(0, PollID, answer)
	if err != nil {
		logger.Error("failed to update answer", "error", err)
		return errs.Err(err).Code("Failed to update answer").Kind(errs.Database)
	}

	return nil
}

func (s *osbbService) DeleteAnswer(UserID, OSBBID, PollID uint64) error {
	logger := s.logger.Named("DeleteAnswer").
		With("user_id", UserID).With("osbb_id", OSBBID).With("poll_id", PollID)

	user, err := s.businessStorage.User.GetUser(UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return errs.M("user not found").Code("user do not exist").Kind(errs.Database)
	}

	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not delete an test answer", "error", err)
		return errs.M("user not osbb head").Code("User can not delete a test answer").Kind(errs.Private)
	}

	poll, err := s.businessStorage.OSBB.GetPoll(PollID, PollFilter{OSBBID: &OSBBID, WithTestAnswers: true})
	if err != nil {
		logger.Error("failed to get poll", "error", err)
		return errs.M("failed to get poll").Code("Failed to get poll").Kind(errs.Database)
	}
	if poll == nil {
		logger.Error("poll do not exist", "error", err)
		return errs.M("poll not found").Code("Poll do not exist").Kind(errs.Database)
	}
	err = s.businessStorage.OSBB.DeleteAnswer(0, AnswerFilter{PollID: &PollID})
	if err != nil {
		logger.Error("failed to delete answer", "error", err)
		return errs.Err(err).Code("Failed to delete answer").Kind(errs.Database)
	}

	return nil
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
	pollResult, err := s.businessStorage.OSBB.GetPollResult(PollID, PollFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get poll results", "error", err)
		return nil, errs.Err(err).Code("Failed to get poll result").Kind(errs.Database)
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
	if user.Role != entity.UserRoleOSBBHead {
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
	if user.Role != entity.UserRoleOSBBHead {
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

func (s *osbbService) ListInhabitants(UserID, OSBBID uint64, filter UserFilter) ([]entity.User, error) {
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
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not create a poll answer", "error", err)
		return nil, errs.M("user not osbb head").Code("User can not create a poll answer").Kind(errs.Private)
	}
	osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBFilter{OSBBID: &OSBBID})
	if osbb == nil {
		logger.Error("osbb do not exist", "error", err)
		return nil, errs.M("osbb not found").Code("Osbb do not exist").Kind(errs.Database)
	}
	inhabitants, err := s.businessStorage.User.ListUsers(filter)
	if err != nil {
		logger.Error("failed to get list users", "error", err)
		return nil, errs.Err(err).Code("Failed to get list users").Kind(errs.Database)
	}
	return inhabitants, nil
}

func (s *osbbService) UpdateInhabitant(UserID, OSBBID uint64, inhabitant entity.EventUserUpdatePayload) error {
	logger := s.logger.Named("UpdateInhabitant").
		With("user_id", UserID).With("osbb_id", OSBBID).With("inhabitant", inhabitant)
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

func (s *osbbService) ApproveInhabitant(UserID, OSBBID uint64, inhabitant entity.EventApproveUser) error {
	logger := s.logger.Named("ApproveInhabitant").
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
	if user.Role != entity.UserRoleOSBBHead {
		logger.Error("User can not delete an test answer", "error", err)
		return errs.M("user not osbb head").Code("User can not delete a test answer").Kind(errs.Private)
	}
	approvedUser, err := s.businessStorage.User.GetUser(inhabitant.UserID, UserFilter{OSBBID: &OSBBID})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if approvedUser == nil {
		logger.Error("approved user do not exist", "error", err)
		return errs.M("approved user not found").Code("approved user do not exist").Kind(errs.Database)
	}

	err = s.businessStorage.User.ApproveUser(inhabitant.UserID, OSBBID, UserFilter{
		OSBBID:     &OSBBID,
		IsApproved: inhabitant.Answer,
	})
	if err != nil {
		logger.Error("failed to approve inhabitant", "error", err)
		return errs.Err(err).Code("Failed to approve inhabitant").Kind(errs.Database)
	}

	return nil
}
