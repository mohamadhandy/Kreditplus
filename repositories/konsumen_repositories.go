package repositories

import "gorm.io/gorm"

type KonsumenRepositoryInterface interface {
}

type konsumenRepository struct {
	db *gorm.DB
}

func InitKonsumenRepository(db *gorm.DB) KonsumenRepositoryInterface {
	return &konsumenRepository{
		db,
	}
}
