package handlers

import "kredit-plus/usecases"

func InitVersionOneKonsumenHandler(u usecases.Repositories) KonsumenHandlerInterface {
	uc := usecases.GetUseCase(u)
	return InitKonsumenHandler(uc.KonsumenUsecase)
}
