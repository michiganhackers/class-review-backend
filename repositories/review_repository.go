package repositories

import (
	"class-review-backend/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type IReviewRepository interface {
	GetReview(uint64) (*models.Review, error)
	TestDB()
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

func (rr *ReviewRepository) TestDB() {
	rows, err := rr.Database.Query("SELECT id, name FROM dev_class_review.test_table")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		// Do something
	}
}
