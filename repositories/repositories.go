package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ReviewRepository IReviewRepository
	Database         *sqlx.DB
}

func DefaultRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		ReviewRepository: DefaultReviewRepository(db),
		Database:         db,
	}
}
