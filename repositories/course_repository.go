package repositories

import (
    "class-review-backend/models"
)

type ICourseRepository interface {
    GetCourse(uint64) (*models.Course, error)
}

// Implements IReviewRepository
type CourseRepository struct {
    // TODO: put database engine here.
}

func DefaultCourseRepository() *CourseRepository {
    // TODO: set up database engine here
    return &CourseRepository{}
}

func (rr *CourseRepository) GetCourse(id uint64) (*models.Course, error) {

    var course models.Course
    return &course, nil

}
