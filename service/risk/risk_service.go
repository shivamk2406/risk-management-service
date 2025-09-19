package service

import (
	"github.com/google/uuid"
	"github.com/shivamk2406/risk-management-service/enums"
	"github.com/shivamk2406/risk-management-service/interfaces"
	"github.com/shivamk2406/risk-management-service/internal/models"
	"github.com/shivamk2406/risk-management-service/internal/repo"
)

type RiskService struct {
	repo repo.API
}

func NewRiskSvc(repo repo.API) interfaces.RiskService {
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

func (r *RiskService) CreateRisk(risk *models.RiskRequestDto) (*models.Risk, error) {
	riskState, err := enums.RiskStateString(risk.State)
	if err != nil {
		return nil, err
	}
	return r.repo.CreateRisk(&models.Risk{
		ID:          uuid.New(),
		State:       riskState,
		Title:       risk.Title,
		Description: risk.Description,
	})
}
