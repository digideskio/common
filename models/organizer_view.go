package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type OrganizerView struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	UserID    nulls.Int    `json:"user_id" db:"user_id"`
	Name      nulls.String `json:"name" db:"name"`
	Email     nulls.String `json:"email" db:"email"`
	Location  nulls.String `json:"location" db:"location"`
	EventID   nulls.Int    `json:"event_id" db:"event_id"`
	EventName nulls.String `json:"event_name" db:"event_name"`
	State     nulls.String `json:"state" db:"state"`
}

// String is not required by pop and may be deleted
func (o OrganizerView) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OrganizerViews is not required by pop and may be deleted
type OrganizerViews []OrganizerView

// String is not required by pop and may be deleted
func (o OrganizerViews) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OrganizerView) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OrganizerView) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OrganizerView) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x OrganizerView) TableName() string { return "organizer_view" }
