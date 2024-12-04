package models

import (
	"github.com/google/uuid"
	"github.com/shivamk2406/risk-management-service/enums"
)

type Risk struct {
	ID          uuid.UUID
	State       enums.RiskState
	Title       string
	Description string
}

type RiskRequestDto struct {
	State       string
	Title       string
	Description string
}
