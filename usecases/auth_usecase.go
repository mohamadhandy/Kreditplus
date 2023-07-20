package usecases

import (
	"kredit-plus/helper"
	"kredit-plus/models"
	"kredit-plus/repositories"
)

type AuthUseCaseInterface interface {
	BeginSession(authRequest models.AuthRequest) helper.Response
}

type authenticationUseCase struct {
	r repositories.AuthRepositoryInterface
}

func InitAuthenticationUseCase(r repositories.AuthRepositoryInterface) AuthUseCaseInterface {
	return &authenticationUseCase{
		r,
	}
}

func (u *authenticationUseCase) BeginSession(req models.AuthRequest) helper.Response {
	return <-u.r.BeginSession(req)
}
