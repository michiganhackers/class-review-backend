package services

import (
	"class-review-backend/repositories"

	"class-review-backend/models"
)

type IProfessorService interface {
	GetAllProfessors() (*[]models.Professor, error)
	GetProfessorByUniqname(string) (*models.Professor, error)
	PostProfessor(*models.Professor) error
	UpdateProfessor(*models.Professor, string) (*models.Professor, error)
	DeleteProfessor(string) error
	GetProfessorStats() (*[]models.ProfessorStats, error)
	GetProfessorStatsByUniqname(string) (*models.ProfessorStats, error)
}

// Implements IProfessorService
type ProfessorService struct {
	Repositories *repositories.Repositories
}

func DefaultProfessorServices(repos *repositories.Repositories) *ProfessorService {
	return &ProfessorService{
		Repositories: repos,
	}
}

func (ps *ProfessorService) GetAllProfessors() (*[]models.Professor, error) {
	return ps.Repositories.ProfessorRepository.GetAllProfessors()
}

func (ps *ProfessorService) GetProfessorByUniqname(uniqname string) (*models.Professor, error) {
	return ps.Repositories.ProfessorRepository.GetProfessorByUniqname(uniqname)
}

func (ps *ProfessorService) PostProfessor(professorInput *models.Professor) error {
	return ps.Repositories.ProfessorRepository.PostProfessor(professorInput)
}

func (ps *ProfessorService) UpdateProfessor(professorInput *models.Professor, uniqname string) (*models.Professor, error) {
	return ps.Repositories.ProfessorRepository.UpdateProfessor(professorInput, uniqname)
}

func (ps *ProfessorService) DeleteProfessor(uniqname string) error {
	return ps.Repositories.ProfessorRepository.DeleteProfessor(uniqname)
}

func (ps *ProfessorService) GetProfessorStats() (*[]models.ProfessorStats, error) {
	return ps.Repositories.ProfessorRepository.GetProfessorStats()
}

func (ps *ProfessorService) GetProfessorStatsByUniqname(uniqname string) (*models.ProfessorStats, error) {
	return ps.Repositories.ProfessorRepository.GetProfessorStatsByUniqname(uniqname)
}