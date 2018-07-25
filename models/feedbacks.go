package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Feedback struct {
	ID           int       `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	SubmissionID nulls.Int `json:"submission_id" db:"submission_id"`
	TalkID       nulls.Int `json:"talk_id" db:"talk_id"`
	UserID       nulls.Int `json:"user_id" db:"user_id"`
	Body         string    `json:"body" db:"body"`
}

// String is not required by pop and may be deleted
func (f Feedback) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Feedbacks is not required by pop and may be deleted
type Feedbacks []Feedback

// String is not required by pop and may be deleted
func (f Feedbacks) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Feedback) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: f.ID, Name: "ID"},
		&validators.StringIsPresent{Field: f.Body, Name: "Body"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Feedback) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Feedback) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Feedback) TableName() string { return "feedbacks" }
