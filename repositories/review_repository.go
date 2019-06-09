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
	_, err := rr.Database.Exec(`CREATE TABLE IF NOT EXISTS test(
		name varchar(255),
		value varchar(255)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8`)

	if err != nil {
		fmt.Println(err, "Could not create table")
	} else {
		fmt.Println("No error")
	}
}
