package services

import (
	"errors"
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint64 `json: "User_id"`
}

type tokenService struct {
	config          *config.Config
	businessStorage Storages
}

func NewTokenService(opts *Options) *tokenService {
	return &tokenService{
		config:          opts.Config,
		businessStorage: opts.Storages,
	}
}

func (s *tokenService) GenerateToken(UserID uint64) (entity.TokenValue, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.config.TokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: UserID,
	})
	signedToken, err := token.SignedString([]byte(s.config.SignInKey))
	if err != nil {
		return "", err
	}
	newToken := &entity.Token{Value: entity.TokenValue(signedToken),
		UserID: UserID, Revoked: false}
	err = s.businessStorage.Token.CreateToken(newToken)
	if err != nil {
		return "", err
	}
	return newToken.Value, nil
}

func (s *tokenService) GetByToken(token string) (*entity.Token, error) {
	return s.businessStorage.Token.GetByToken(token)
}

func (s *tokenService) ParseToken(accessToken string) (uint64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.config.SignInKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserID, nil
}

func (s *tokenService) RefreshToken(UserID uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.config.TokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: UserID,
	})
	tokenSigned, err := token.SignedString([]byte(s.config.SignInKey))
	err = s.businessStorage.Token.UpdateByUser(&entity.Token{UserID: UserID})
	if err != nil {
		return "", err
	}
	return tokenSigned, err
}