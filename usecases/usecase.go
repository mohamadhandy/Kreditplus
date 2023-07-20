package usecases

import (
	"kredit-plus/repositories"

	"gorm.io/gorm"
)

type Repositories struct {
	KonsumenRepository    repositories.KonsumenRepositoryInterface
	AuthRepository        repositories.AuthRepositoryInterface
	ProductRepository     repositories.ProductRepositoryInterface
	TransactionRepository repositories.TransaksiRepositoryInterface
}

type Usecases struct {
	KonsumenUsecase  KonsumenUsecaseInterface
	AuthUsecase      AuthUseCaseInterface
	ProductUsecase   ProductUsecaseInterface
	TransaksiUsecase TransaksiUsecaseInterface
}

var useCaseInstance Usecases

func InitRepository(db *gorm.DB) Repositories {
	return Repositories{
		KonsumenRepository:    repositories.InitKonsumenRepository(db),
		AuthRepository:        repositories.InitAuthRepository(db),
		ProductRepository:     repositories.InitProductRepository(db),
		TransactionRepository: repositories.InitTransaksiRepository(db),
	}
}

func GetUseCase(r Repositories) *Usecases {
	if (Usecases{}) == useCaseInstance {
		useCaseInstance = Usecases{
			KonsumenUsecase:  InitKonsumenUsecase(r.KonsumenRepository),
			AuthUsecase:      InitAuthenticationUseCase(r.AuthRepository),
			ProductUsecase:   InitProductUseCase(r.ProductRepository),
			TransaksiUsecase: InitTransaksiUsecase(r.TransactionRepository),
		}
	}
	return &useCaseInstance
}
