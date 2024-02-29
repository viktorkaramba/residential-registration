package storages

import (
	"errors"
	"residential-registration/backend/internal/entity"

	"github.com/a631807682/zerofield"
	"gorm.io/gorm"
)

type tokenStorage struct {
	db *gorm.DB
}

func NewTokenStorage(db *gorm.DB) *tokenStorage {
	return &tokenStorage{
		db: db,
	}
}

func (s *tokenStorage) CreateToken(token *entity.Token) error {
	return s.db.Create(token).Error
}

func (s *tokenStorage) GetByToken(token string) (*entity.Token, error) {
	newToken := &entity.Token{}
	err := s.db.Model(&entity.Token{}).Where(entity.Token{Value: entity.TokenValue(token)}).First(newToken).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return newToken, err
}

func (s *tokenStorage) Update(token *entity.Token) error {
	return s.db.
		Where(&entity.Token{Value: token.Value}).
		Scopes(zerofield.UpdateScopes()).
		Updates(token).
		Error
}

func (s *tokenStorage) UpdateByUser(token *entity.Token) error {
	return s.db.
		Where(&entity.Token{UserID: token.UserID}).
		Updates(token).
		Error
}
