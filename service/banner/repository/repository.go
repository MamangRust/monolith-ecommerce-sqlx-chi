package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	BannerQuery   BannerQueryRepository
	BannerCommand BannerCommandRepository
}

func NewRepositories(DB *sqlx.DB) *Repositories {

	return &Repositories{
		BannerQuery:   NewBannerQueryRepository(DB),
		BannerCommand: NewBannerCommandRepository(DB),
	}
}
