package enums

const (
	UnspecifiedRiskStateRequest RiskState = iota
	OPEN
	CLOSED
	ACCEPTED
	INVESTIGATING
)

type RiskState int

var RiskStateMap = map[string]RiskState{
	"OPEN":          OPEN,
	"CLOSED":        CLOSED,
	"ACCEPTED":      ACCEPTED,
	"INVESTIGATING": INVESTIGATING,
}
