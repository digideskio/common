package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Plan struct {
	ID           int        `json:"id" db:"id"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	Name         string     `json:"name" db:"name"`
	Price        int        `json:"price" db:"price"`
	State        string     `json:"state" db:"state"`
	Key          string     `json:"key" db:"key"`
	Entitlements slices.Map `json:"entitlements" db:"entitlements"`
}

// String is not required by pop and may be deleted
func (p Plan) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Plans is not required by pop and may be deleted
type Plans []Plan

// String is not required by pop and may be deleted
func (p Plans) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Plan) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.ID, Name: "ID"},
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.IntIsPresent{Field: p.Price, Name: "Price"},
		&validators.StringIsPresent{Field: p.State, Name: "State"},
		&validators.StringIsPresent{Field: p.Key, Name: "Key"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Plan) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Plan) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Plan) TableName() string { return "plans" }
