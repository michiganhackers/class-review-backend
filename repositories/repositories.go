package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ReviewRepository IReviewRepository
	Database         *sqlx.DB
}

func DefaultRepositories() *Repositories {
	return &Repositories{
		ReviewRepository: DefaultReviewRepository(),
		Database:         CreateDB(),
	}
}
