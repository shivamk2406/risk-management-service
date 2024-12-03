package service

import (
	"github.com/shivamk2406/risk-management-service/interfaces"
	"github.com/shivamk2406/risk-management-service/models"
	"github.com/shivamk2406/risk-management-service/repo"
)

type RiskService struct {
	repo repo.API
}


func NewRiskSvc(repo repo.API) interfaces.RiskService{
	return &RiskService{
		repo: repo,
	}
}

func (r *RiskService) GetRisks() ([]*models.Risk, error) {
	return r.repo.GetRisks()
}

func (r *RiskService) GetRiskById(id string) (*models.Risk, error) {
	return r.repo.GetRiskById(id)
}

func (r *RiskService) CreateRisk(risk *models.Risk) (*models.Risk, error) {
	return r.repo.CreateRisk(risk)
}