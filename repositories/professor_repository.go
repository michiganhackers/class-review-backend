package repositories

import (
	"class-review-backend/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type IProfessorRepository interface {
	GetAllProfessors() (*[]models.Professor, error)
	GetProfessorByUniqname(string)(*models.Professor, error)
	PostProfessor(*models.Professor) error
	UpdateProfessor(*models.Professor, string) (*models.Professor, error)
	DeleteProfessor(string) error
	GetProfessorStats() (*[]models.ProfessorStats, error)
	GetProfessorStatsByUniqname(string) (*models.ProfessorStats, error)
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

func (pr *ProfessorRepository) GetAllProfessors() (*[]models.Professor, error) {
	var professors []models.Professor
	err := pr.Database.Select(&professors, `SELECT professor_uniqname, 
 												   professor_name 
										    FROM professors`)
	if err != nil {
		log.Println("Error in GetAllProfessors:", err)
		return &professors, err
	}

	return &professors, nil
}

func (pr *ProfessorRepository) GetProfessorByUniqname(uniqname string) (*models.Professor, error) {
	var professor models.Professor
	err := pr.Database.Get(&professor, `SELECT professor_uniqname, 
											   professor_name 
									 	FROM professors WHERE professor_uniqname=?`, uniqname)
	if err != nil {
		log.Println("Error in GetProfessorByUniqname:", err)
		return &professor, err
	}

	return &professor, nil
}

func (pr *ProfessorRepository) PostProfessor(professorInput *models.Professor) error {
	_, err := pr.Database.NamedExec(`INSERT INTO professors 
									 VALUES (:professor_uniqname, 
											 :professor_name)`, *professorInput)
	if err != nil {
		log.Println("Error in PostProfessor:", err)
		return err
	}
	return nil
}

func (pr *ProfessorRepository) UpdateProfessor(professorInput *models.Professor, uniqname string) (*models.Professor, error) {
	professorInput.Uniqname = uniqname
	_, err := pr.Database.NamedExec(`UPDATE professors 
									 SET professor_name = :professor_name
									 WHERE professor_uniqname = :professor_uniqname`, *professorInput)
	if err != nil {
		log.Println("Error in PutProfessor:", err)
		return professorInput, err
	}
	return professorInput, nil
}

func (pr *ProfessorRepository) DeleteProfessor(uniqname string) error {
	_, err := pr.Database.Exec(`DELETE FROM professors WHERE professor_uniqname=?`, uniqname)
	if err != nil {
		log.Println("Error in DeleteProfessor:", err)
		return err
	}

	return nil
}

func (pr *ProfessorRepository) GetProfessorStats() (*[]models.ProfessorStats, error) {
	var professorStats []models.ProfessorStats
	var professorUniqnames []string
	err := pr.Database.Select(&professorUniqnames, `SELECT DISTINCT professor_uniqname FROM reviews`)
	if err != nil {
		log.Println("Error in GetProfessorStats:", err)
		return &professorStats, err
	}

	for _, element := range professorUniqnames {
		individualProfStats, err := pr.GetProfessorStatsByUniqname(element)
		if err != nil {
			log.Println("Error in GetProfessorStats:", err)
			return &professorStats, err
		}
		professorStats = append(professorStats, *individualProfStats)
	}

	return &professorStats, nil
}

func (pr *ProfessorRepository) GetProfessorStatsByUniqname(uniqname string) (*models.ProfessorStats, error) {
	var professorStats models.ProfessorStats
	err := pr.Database.Get(&professorStats, `SELECT professor_uniqname, 
													AVG(rating), 
													AVG(difficulty), 
													AVG(interest), 
													SUM(helpfulCount), 
													SUM(notHelpfulCount)
											 FROM reviews WHERE professor_uniqname=?`, uniqname)
	if err != nil {
		log.Println("Error in GetProfessorStatsByName:", err)
		return &professorStats, err
	}

	return &professorStats, nil
}