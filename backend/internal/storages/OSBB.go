package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/internal/services"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type OSBBStorage struct {
	db *gorm.DB
}

func NewOSBBStorage(db *gorm.DB) *OSBBStorage {
	return &OSBBStorage{
		db: db,
	}
}

func (s *OSBBStorage) CreateOSBB(OSBB *entity.OSBB) error {
	err := s.db.Create(OSBB).Error
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		if pgErr.Code == "23505" {
			if strings.Contains(err.Error(), "\"idx_users_phone_number\"") {
				return services.ErrPhoneNumberDuplicate
			} else if strings.Contains(err.Error(), "\"idx_name\"") {
				return services.ErrEDRPOUDuplicate
			}
		}
	}
	return err
}

func (s *OSBBStorage) ListOSBBS(filter services.OSBBFilter) ([]entity.OSBB, error) {
	stmt := s.db.
		Model(&entity.OSBB{})

	if filter.WithOSBBHead {
		stmt.Preload("OSBBHead")
	}
	if filter.WithBuilding {
		stmt.Preload("Building")
	}
	if filter.WithAnnouncements {
		stmt.Preload("Announcement").Order("created_at DESC")
	}

	var osbbs []entity.OSBB
	return osbbs, stmt.Order("created_at DESC").Find(&osbbs).Error
}

func (s *OSBBStorage) GetOSBB(filter services.OSBBFilter) (*entity.OSBB, error) {
	osbb := &entity.OSBB{}
	stmt := s.db.
		Model(&entity.OSBB{})
	if filter.OSBBID != nil {
		stmt = stmt.Where("id = ?", *filter.OSBBID)
	}
	if filter.WithOSBBHead {
		stmt.Preload("OSBBHead")
	}
	if filter.WithBuilding {
		stmt.Preload("Building")
	}
	if filter.WithAnnouncements {
		stmt.Preload("Announcement").Order("created_at DESC")
	}
	err := stmt.First(osbb).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return osbb, err
}

func (s *OSBBStorage) UpdateOSBB(OSBBID uint64, opts *entity.EventOSBBUpdatePayload) error {
	stmt := s.db.Model(&entity.OSBB{})
	var osbb entity.OSBB

	if OSBBID != 0 {
		stmt = stmt.Where("id = ?", OSBBID)
	}

	if opts.Name != nil {
		osbb.Name = *opts.Name
	}
	if opts.EDRPOU != nil {
		osbb.EDRPOU = *opts.EDRPOU
	}
	if opts.Rent != nil {
		osbb.Rent = *opts.Rent
	}
	if opts.Photo != nil {
		osbb.Photo = opts.Photo
	}
	err := stmt.Updates(osbb).Error
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		if pgErr.Code == "23505" {
			if strings.Contains(err.Error(), "\"idx_name\"") {
				return services.ErrEDRPOUDuplicate
			}
		}
	}
	return nil
}

func (s *OSBBStorage) CreateAnnouncement(announcement *entity.Announcement) error {
	return s.db.Create(announcement).Error
}

func (s *OSBBStorage) GetAnnouncement(AnnouncementID uint64, filter services.AnnouncementFilter) (*entity.Announcement, error) {
	announcement := &entity.Announcement{}
	stmt := s.db.
		Model(&entity.Announcement{})
	if AnnouncementID != 0 {
		stmt = stmt.Where("id = ?", AnnouncementID)
	}
	if filter.OSBBID != nil {
		stmt = stmt.Where("osbb_id = ?", *filter.OSBBID)
	}

	err := stmt.First(announcement).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return announcement, err
}

func (s *OSBBStorage) ListAnnouncements(filter services.AnnouncementFilter) ([]entity.Announcement, error) {
	stmt := s.db.
		Model(&entity.Announcement{})

	if filter.OSBBID != nil {
		stmt = stmt.Where("osbb_id = ?", *filter.OSBBID)
	}
	var announcements []entity.Announcement
	return announcements, stmt.Order("created_at DESC").Find(&announcements).Error
}

func (s *OSBBStorage) UpdateAnnouncement(AnnouncementID uint64, opts *entity.EventAnnouncementUpdatePayload) error {
	stmt := s.db.Model(&entity.Announcement{})
	var announcement entity.Announcement

	if AnnouncementID != 0 {
		stmt = stmt.Where("id = ?", AnnouncementID)
	}

	if opts.Title != nil {
		announcement.Title = *opts.Title
	}
	if opts.Content != nil {
		announcement.Content = *opts.Content
	}

	return stmt.Updates(announcement).Error
}

