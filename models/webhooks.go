package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Webhook struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	EventID   nulls.Int `json:"event_id" db:"event_id"`
	URL       string    `json:"url" db:"url"`
	State     string    `json:"state" db:"state"`
}

// String is not required by pop and may be deleted
func (w Webhook) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Webhooks is not required by pop and may be deleted
type Webhooks []Webhook

// String is not required by pop and may be deleted
func (w Webhooks) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *Webhook) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: w.ID, Name: "ID"},
		&validators.StringIsPresent{Field: w.URL, Name: "URL"},
		&validators.StringIsPresent{Field: w.State, Name: "State"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *Webhook) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *Webhook) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Webhook) TableName() string { return "webhooks" }
