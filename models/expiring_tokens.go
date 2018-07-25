package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type ExpiringToken struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    nulls.Int `json:"user_id" db:"user_id"`
	Token     string    `json:"token" db:"token"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

// String is not required by pop and may be deleted
func (e ExpiringToken) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// ExpiringTokens is not required by pop and may be deleted
type ExpiringTokens []ExpiringToken

// String is not required by pop and may be deleted
func (e ExpiringTokens) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *ExpiringToken) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: e.ID, Name: "ID"},
		&validators.StringIsPresent{Field: e.Token, Name: "Token"},
		&validators.TimeIsPresent{Field: e.ExpiresAt, Name: "ExpiresAt"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *ExpiringToken) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *ExpiringToken) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x ExpiringToken) TableName() string { return "expiring_tokens" }
