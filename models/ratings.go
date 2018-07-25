package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Rating struct {
	ID           int          `json:"id" db:"id"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	SubmissionID int          `json:"submission_id" db:"submission_id"`
	UserID       int          `json:"user_id" db:"user_id"`
	Value        int          `json:"value" db:"value"`
	Comments     nulls.String `json:"comments" db:"comments"`
	CfpID        int          `json:"cfp_id" db:"cfp_id"`
}

// String is not required by pop and may be deleted
func (r Rating) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Ratings is not required by pop and may be deleted
type Ratings []Rating

// String is not required by pop and may be deleted
func (r Ratings) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Rating) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: r.ID, Name: "ID"},
		&validators.IntIsPresent{Field: r.SubmissionID, Name: "SubmissionID"},
		&validators.IntIsPresent{Field: r.UserID, Name: "UserID"},
		&validators.IntIsPresent{Field: r.Value, Name: "Value"},
		&validators.IntIsPresent{Field: r.CfpID, Name: "CfpID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Rating) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Rating) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Rating) TableName() string { return "ratings" }
