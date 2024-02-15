package services

import (
	"errors"
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"time"
)

type osbbService struct {
	config          *config.Config
	businessStorage Storages
}

func NewOSBBService(opts *Options) *osbbService {
	return &osbbService{
		config:          opts.Config,
		businessStorage: opts.Storages,
	}
}

func (s *osbbService) AddOSBB(inputOSBB entity.InputOSBB) (*entity.OSBB, error) {
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
		return nil, err
	}

	return osbb, nil
}

func (s *osbbService) AddAnnouncement(UserID, OSBBID uint64, inputAnnouncement entity.InputAnnouncement) (*entity.Announcement, error) {
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		return nil, err
	}
	if user.Role == entity.UserRoleOSBBHEad {
		osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBID)
		if osbb == nil {
			return nil, errors.New("there are no osbb with this id")
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
			return nil, err
		}
		return announcement, nil
	}
	return nil, errors.New("user not osbb head")
}

func (s *osbbService) AddPoll(UserID, OSBBID uint64, inputPoll entity.InputPoll) (*entity.Poll, error) {
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		return nil, err
	}
	if user.Role == entity.UserRoleOSBBHEad {
		osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBID)
		if osbb == nil {
			return nil, errors.New("there are no osbb with this id")
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
			return nil, err
		}
		return poll, nil
	}
	return nil, errors.New("user not osbb head")
}

func (s *osbbService) AddPollTest(UserID, OSBBID uint64, inputPollTest entity.InputPollTest) (*entity.Poll, error) {
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		return nil, err
	}
	if user.Role == entity.UserRoleOSBBHEad {
		osbb, err := s.businessStorage.OSBB.GetOSBB(OSBBID)
		if osbb == nil {
			return nil, errors.New("there are no osbb with this id")
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
			return nil, err
		}
		return poll, nil
	}
	return nil, errors.New("user not osbb head")
}

func (s *osbbService) AddPollAnswer(UserID, PollID uint64, inputPollAnswer entity.InputPollAnswer) (*entity.Answer, error) {
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		return nil, err
	}
	if user.Role == entity.UserRoleOSBBHEad {
		poll, err := s.businessStorage.OSBB.GetPoll(PollID)
		if err != nil {
			return nil, err
		}
		if poll == nil {
			return nil, errors.New("there are no poll with this id")
		}
		answer := &entity.Answer{
			PollID:  PollID,
			UserID:  UserID,
			Content: inputPollAnswer.Content,
		}
		err = s.businessStorage.OSBB.CreatAnswer(answer)
		if err != nil {
			return nil, err
		}
		return answer, nil
	}
	return nil, errors.New("user not osbb head")
}

func (s *osbbService) AddPollAnswerTest(UserID, PollID uint64, inputPollAnswerTest entity.InputPollAnswerTest) (*entity.Answer, error) {
	user, err := s.businessStorage.User.GetUser(UserID)
	if err != nil {
		return nil, err
	}
	if user.Role == entity.UserRoleOSBBHEad {
		poll, err := s.businessStorage.OSBB.GetPoll(PollID)
		if err != nil {
			return nil, err
		}
		if poll == nil {
			return nil, errors.New("there are no poll with this id")
		}
		var isExist bool
		for _, answer := range poll.TestAnswer {
			if answer.PollID == PollID && answer.ID == inputPollAnswerTest.TestAnswerID {
				isExist = true
			}
		}
		if !isExist {
			return nil, errors.New("there are no test answer in this poll")
		}
		answer := &entity.Answer{
			PollID:       PollID,
			UserID:       UserID,
			TestAnswerID: inputPollAnswerTest.TestAnswerID,
		}
		err = s.businessStorage.OSBB.CreatAnswer(answer)
		if err != nil {
			return nil, err
		}
		return answer, nil
	}
	return nil, errors.New("user not osbb head")
}
