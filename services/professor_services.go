package services

import (
	"class-review-backend/repositories"

	"class-review-backend/models"
)

type IProfessorService interface {
	GetProfessors() (*[]models.Professor, error)
	GetProfessorByReviewID(uint64) (*models.Professor, error)
	GetProfessorsByCourseID(uint64) (*[]models.Professor, error)
	GetProfessorStats() (*[]models.ProfessorStats, error)
	GetProfessorStatsByName(string) (*models.ProfessorStats, error)
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

func (ps *ProfessorService) GetProfessors() (*[]models.Professor, error) {
	return ps.Repositories.ProfessorRepository.GetProfessors()
}

func (ps *ProfessorService) GetProfessorByReviewID(id uint64) (*models.Professor, error) {
	return ps.Repositories.ProfessorRepository.GetProfessorByReviewID(id)
}

func (ps *ProfessorService) GetProfessorsByCourseID(id uint64) (*[]models.Professor, error) {
	return ps.Repositories.ProfessorRepository.GetProfessorsByCourseID(id)
}

func (ps *ProfessorService) GetProfessorStats() (*[]models.ProfessorStats, error) {
	return ps.Repositories.ProfessorRepository.GetProfessorStats()
}

func (ps *ProfessorService) GetProfessorStatsByName(name string) (*models.ProfessorStats, error) {
	return ps.Repositories.ProfessorRepository.GetProfessorStatsByName(name)
}