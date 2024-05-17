package libraries

import (
	"crud_operation/configs"
	"crud_operation/repository"
)

type AuthenticationService interface {
	AuthenticateUser(username, password string) (bool, error)
	Login(username, password string) (bool, error)
	ResetPassword(username, password string) (bool, error)
}

type authenticationService struct {
	config   configs.ApiConfig
	authRepo repository.AuthenticationRepository
}

func NewAuthenticationService(authRepo repository.AuthenticationRepository) AuthenticationService {
	return &authenticationService{
		config:   configs.GetConfig(),
		authRepo: authRepo,
	}
}

func (auth *authenticationService) AuthenticateUser(username, password string) (bool, error) {
	return auth.authRepo.AuthenticateUser(username, password)
}

func (auth *authenticationService) Login(username, password string) (bool, error) {
	return auth.authRepo.Login(username, password)
}

func (auth *authenticationService) ResetPassword(username, password string) (bool, error) {
	return auth.authRepo.ResetPassword(username, password)
}
