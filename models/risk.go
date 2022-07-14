package models

// Risk represents the risk involved in credit
type Risk string

// All pre-defined Status
const (
	RiskLow    Risk = "LOW"
	RiskMedium Risk = "MEDIUM"
	RiskHigh   Risk = "HIGH"
)

var (
	statuses = map[string]Risk{
		"LOW":    RiskLow,
		"MEDIUM": RiskMedium,
		"HIGH":   RiskHigh,
	}
)

// String converts a risk to a string representation.
func (s Risk) String() string {
	return string(s)
}
