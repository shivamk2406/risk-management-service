package interfaces

import "github.com/shivamk2406/risk-management-service/models"

type RiskService interface {
	GetRisks() ([]*models.Risk, error)
	GetRiskById(id string) (*models.Risk, error)
	CreateRisk(risk *models.Risk) (*models.Risk, error)
}