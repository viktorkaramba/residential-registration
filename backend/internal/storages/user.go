package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/internal/services"

	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *userStorage {
	return &userStorage{
		db: db,
	}
}

func (s *userStorage) CreateUser(User *entity.User) error {
	return s.db.Create(User).Error
}

func (s *userStorage) GetUser(UserID uint64, filter services.UserFilter) (*entity.User, error) {
	stmt := s.db.
		Model(&entity.User{})
	if UserID != 0 {
		stmt = stmt.Where(entity.User{ID: UserID})
	}
	if filter.OSBBID != nil {
		stmt = stmt.Where(entity.User{OSBBID: *filter.OSBBID})
	}
	if filter.PhoneNumber != nil {
		stmt = stmt.Where(entity.User{PhoneNumber: *filter.PhoneNumber})
	}
	if filter.UserRole != nil {
		stmt = stmt.Where(entity.User{Role: *filter.UserRole})
	}
	var user *entity.User
	err := stmt.First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, err
}

func (s *userStorage) GetUserByPhoneNumber(phoneNumber entity.PhoneNumber) (*entity.User, error) {
	User := &entity.User{}
	err := s.db.Model(&entity.User{}).Where(entity.User{PhoneNumber: phoneNumber}).First(User).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return User, err
}

func (s *userStorage) ListUsers(filter services.UserFilter) ([]entity.User, error) {
	stmt := s.db.
		Model(&entity.User{})

	if filter.OSBBID != nil {
		stmt = stmt.Where(entity.User{OSBBID: *filter.OSBBID})
	}
	var users []entity.User
	return users, stmt.Find(&users).Error
}

func (s *userStorage) UpdateUser(UserID, OSBBID uint64, opts *entity.EventUserUpdatePayload) error {
	stmt := s.db.Model(&entity.User{})
	var user entity.User

	if UserID != 0 {
		user.ID = UserID
		stmt = stmt.Where(entity.User{ID: UserID})
	}
	if OSBBID != 0 {
		stmt = stmt.Where(entity.User{OSBBID: OSBBID})
	}

	if opts.FirstName != nil {
		user.FirstName = *opts.FirstName
	}
	if opts.Surname != nil {
		user.Surname = *opts.Surname
	}
	if opts.Patronymic != nil {
		user.Patronymic = *opts.Patronymic
	}
	if opts.ApartmentArea != nil {
		user.Apartment.Area = *opts.ApartmentArea
	}
	if opts.ApartmentNumber != nil {
		user.Apartment.Number = *opts.ApartmentNumber
	}
	if opts.PhoneNumber != nil {
		user.PhoneNumber = *opts.PhoneNumber
	}

	return stmt.Updates(user).Error
}
