package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Organization struct {
	ID        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Name      nulls.String `json:"name" db:"name"`
}

// String is not required by pop and may be deleted
func (o Organization) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Organizations is not required by pop and may be deleted
type Organizations []Organization

// String is not required by pop and may be deleted
func (o Organizations) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *Organization) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: o.ID, Name: "ID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *Organization) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *Organization) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Organization) TableName() string { return "organizations" }
