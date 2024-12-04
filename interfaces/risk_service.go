package interfaces

import "github.com/shivamk2406/risk-management-service/internal/models"

type RiskService interface {
	GetRisks() ([]*models.Risk, error)
	GetRiskById(id string) (*models.Risk, error)
	CreateRisk(risk *models.RiskRequestDto) (*models.Risk, error)
}