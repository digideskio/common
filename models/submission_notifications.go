package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type SubmissionNotification struct {
	ID           int       `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	SubmissionID int       `json:"submission_id" db:"submission_id"`
	UserID       int       `json:"user_id" db:"user_id"`
	CfpID        int       `json:"cfp_id" db:"cfp_id"`
	EventID      int       `json:"event_id" db:"event_id"`
	Body         string    `json:"body" db:"body"`
	State        string    `json:"state" db:"state"`
}

// String is not required by pop and may be deleted
func (s SubmissionNotification) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// SubmissionNotifications is not required by pop and may be deleted
type SubmissionNotifications []SubmissionNotification

// String is not required by pop and may be deleted
func (s SubmissionNotifications) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *SubmissionNotification) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: s.ID, Name: "ID"},
		&validators.IntIsPresent{Field: s.SubmissionID, Name: "SubmissionID"},
		&validators.IntIsPresent{Field: s.UserID, Name: "UserID"},
		&validators.IntIsPresent{Field: s.CfpID, Name: "CfpID"},
		&validators.IntIsPresent{Field: s.EventID, Name: "EventID"},
		&validators.StringIsPresent{Field: s.Body, Name: "Body"},
		&validators.StringIsPresent{Field: s.State, Name: "State"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *SubmissionNotification) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *SubmissionNotification) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x SubmissionNotification) TableName() string { return "submission_notifications" }
