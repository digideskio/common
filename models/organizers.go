package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Organizer struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	EventID   int       `json:"event_id" db:"event_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	RoleID    nulls.Int `json:"role_id" db:"role_id"`
	Token     string    `json:"token" db:"token"`
}

// String is not required by pop and may be deleted
func (o Organizer) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Organizers is not required by pop and may be deleted
type Organizers []Organizer

// String is not required by pop and may be deleted
func (o Organizers) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *Organizer) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: o.ID, Name: "ID"},
		&validators.IntIsPresent{Field: o.EventID, Name: "EventID"},
		&validators.IntIsPresent{Field: o.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: o.Token, Name: "Token"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *Organizer) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *Organizer) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Organizer) TableName() string { return "organizers" }
