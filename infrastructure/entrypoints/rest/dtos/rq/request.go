package rq

import "time"

type Request struct {
	ID               string    `json:"id"`
	Brand            string    `json:"brand"`
	Address          string    `json:"address"`
	InstallationDate time.Time `json:"installation_date"`
	RetirementDate   time.Time `json:"retirement_date"`
	Serial           string    `json:"serial"`
	Lines            int       `json:"lines"`
	IsActive         bool      `json:"active"`
	CreatedAt        time.Time `json:"created_at"`
}
