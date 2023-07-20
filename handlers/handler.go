package handlers

import "kredit-plus/usecases"

func InitVersionOneKonsumenHandler(u usecases.Repositories) KonsumenHandlerInterface {
	uc := usecases.GetUseCase(u)
	return InitKonsumenHandler(uc.KonsumenUsecase)
}

func InitVersionOneAuthHandler(u usecases.Repositories) AuthenticationHandlerInterface {
	uc := usecases.GetUseCase(u)
	return InitAuthenticationHandler(uc.AuthUsecase)
}
