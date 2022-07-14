package manager

import (
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"github.com/honestbank/tech-assignment-backend-engineer/risk"
	"strconv"
)

func (m *mgr) AddCreditDetails(data *models.RecordData) models.Response {
	if data.Age >= MinimumAgeLimit && data.Income >= MinimumEarningLimit && risk.CalculateCreditRisk(data.Age, data.NumberOfCreditCards) == models.RiskLow.String() &&
		!data.PoliticallyExposed && checkAreaCode(data.PhoneNumber) {
		fmt.Println("here", data.Age)
		return models.Response{Status: "approved"}
	} else {
		return models.Response{Status: "declined"}
	}
}

func checkAreaCode(phoneNumber string) bool {
	phoneNumberPrefix := phoneNumber[0:1]
	phoneNumberValue, _ := strconv.Atoi(phoneNumberPrefix)

	switch phoneNumberValue {
	case 0, 2, 5, 8:
		return true
	}

	return false
}
