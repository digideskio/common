package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type ApiKey struct {
	ID        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Name      nulls.String `json:"name" db:"name"`
	Token     nulls.String `json:"token" db:"token"`
}

// String is not required by pop and may be deleted
func (a ApiKey) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// ApiKeys is not required by pop and may be deleted
type ApiKeys []ApiKey

// String is not required by pop and may be deleted
func (a ApiKeys) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *ApiKey) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: a.ID, Name: "ID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *ApiKey) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *ApiKey) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x ApiKey) TableName() string { return "api_keys" }
