package repositories

import (
	"class-review-backend/models"
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/jmoiron/sqlx"
)

type IReviewRepository interface {
	GetAllReviews() ([]models.Review, error)
	GetReviewById(uint64) (*models.Review, error)
	PostReview(*models.Review) error
	UpdateReview(*models.Review, uint64) (*models.Review, error)
	DeleteReview(uint64) error
	GetRatingByReviewId(uint64) (*models.RatingCount, error)
	UpdateRating(*models.UserRating) (*models.UserRating, error)
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
	_, err := rr.Database.NamedExec(`INSERT INTO reviews 
									 VALUES (DEFAULT,
											 :rating, 
											 :difficulty, 
											 :interest, 
											 :courseId, 
											 :review_date,
											 :review_text, 
											 :professor_uniqname, 
											 :semester, 
											 :userEmail)`, reviewInput)
	if err != nil {
		log.Println("Error in PostReview:", err)
		return err
	}
	return nil
}

func (rr *ReviewRepository) UpdateReview(reviewInput *models.Review, id uint64) (*models.Review, error) {
	reviewInput.Id = id
	_, err := rr.Database.NamedExec(`UPDATE reviews 
									 SET rating = :rating, 
										 difficulty = :difficulty, 
										 interest = :interest, 
										 courseId = :courseId, 
										 review_date = :review_date, 
										 review_text = :review_text, 
										 professor_uniqname = :professor_uniqname, 
										 semester = :semester, 
										 userEmail = :userEmail 
									 WHERE id = :id`, reviewInput)
	if err != nil {
		log.Println("Error in UpdateReview:", err)
		return reviewInput, err
	}
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

func (rr *ReviewRepository) GetRatingByReviewId(reviewId uint64) (*models.RatingCount, error) {
	var rating models.RatingCount
	err := rr.Database.Get(&rating, `SELECT reviewId,
       								 		SUM(CASE WHEN helpful = 1 THEN 1 ELSE 0 END) AS helpfulCount,
       								 		SUM(CASE WHEN helpful = 0 THEN 1 ELSE 0 END) AS notHelpfulCount
									 FROM ratings WHERE reviewId = ?`, reviewId)
	if err != nil {
		log.Println("Error in GetRatingByReviewId:", err)
		return &rating, err
	}

	return &rating, nil
}

func (rr *ReviewRepository) UpdateRating(ratingInput *models.UserRating) (*models.UserRating, error) {
	// First, get all ratings matching a reviewId, because two equivalent hashed and salted emails may differ, so we
	// 	compare them with bcrypt's function instead of a WHERE clause in sql
	var ratings []models.UserRating
	err := rr.Database.Select(&ratings, `SELECT userEmail,
											    reviewId, 
											    helpful 
	       							 	 FROM ratings WHERE reviewId = ?`, ratingInput.ReviewId)
	if err != nil {
		log.Println("Error in UpdateRating:", err)
		return ratingInput, err
	}

	// For every rating matching the reviewId, if the email input matches a hashed email, we determine the effect based
	// 	on the input and the database value, and then return directly from inside the loop
	for _, element := range ratings {
		if bcrypt.CompareHashAndPassword([]byte(element.UserEmail), []byte(ratingInput.UserEmail)) == nil {
			// If the user changes their rating, update the database to reflect that
			if element.Helpful != ratingInput.Helpful {
				element.Helpful = ratingInput.Helpful
				_, err := rr.Database.NamedExec(`UPDATE ratings 
												 SET helpful = :helpful
												 WHERE userEmail = :userEmail AND reviewId = :reviewId`, element)
				if err != nil {
					log.Println("Error in UpdateRating:", err)
					return ratingInput, err
				}

				return ratingInput, nil
			} else {	// Otherwise, if the user rated the review the same way twice, delete the rating
				_, err := rr.Database.NamedExec(`DELETE FROM ratings 
												 WHERE userEmail = :userEmail AND reviewId = :reviewId`, element)
				if err != nil {
					log.Println("Error in UpdateRating:", err)
					return nil, err
				}

				return nil, nil
			}
		}
	}

	// At this point, we know a rating with this user and review combo does not exist, so we insert it into the table
	hash, err := bcrypt.GenerateFromPassword([]byte(ratingInput.UserEmail), 12)
	if err != nil {
		log.Println("Error in UpdateRating:", err)
		return ratingInput, err
	}
	// This new rating variable is used so we can return ratingInput with an unhashed email
	rating := *ratingInput
	rating.UserEmail = string(hash)
	_, err = rr.Database.NamedExec(`INSERT INTO ratings 
								 	VALUES (:userEmail, 
										 	:reviewId, 
										 	:helpful)`, rating)
	if err != nil {
		log.Println("Error in UpdateRating:", err)
		return ratingInput, err
	}

	return ratingInput, nil
}