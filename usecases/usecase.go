package usecases

import (
	"kredit-plus/repositories"

	"gorm.io/gorm"
)

type Repositories struct {
	KonsumenRepository repositories.KonsumenRepositoryInterface
}

type Usecases struct {
	KonsumenUsecase KonsumenUsecaseInterface
}

var useCaseInstance Usecases

func InitRepository(db *gorm.DB) Repositories {
	return Repositories{
		KonsumenRepository: repositories.InitKonsumenRepository(db),
	}
}

func GetUseCase(r Repositories) *Usecases {
	if (Usecases{}) == useCaseInstance {
		useCaseInstance = Usecases{
			KonsumenUsecase: InitKonsumenUsecase(r.KonsumenRepository),
		}
	}
	return &useCaseInstance
}
