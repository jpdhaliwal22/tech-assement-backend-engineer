package manager

import (
	"github.com/honestbank/tech-assignment-backend-engineer/models"
)

type Manager interface {
	AddCreditDetails(data *models.RecordData) models.Response
}

type mgr struct {
}

func NewManager() Manager {
	return &mgr{}
}
