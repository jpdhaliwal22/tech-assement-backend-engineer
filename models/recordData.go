package models

type RecordData struct {
	Income              int    `json:"income,omitempty"`
	NumberOfCreditCards int    `json:"number_of_credit_cards,omitempty"`
	Age                 int    `json:"age,omitempty"`
	PoliticallyExposed  bool   `json:"politically_exposed,omitempty"`
	JobIndustryCode     string `json:"job_industry_code,omitempty"`
	PhoneNumber         string `json:"phone_number,omitempty"`
}
