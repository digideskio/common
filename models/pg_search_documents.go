package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type PgSearchDocument struct {
	ID             int          `json:"id" db:"id"`
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at" db:"updated_at"`
	Content        nulls.String `json:"content" db:"content"`
	SearchableID   nulls.Int    `json:"searchable_id" db:"searchable_id"`
	SearchableType nulls.String `json:"searchable_type" db:"searchable_type"`
}

// String is not required by pop and may be deleted
func (p PgSearchDocument) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// PgSearchDocuments is not required by pop and may be deleted
type PgSearchDocuments []PgSearchDocument

// String is not required by pop and may be deleted
func (p PgSearchDocuments) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *PgSearchDocument) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.ID, Name: "ID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *PgSearchDocument) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *PgSearchDocument) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x PgSearchDocument) TableName() string { return "pg_search_documents" }
