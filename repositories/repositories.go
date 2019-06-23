package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ReviewRepository    IReviewRepository
	CourseRepository    ICourseRepository
	ProfessorRepository IProfessorRepository
	Database            *sqlx.DB
}

func DefaultRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		ReviewRepository:    DefaultReviewRepository(db),
		CourseRepository:    DefaultCourseRepository(),
		ProfessorRepository: DefaultProfessorRepository(db),
		Database:            db,
	}
}
