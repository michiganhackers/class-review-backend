package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ReviewRepository IReviewRepository
	Database         *sqlx.DB
	// TODO: put database engine here
}

func DefaultRepositories() *Repositories {
	return &Repositories{
		ReviewRepository: DefaultReviewRepository(),
		Database:         CreateDB(),
	}
}
