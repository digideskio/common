package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Attendee struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    nulls.Int `json:"user_id" db:"user_id"`
	EventID   nulls.Int `json:"event_id" db:"event_id"`
}

// String is not required by pop and may be deleted
func (a Attendee) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Attendees is not required by pop and may be deleted
type Attendees []Attendee

// String is not required by pop and may be deleted
func (a Attendees) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Attendee) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: a.ID, Name: "ID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Attendee) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Attendee) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Attendee) TableName() string { return "attendees" }
