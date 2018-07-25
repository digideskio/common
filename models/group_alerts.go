package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type GroupAlert struct {
	ID             int          `json:"id" db:"id"`
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at" db:"updated_at"`
	OrganizationID nulls.Int    `json:"organization_id" db:"organization_id"`
	State          nulls.String `json:"state" db:"state"`
	Schedule       nulls.String `json:"schedule" db:"schedule"`
	LastAlertedAt  time.Time    `json:"last_alerted_at" db:"last_alerted_at"`
	Keywords       nulls.String `json:"keywords" db:"keywords"`
	UserIds        slices.Int   `json:"user_ids" db:"user_ids"`
}

// String is not required by pop and may be deleted
func (g GroupAlert) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// GroupAlerts is not required by pop and may be deleted
type GroupAlerts []GroupAlert

// String is not required by pop and may be deleted
func (g GroupAlerts) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (g *GroupAlert) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: g.ID, Name: "ID"},
		&validators.TimeIsPresent{Field: g.LastAlertedAt, Name: "LastAlertedAt"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (g *GroupAlert) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (g *GroupAlert) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x GroupAlert) TableName() string { return "group_alerts" }