func (s *OSBBStorage) DeleteAnnouncement(AnnouncementID uint64, filter services.AnnouncementFilter) error {
	stmt := s.db.Model(&entity.Announcement{})
	var announcement *entity.Announcement

	if AnnouncementID != 0 {
		stmt = stmt.Where("id = ?", AnnouncementID)
	}

	if filter.OSBBID != nil {
		stmt = stmt.Where("osbb_id = ?", *filter.OSBBID)
	}

	return stmt.Delete(&announcement).Error
}

func (s *OSBBStorage) CreatePoll(poll *entity.Poll) error {
	return s.db.Create(poll).Error
}

func (s *OSBBStorage) ListPolls(filter services.PollFilter) ([]entity.Poll, error) {
	stmt := s.db.
		Model(&entity.Poll{})

	if filter.OSBBID != nil {
		stmt = stmt.Where("osbb_id = ?", *filter.OSBBID)
	}
	if filter.WithTestAnswers {
		stmt = stmt.Preload("TestAnswers")
	}
	var polls []entity.Poll
	return polls, stmt.Order("created_at DESC").Find(&polls).Error
}

func (s *OSBBStorage) UpdatePoll(PollID uint64, opts *entity.EventPollUpdatePayload) error {
	stmt := s.db.Model(&entity.Poll{})
	var poll entity.Poll

	if PollID != 0 {
		stmt = stmt.Where("id = ?", PollID)
	}

	if opts.Question != nil {
		poll.Question = *opts.Question
	}

	if opts.IsClosed != nil {
		poll.IsClosed = *opts.IsClosed
	}

	if opts.FinishedAt != nil {
		poll.FinishedAt = *opts.FinishedAt
	}

	return stmt.Updates(poll).Error
}

func (s *OSBBStorage) DeletePoll(PollID uint64, filter services.PollFilter) error {
	stmt := s.db.Model(&entity.Poll{})
	var poll *entity.Poll

	if PollID != 0 {
		stmt = stmt.Where("id = ?", PollID)
	}

	if filter.OSBBID != nil {
		stmt = stmt.Where("osbb_id = ?", *filter.OSBBID)
	}

	return stmt.Delete(&poll).Error
}

func (s *OSBBStorage) UpdateTestAnswer(TestAnswerID uint64, opts *entity.EventTestAnswerUpdatePayload) error {
	stmt := s.db.Model(&entity.TestAnswer{})
	var testAnswer entity.TestAnswer

	if TestAnswerID != 0 {
		stmt = stmt.Where("id = ?", TestAnswerID)
	}

	if opts.Content != nil {
		testAnswer.Content = *opts.Content
	}

	return stmt.Updates(testAnswer).Error
}

func (s *OSBBStorage) DeleteTestAnswer(TestAnswerID uint64, filter services.TestAnswerFilter) error {
	stmt := s.db.Model(&entity.TestAnswer{})
	var testAnswer *entity.TestAnswer

	if TestAnswerID != 0 {
		stmt = stmt.Where("id = ?", TestAnswerID)
	}

	if filter.PollID != nil {
		stmt = stmt.Where("poll_id = ?", *filter.PollID)
	}

	return stmt.Delete(&testAnswer).Error
}

func (s *OSBBStorage) GetPoll(PollID uint64, filter services.PollFilter) (*entity.Poll, error) {
	poll := &entity.Poll{}
	stmt := s.db.Model(&entity.Poll{})
	if filter.OSBBID != nil {
		stmt = stmt.Where("osbb_id = ?", *filter.OSBBID)
	}
	if filter.WithTestAnswers {
		stmt = stmt.Preload("TestAnswers")
	}
	err := stmt.Where("id = ?", PollID).First(poll).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return poll, err
}

