package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Submission struct {
	ID             int           `json:"id" db:"id"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
	TalkID         int           `json:"talk_id" db:"talk_id"`
	CfpID          int           `json:"cfp_id" db:"cfp_id"`
	ProfileID      int           `json:"profile_id" db:"profile_id"`
	State          string        `json:"state" db:"state"`
	ParentID       int           `json:"parent_id" db:"parent_id"`
	EventID        int           `json:"event_id" db:"event_id"`
	Confirmed      bool          `json:"confirmed" db:"confirmed"`
	AdditionalInfo nulls.String  `json:"additional_info" db:"additional_info"`
	UserID         int           `json:"user_id" db:"user_id"`
	Locked         bool          `json:"locked" db:"locked"`
	Tags           slices.String `json:"tags" db:"tags"`
}

// String is not required by pop and may be deleted
func (s Submission) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Submissions is not required by pop and may be deleted
type Submissions []Submission

// String is not required by pop and may be deleted
func (s Submissions) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Submission) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: s.ID, Name: "ID"},
		&validators.IntIsPresent{Field: s.TalkID, Name: "TalkID"},
		&validators.IntIsPresent{Field: s.CfpID, Name: "CfpID"},
		&validators.IntIsPresent{Field: s.ProfileID, Name: "ProfileID"},
		&validators.StringIsPresent{Field: s.State, Name: "State"},
		&validators.IntIsPresent{Field: s.ParentID, Name: "ParentID"},
		&validators.IntIsPresent{Field: s.EventID, Name: "EventID"},
		&validators.IntIsPresent{Field: s.UserID, Name: "UserID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Submission) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Submission) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Submission) TableName() string { return "submissions" }
