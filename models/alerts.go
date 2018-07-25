package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Alert struct {
	ID               int          `json:"id" db:"id"`
	CreatedAt        time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at" db:"updated_at"`
	UserID           nulls.Int    `json:"user_id" db:"user_id"`
	State            nulls.String `json:"state" db:"state"`
	Schedule         nulls.String `json:"schedule" db:"schedule"`
	TravelAssistance bool         `json:"travel_assistance" db:"travel_assistance"`
	LastAlertedAt    time.Time    `json:"last_alerted_at" db:"last_alerted_at"`
	Keywords         nulls.String `json:"keywords" db:"keywords"`
}

// String is not required by pop and may be deleted
func (a Alert) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Alerts is not required by pop and may be deleted
type Alerts []Alert

// String is not required by pop and may be deleted
func (a Alerts) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Alert) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: a.ID, Name: "ID"},
		&validators.TimeIsPresent{Field: a.LastAlertedAt, Name: "LastAlertedAt"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Alert) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Alert) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Alert) TableName() string { return "alerts" }
