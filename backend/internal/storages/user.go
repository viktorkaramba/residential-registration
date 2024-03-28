package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/internal/services"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
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
	err := s.db.Create(User).Error
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		if pgErr.Code == "23505" {
			if strings.Contains(err.Error(), "\"idx_users_phone_number\"") {
				return services.ErrPhoneNumberDuplicate
			}
		}
	}
	return err
}

func (s *userStorage) GetUser(UserID uint64, filter services.UserFilter) (*entity.User, error) {
	stmt := s.db.
		Model(&entity.User{})
	if UserID != 0 {
		stmt = stmt.Where("id = ?", UserID)
	}
	if filter.OSBBID != nil {
		stmt = stmt.Where("osbb_id = ?", *filter.OSBBID)
	}
	if filter.PhoneNumber != nil {
		stmt = stmt.Where("phone_number = ?", *filter.PhoneNumber)
	}
	if filter.UserRole != nil {
		stmt = stmt.Where("role = ?", *filter.UserRole)
	}

	if filter.IsApproved != nil {
		stmt = stmt.Where("is_approved = ?", *filter.IsApproved)
	}

	stmt = stmt.Preload("Apartment")
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
	stmt = stmt.Preload("Apartment")
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
		stmt = stmt.Where("id = ?", UserID)
	}
	if OSBBID != 0 {
		stmt = stmt.Where("osbb_id = ?", OSBBID)
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
	if opts.Photo != nil {
		user.Photo = opts.Photo
	}
	err := stmt.Updates(user).Error
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		if pgErr.Code == "23505" {
			if strings.Contains(err.Error(), "\"idx_users_phone_number\"") {
				return services.ErrPhoneNumberDuplicate
			}
		}
	}
	return stmt.Updates(user).Error
}
