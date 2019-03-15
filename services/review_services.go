package services

import (
	"class-review-backend/repositories"

	"class-review-backend/models"
)

type IReviewService interface {
	GetReview(uint64) (*models.Review, error)
}

// Implements IReviewService
type ReviewService struct {
	Repositories *repositories.Repositories
}

func DefaultReviewServices(repos *repositories.Repositories) *ReviewService {
	return &ReviewService{
		Repositories: repos,
	}
}

func (rs *ReviewService) GetReview(id uint64) (*models.Review, error) {
	return rs.Repositories.ReviewRepository.GetReview(id)
}
