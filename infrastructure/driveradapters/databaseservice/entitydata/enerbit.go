package entitydata

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type EnerBitData struct {
	ID               uuid.UUID    `gorm:"primaryKey" json:"id" gorm:"type:uuid;"`
	Brand            string       `json:"brand"`
	Address          string       `json:"address"`
	InstallationDate time.Time    `json:"installation_date"`
	RetirementDate   sql.NullTime `json:"retirement_date"`
	Serial           string       `json:"serial"`
	Lines            int          `json:"lines"`
	IsActive         bool         `json:"is_active"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdateAt         time.Time    `json:"update_at"`
}
