package entities

import "time"

type EnerBitEntities struct {
	ID               string
	Brand            string
	Address          string
	InstallationDate time.Time
	RetirementDate   time.Time
	Serial           string
	Lines            int
	IsActive         bool
	CreatedAt        time.Time
}
