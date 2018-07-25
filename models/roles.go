package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Role struct {
	ID           int        `json:"id" db:"id"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	RoleableID   int        `json:"roleable_id" db:"roleable_id"`
	Name         string     `json:"name" db:"name"`
	Permissions  slices.Map `json:"permissions" db:"permissions"`
	Admin        bool       `json:"admin" db:"admin"`
	RoleableType string     `json:"roleable_type" db:"roleable_type"`
}

// String is not required by pop and may be deleted
func (r Role) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Roles is not required by pop and may be deleted
type Roles []Role

// String is not required by pop and may be deleted
func (r Roles) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Role) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: r.ID, Name: "ID"},
		&validators.IntIsPresent{Field: r.RoleableID, Name: "RoleableID"},
		&validators.StringIsPresent{Field: r.Name, Name: "Name"},
		&validators.StringIsPresent{Field: r.RoleableType, Name: "RoleableType"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Role) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Role) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Role) TableName() string { return "roles" }
