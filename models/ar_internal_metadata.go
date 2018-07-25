package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type ArInternalMetadatum struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Key       string       `json:"key" db:"key"`
	Value     nulls.String `json:"value" db:"value"`
}

// String is not required by pop and may be deleted
func (a ArInternalMetadatum) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// ArInternalMetadata is not required by pop and may be deleted
type ArInternalMetadata []ArInternalMetadatum

// String is not required by pop and may be deleted
func (a ArInternalMetadata) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *ArInternalMetadatum) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Key, Name: "Key"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *ArInternalMetadatum) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *ArInternalMetadatum) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x ArInternalMetadatum) TableName() string { return "ar_internal_metadata" }
