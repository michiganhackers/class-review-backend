package repositories

import (
	"class-review-backend/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type IReviewRepository interface {
	GetAllReviews() ([]models.Review, error)
	GetReviewById(uint64) (*models.Review, error)
	PostReview(*models.Review) error
	UpdateReview(*models.Review, uint64) (*models.Review, error)
	DeleteReview(uint64) error
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

func (rr *ReviewRepository) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := rr.Database.Select(&reviews, `SELECT id,
												rating, 
												difficulty, 
												interest, 
												courseId, 
												review_date, 
												is_anonymous, 
												review_text, 
												professor_uniqname, 
												semester, 
												userEmail					
										 FROM reviews`)
	if err != nil {
		log.Println("Error in GetAllReviews:", err)
		return reviews, err
	}

	return reviews, nil
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
											semester, 
											userEmail					
	       							 FROM reviews WHERE id=?`, id)
	if err != nil {
		log.Println("Error in GetReviewById:", err)
		return &review, err
	}

	return &review, nil
}

func (rr *ReviewRepository) PostReview(reviewInput *models.Review) error {
	return nil
}

func (rr *ReviewRepository) UpdateReview(reviewInput *models.Review, id uint64) (*models.Review, error) {
	return reviewInput, nil
}

func (rr *ReviewRepository) DeleteReview(id uint64) error {
	_, err := rr.Database.Exec(`DELETE FROM reviews WHERE id=?`, id)
	if err != nil {
		log.Println("Error in DeleteReview:", err)
		return err
	}

	return nil
}