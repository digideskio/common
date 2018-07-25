package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type SchemaMigration struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Version   string    `json:"version" db:"version"`
}

// String is not required by pop and may be deleted
func (s SchemaMigration) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// SchemaMigrations is not required by pop and may be deleted
type SchemaMigrations []SchemaMigration

// String is not required by pop and may be deleted
func (s SchemaMigrations) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *SchemaMigration) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Version, Name: "Version"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *SchemaMigration) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *SchemaMigration) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x SchemaMigration) TableName() string { return "schema_migrations" }
