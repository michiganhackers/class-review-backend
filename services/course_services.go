package services

import (
	"class-review-backend/repositories"

	"class-review-backend/models"
)

type ICourseService interface {
	GetCourse(uint64) (*models.Course, error)
}

// Implements ICourseService
type CourseService struct {
	Repositories *repositories.Repositories
}

func DefaultCourseServices(repos *repositories.Repositories) *CourseService {
	return &CourseService{
		Repositories: repos,
	}
}

func (cs *CourseService) GetCourse(id uint64) (*models.Course, error) {
	return cs.Repositories.CourseRepository.GetCourse(id)
}
