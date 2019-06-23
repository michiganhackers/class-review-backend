package repositories

import (
	"class-review-backend/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type IProfessorRepository interface {
	GetProfessors() (*[]models.Professor, error)
	GetProfessorByReviewID(uint64) (*models.Professor, error)
	GetProfessorsByCourseID(uint64) (*[]models.Professor, error)
	GetProfessorStats() (*[]models.ProfessorStats, error)
	GetProfessorStatsByName(string) (*models.ProfessorStats, error)
}

// Implements IProfessorRepository
type ProfessorRepository struct {
	Database *sqlx.DB
}

func DefaultProfessorRepository(db *sqlx.DB) *ProfessorRepository {
	return &ProfessorRepository{
		Database: db,
	}
}

func (pr *ProfessorRepository) GetProfessors() (*[]models.Professor, error) {
	var professors []models.Professor
	err := pr.Database.Select(&professors, `SELECT DISTINCT professor_name FROM reviews`)
	if err != nil {
		log.Println("Error in GetProfessors: ", err)
		return &professors, err
	}

	return &professors, nil
}

func (pr *ProfessorRepository) GetProfessorByReviewID(id uint64) (*models.Professor, error) {
	var professor models.Professor
	err := pr.Database.Get(&professor, `SELECT professor_name FROM reviews WHERE id=?`, id)
	if err != nil {
		log.Println("Error in GetProfessorByReviewID: ", err)
		return &professor, err
	}

	return &professor, nil
}

func (pr *ProfessorRepository) GetProfessorsByCourseID(id uint64) (*[]models.Professor, error) {
	var professors []models.Professor
	err := pr.Database.Select(&professors, `SELECT DISTINCT professor_name FROM reviews WHERE courseID=?`, id)
	if err != nil {
		log.Println("Error in GetProfessorsByCourseID: ", err)
		return &professors, err
	}

	return &professors, nil
}

func (pr *ProfessorRepository) GetProfessorStats() (*[]models.ProfessorStats, error) {
	var professorStats []models.ProfessorStats
	professors, err := pr.GetProfessors()
	if err != nil {
		log.Println("Error in GetProfessorStats: ", err)
		return &professorStats, err
	}

	for _, element := range *professors {
		individualProfStats, err := pr.GetProfessorStatsByName(element.Name)
		if err != nil {
			log.Println("Error in GetProfessorStats: ", err)
			return &professorStats, err
		}
		professorStats = append(professorStats, *individualProfStats)
	}

	return &professorStats, nil
}

func (pr *ProfessorRepository) GetProfessorStatsByName(name string) (*models.ProfessorStats, error) {
	var professorStats models.ProfessorStats
	err := pr.Database.Get(&professorStats, `SELECT professor_name,
													AVG(rating), 
													AVG(difficulty), 
													AVG(interest), 
													SUM(helpfulCount), 
													SUM(notHelpfulCount)
											 FROM reviews WHERE professor_name=?`, name)

	if err != nil {
		log.Println("Error in GetProfessorStatsByName: ", err)
		return &professorStats, err
	}

	return &professorStats, nil
}