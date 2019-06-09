package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ReviewRepository IReviewRepository
  CourseRepository ICourseRepository
	Database         *sqlx.DB
}

func DefaultRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		ReviewRepository: DefaultReviewRepository(db),
    CourseRepository: DefaultCourseRepository(),
		Database:         db,
	}
}
