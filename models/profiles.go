package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Profile struct {
	ID        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	UserID    int          `json:"user_id" db:"user_id"`
	Bio       nulls.String `json:"bio" db:"bio"`
	Company   nulls.String `json:"company" db:"company"`
	Twitter   nulls.String `json:"twitter" db:"twitter"`
	URL       nulls.String `json:"url" db:"url"`
	Name      string       `json:"name" db:"name"`
	ShirtSize nulls.String `json:"shirt_size" db:"shirt_size"`
}

// String is not required by pop and may be deleted
func (p Profile) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Profiles is not required by pop and may be deleted
type Profiles []Profile

// String is not required by pop and may be deleted
func (p Profiles) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Profile) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.ID, Name: "ID"},
		&validators.IntIsPresent{Field: p.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Profile) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Profile) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Profile) TableName() string { return "profiles" }
