package repositories

import (
    "class-review-backend/models"
)

type IReviewRepository interface {
    GetReview(uint64) (*models.Review, error)
}

// Implements IReviewRepository
type ReviewRepository struct {
    // TODO: put database engine here.
}

func DefaultReviewRepository() *ReviewRepository {
    // TODO: set up database engine here
    return &ReviewRepository{}
}

func (rr *ReviewRepository) GetReview(id uint64) (*models.Review, error) {

    var review models.Review
    return &review, nil

}
