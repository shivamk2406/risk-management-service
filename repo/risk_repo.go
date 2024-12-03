package repo

import (
	"fmt"

	"github.com/shivamk2406/risk-management-service/models"
)

type API interface {
	GetRisk() ([]*models.Risk, error)
	GetRiskById(id string) (*models.Risk, error)
	CreateRisk(risk *models.Risk) (*models.Risk, error)
}

type RiskRepo struct {
	Risks map[string]*models.Risk
}

func NewRiskRepo() API {
	return &RiskRepo{
		Risks: make(map[string]*models.Risk),
	}
}

func (r *RiskRepo) GetRisk() ([]*models.Risk, error) {
	var risks []*models.Risk

	for _, risk := range r.Risks {
		risks = append(risks, risk)
	}

	return risks, nil
}

func (r *RiskRepo) GetRiskById(id string) (*models.Risk, error) {
	risk, ok := r.Risks[id]
	if !ok {
		return nil, fmt.Errorf("no such risk exist with id %s", id)
	}
	return risk, nil
}

func (r *RiskRepo) CreateRisk(risk *models.Risk) (*models.Risk, error) {
	r.Risks[risk.ID.String()]=risk
	return r.Risks[risk.ID.String()], nil
}
