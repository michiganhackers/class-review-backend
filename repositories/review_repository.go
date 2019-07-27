package repositories

import (
	"class-review-backend/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type IReviewRepository interface {
	GetReviewById(uint64) (*models.Review, error)
}

// Implements IReviewRepository
type ReviewRepository struct {
	Database *sqlx.DB
}

func DefaultReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{
		Database: db,
	}
}

func (rr *ReviewRepository) GetReviewById(id uint64) (*models.Review, error) {
	var review models.Review
	err := rr.Database.Get(&review, `SELECT id,
											rating, 
											difficulty, 
											interest, 
											courseId, 
											review_date, 
											is_anonymous, 
											review_text, 
											professor_uniqname, 
											helpfulCount, 
											notHelpfulCount, 
											semester, 
											userEmail					
	       							 FROM reviews WHERE id=?`, id)
	if err != nil {
		log.Println("Error in GetReviewById:", err)
		return &review, err
	}

	return &review, nil
}
