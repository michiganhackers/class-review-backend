package repositories

import (
	"class-review-backend/models"

	"github.com/jmoiron/sqlx"
)

type IReviewRepository interface {
	GetReview(uint64) (*models.Review, error)
}

// Implements IReviewRepository
type ReviewRepository struct {
	Database *sqlx.DB
}

func DefaultReviewRepository(db *sqlx.DB) *ReviewRepository {
	// TODO: set up database engine here
	return &ReviewRepository{
		Database: db,
	}
}

func (rr *ReviewRepository) GetReview(id uint64) (*models.Review, error) {

	var review models.Review
	return &review, nil

}
