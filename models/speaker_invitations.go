package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type SpeakerInvitation struct {
	ID           int       `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	EventID      int       `json:"event_id" db:"event_id"`
	CfpID        int       `json:"cfp_id" db:"cfp_id"`
	InviterID    int       `json:"inviter_id" db:"inviter_id"`
	UserID       nulls.Int `json:"user_id" db:"user_id"`
	SubmissionID nulls.Int `json:"submission_id" db:"submission_id"`
	Email        string    `json:"email" db:"email"`
	Token        string    `json:"token" db:"token"`
	State        string    `json:"state" db:"state"`
	Body         string    `json:"body" db:"body"`
}

// String is not required by pop and may be deleted
func (s SpeakerInvitation) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// SpeakerInvitations is not required by pop and may be deleted
type SpeakerInvitations []SpeakerInvitation

// String is not required by pop and may be deleted
func (s SpeakerInvitations) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *SpeakerInvitation) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: s.ID, Name: "ID"},
		&validators.IntIsPresent{Field: s.EventID, Name: "EventID"},
		&validators.IntIsPresent{Field: s.CfpID, Name: "CfpID"},
		&validators.IntIsPresent{Field: s.InviterID, Name: "InviterID"},
		&validators.StringIsPresent{Field: s.Email, Name: "Email"},
		&validators.StringIsPresent{Field: s.Token, Name: "Token"},
		&validators.StringIsPresent{Field: s.State, Name: "State"},
		&validators.StringIsPresent{Field: s.Body, Name: "Body"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *SpeakerInvitation) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *SpeakerInvitation) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x SpeakerInvitation) TableName() string { return "speaker_invitations" }
