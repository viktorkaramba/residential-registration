package services

import (
	"errors"
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/errs"
	"residential-registration/backend/pkg/logging"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint64 `json: "User_id"`
}

type tokenService struct {
	logger          logging.Logger
	config          *config.Config
	businessStorage Storages
}

func NewTokenService(opts *Options) *tokenService {
	return &tokenService{
		logger:          opts.Logger.Named("TokenService"),
		config:          opts.Config,
		businessStorage: opts.Storages,
	}
}

func (s *tokenService) GenerateToken(UserID uint64) (entity.TokenValue, error) {
	logger := s.logger.Named("GenerateToken").
		With("user_id", UserID)
	err := s.businessStorage.Token.UpdateByUser(&entity.Token{
		Revoked: true,
		UserID:  UserID,
	})
	if err != nil {
		logger.Error("failed to revoke all users tokens", "error", err)
		return "", errs.Err(err).Code("Failed to revoke all users tokens").Kind(errs.Database)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.config.TokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: UserID,
	})
	signedToken, err := token.SignedString([]byte(s.config.SignInKey))
	if err != nil {
		logger.Error("failed to sign token", "error", err)
		return "", errs.Err(err).Code("Failed to sign token").Kind(errs.Private)
	}
	newToken := &entity.Token{Value: entity.TokenValue(signedToken),
		UserID: UserID, Revoked: false}
	err = s.businessStorage.Token.CreateToken(newToken)
	if err != nil {
		logger.Error("failed to create token", "error", err)
		return "", errs.Err(err).Code("Failed to create token").Kind(errs.Database)
	}

	return newToken.Value, nil
}

func (s *tokenService) GetByToken(token string) (*entity.Token, error) {
	logger := s.logger.Named("GetByToken").
		With("token", token)
	existToken, err := s.businessStorage.Token.GetByToken(token)
	if err != nil {
		logger.Error("failed to get token", "error", err)
		return nil, errs.Err(err).Code("Failed to get token").Kind(errs.Database)
	}
	if existToken == nil {
		logger.Error("token does not exist")
		return nil, errs.M("token not found").Code("Token does not exist").Kind(errs.NotExist)
	}
	return existToken, nil
}

func (s *tokenService) ParseToken(accessToken string) (uint64, error) {
	logger := s.logger.Named("ParseToken").
		With("accessToken", accessToken)
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.config.SignInKey), nil
	})

	if err != nil {
		logger.Error("failed to parse token with claims", "error", err)
		return 0, errs.Err(err).Code("Failed to parse token with claims").Kind(errs.Private)
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		logger.Error("token claims are not of type *tokenClaims", "error", err)
		return 0, errs.Err(err).Code("Token claims are not of type *tokenClaims").Kind(errs.Private)
	}
	return claims.UserID, nil
}

func (s *tokenService) RefreshToken(UserID uint64) (string, error) {
	logger := s.logger.Named("RefreshToken").
		With("user_id", UserID)

	err := s.businessStorage.Token.UpdateByUser(&entity.Token{UserID: UserID, Revoked: true})
	if err != nil {
		logger.Error("failed to update token by user", "error", err)
		return "", errs.Err(err).Code("Failed to update token by user").Kind(errs.Database)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.config.TokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: UserID,
	})
	tokenSigned, err := token.SignedString([]byte(s.config.SignInKey))
	if err != nil {
		logger.Error("failed to sign token", "error", err)
		return "", errs.Err(err).Code("Failed to sign token").Kind(errs.Private)
	}

	return tokenSigned, err
}
