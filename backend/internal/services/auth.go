package services

import (
	"crypto/sha1"
	"fmt"
	"residential-registration/backend/config"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/errs"
	"residential-registration/backend/pkg/logging"
)

type authService struct {
	logger          logging.Logger
	config          *config.Config
	businessStorage Storages
}

func NewAuthService(opts *Options) *authService {
	return &authService{
		logger:          opts.Logger.Named("AuthService"),
		config:          opts.Config,
		businessStorage: opts.Storages,
	}
}

func (s *authService) AddUser(OSBBID uint64, inputUser entity.EventUserPayload) (*entity.User, error) {
	logger := s.logger.Named("AddUser").
		With("osbb_id", OSBBID).With("input_user", inputUser)

	building, err := s.businessStorage.Building.GetByOSBBID(OSBBID)

	if err != nil {
		logger.Error("failed to get building", "error", err)
		return nil, errs.Err(err).Code("Failed to get building").Kind(errs.Database)
	}
	if building == nil {
		logger.Error("building do not exist", "error", err)
		return nil, errs.M("building not found").Code("Building do not exist").Kind(errs.Database)
	}

	User := &entity.User{
		Apartment: entity.Apartment{
			BuildingID: building.ID,
			Number:     inputUser.ApartmentNumber,
			Area:       inputUser.ApartmentArea,
		},
		OSBBID: OSBBID,
		FullName: entity.FullName{
			FirstName:  inputUser.FirstName,
			Surname:    inputUser.Surname,
			Patronymic: inputUser.Patronymic,
		},
		Password:    s.generatePasswordHash(string(inputUser.Password)),
		PhoneNumber: inputUser.PhoneNumber,
		Role:        entity.UserRoleInhabitant,
	}
	err = s.businessStorage.User.CreateUser(User)
	if err != nil {
		logger.Error("failed to сreate user", "error", err)
		return nil, errs.Err(err).Code("Failed to сreate user").Kind(errs.Database)
	}

	return User, nil
}

func (s *authService) Login(inputLogin entity.EventLoginPayload) (*entity.User, error) {
	logger := s.logger.Named("Login").
		With("input_login", inputLogin)

	user, err := s.businessStorage.User.GetUser(0, UserFilter{
		PhoneNumber: &inputLogin.PhoneNumber,
	})
	if err != nil {
		logger.Error("failed to get user", "error", err)
		return nil, errs.Err(err).Code("Failed to get user").Kind(errs.Database)
	}
	if user == nil {
		logger.Error("user do not exist", "error", err)
		return nil, errs.M("user not found").Code("User do not exist").Kind(errs.NotExist)
	}

	if user.IsApproved == nil {
		logger.Error("user wait approve", "error", err)
		return nil, errs.M("user not approve").Code("Failed to login").Kind(errs.Private)
	}

	if !*user.IsApproved {
		logger.Error("user not approved", "error", err)
		return nil, errs.M("user not approved").Code("Failed to login").Kind(errs.Private)
	}
  
	if user.Password != GeneratePasswordHash(s.config.Salt, string(inputLogin.Password)) {
		logger.Error("incorrect password", "error", err)
		return nil, errs.M("incorrect password").Code("Failed to login").Kind(errs.Private)
	}

	return user, nil
}

func (s *authService) Logout(token entity.TokenValue) error {
	logger := s.logger.Named("Logout").
		With("token", token)

	existToken, err := s.businessStorage.Token.GetByToken(string(token))
	if err != nil {
		logger.Error("failed to get token", "error", err)
		return errs.Err(err).Code("Failed to get token").Kind(errs.Database)
	}

	if existToken == nil {
		logger.Error("token does not exist")
		return errs.M("token not found").Code("Token does not exist").Kind(errs.NotExist)
	}

	err = s.businessStorage.Token.Update(&entity.Token{
		Revoked: true,
		Value:   token,
	})
	if err != nil {
		logger.Error("failed to revoke all users tokens", "error", err)
		return errs.Err(err).Code("Failed to revoke all users tokens").Kind(errs.Database)
	}
	return nil
}

func (s *authService) generatePasswordHash(password string) entity.Password {
	hash := sha1.New()
	hash.Write([]byte(password))
	return entity.Password(fmt.Sprintf("%x", hash.Sum([]byte(s.config.Salt))))
}
