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

type OpenCfpView struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Count     nulls.Int `json:"count" db:"count"`
	Day       time.Time `json:"day" db:"day"`
}

// String is not required by pop and may be deleted
func (o OpenCfpView) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OpenCfpViews is not required by pop and may be deleted
type OpenCfpViews []OpenCfpView

// String is not required by pop and may be deleted
func (o OpenCfpViews) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OpenCfpView) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.TimeIsPresent{Field: o.Day, Name: "Day"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OpenCfpView) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OpenCfpView) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x OpenCfpView) TableName() string { return "open_cfp_view" }
