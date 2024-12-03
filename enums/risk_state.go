package enums

const (
	UnspecifiedTenantRequest RiskState = iota
	OPEN
	CLOSED
	ACCEPTED
	INVESTIGATING
)

type RiskState int