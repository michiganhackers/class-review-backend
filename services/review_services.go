package services

import (
    "class-review-backend/repositories"

    "class-review-backend/models"
)

type IReviewService interface {
    GetAllReviews() ([]models.Review, error)
    GetReviewById(uint64) (*models.Review, error)
    PostReview(*models.Review) error
    UpdateReview(*models.Review, uint64) (*models.Review, error)
    DeleteReview(uint64) error
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

func (rs *ReviewService) GetAllReviews() ([]models.Review, error) {
    return rs.Repositories.ReviewRepository.GetAllReviews()
}

func (rs *ReviewService) GetReviewById(id uint64) (*models.Review, error) {
    return rs.Repositories.ReviewRepository.GetReviewById(id)
}

func (rs *ReviewService) PostReview(reviewInput *models.Review) error {
    return rs.Repositories.ReviewRepository.PostReview(reviewInput)
}

func (rs *ReviewService) UpdateReview(reviewInput *models.Review, id uint64) (*models.Review, error) {
    return rs.Repositories.ReviewRepository.UpdateReview(reviewInput, id)
}

func (rs *ReviewService) DeleteReview(id uint64) error {
    return rs.Repositories.ReviewRepository.DeleteReview(id)
}