func (s *OSBBStorage) GetPollResult(PollID uint64, filter services.PollFilter) (*entity.PollResult, error) {
	pollResult := &entity.PollResult{}

	stmt := s.db.
		Model(&entity.Poll{})

	poll := &entity.Poll{}

	stmt = stmt.Preload("UserAnswers")
	err := stmt.Where("id = ?", PollID).First(poll).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	pollResult.Answer = poll.UserAnswers
	pollResult.CountOfAllAnswers = uint64(len(poll.UserAnswers))
	if poll.Type == entity.PollTypeTest {
		var testAnswerCount []entity.TestAnswerCount
		err = s.db.Raw(`
			SELECT ts.id, 
			       COUNT(a.test_answer_id)
			FROM test_answers ts
			    LEFT JOIN answers a
			        ON ts.id = a.test_answer_id
			WHERE ts.poll_id = ?
			GROUP BY ts.id;
	`, PollID).Scan(&testAnswerCount).Error
		pollResult.CountOfTestAnswers = testAnswerCount
	}
	return pollResult, nil
}

func (s *OSBBStorage) CreatAnswer(answer *entity.Answer) error {
	return s.db.Create(answer).Error
}

func (s *OSBBStorage) ListAnswers(filter services.AnswerFilter) ([]entity.Answer, error) {
	stmt := s.db.
		Model(&entity.Answer{})

	if filter.TestAnswerID != nil {
		stmt = stmt.Where("test_answer_id = ?", *filter.TestAnswerID)
	}
	if filter.PollID != nil {
		stmt = stmt.Where("poll_id = ?", *filter.PollID)
	}
	if filter.UserID != nil {
		stmt = stmt.Where("user_id = ?", *filter.UserID)
	}
	if filter.Content != nil {
		stmt = stmt.Where("content = ?", *filter.Content)
	}
	if filter.CreatedAt != nil {
		stmt = stmt.Where("created_at = ?", *filter.CreatedAt)
	}
	if filter.UpdateAt != nil {
		stmt = stmt.Where("update_at = ?", *filter.UpdateAt)
	}

	var answers []entity.Answer
	return answers, stmt.Order("created_at DESC").Find(&answers).Error
}

func (s *OSBBStorage) GetAnswer(AnswerID uint64, filter services.AnswerFilter) (*entity.Answer, error) {
	answer := &entity.Answer{}
	stmt := s.db.Model(&entity.Answer{})
	if AnswerID != 0 {
		stmt = stmt.Where(entity.Answer{ID: AnswerID})
	}
	if filter.TestAnswerID != nil {
		stmt = stmt.Where("test_answer_id = ?", *filter.TestAnswerID)
	}
	if filter.PollID != nil {
		stmt = stmt.Where("poll_id = ?", *filter.PollID)
	}
	if filter.UserID != nil {
		stmt = stmt.Where("user_id = ?", *filter.UserID)
	}
	if filter.Content != nil {
		stmt = stmt.Where("content = ?", *filter.Content)
	}
	if filter.CreatedAt != nil {
		stmt = stmt.Where("created_at = ?", *filter.CreatedAt)
	}
	if filter.UpdateAt != nil {
		stmt = stmt.Where("update_at = ?", *filter.UpdateAt)
	}
	err := stmt.First(answer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return answer, err
}

func (s *OSBBStorage) UpdateAnswer(AnswerID, PollID uint64, opts *entity.EventUserAnswerUpdatePayload) error {
	stmt := s.db.Model(&entity.Answer{})
	var answer entity.Answer

	if AnswerID != 0 {
		stmt = stmt.Where("id = ?", AnswerID)
	}

	if PollID != 0 {
		stmt = stmt.Where("poll_id = ?", PollID)
	}
	if opts.TestAnswerID != nil {
		answer.TestAnswerID = *opts.TestAnswerID
	}

	if opts.Content != nil {
		answer.Content = *opts.Content
	}

	return stmt.Updates(answer).Error
}

func (s *OSBBStorage) DeleteAnswer(AnswerID uint64, filter services.AnswerFilter) error {
	stmt := s.db.Model(&entity.Answer{})
	var answer *entity.Answer

	if AnswerID != 0 {
		stmt = stmt.Where("id = ?", AnswerID)
	}

	if filter.PollID != nil {
		stmt = stmt.Where("poll_id = ?", *filter.PollID)
	}

	return stmt.Delete(&answer).Error
}

func (s *OSBBStorage) CreatePayment(payment *entity.Payment) error {
	return s.db.Create(payment).Error
}

func (s *OSBBStorage) CreateUserPayment(userPayment *entity.Purchase) error {
	return s.db.Create(userPayment).Error
}
