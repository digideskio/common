package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type LifecycleEmailLog struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Type      string    `json:"type" db:"type"`
	SentAt    time.Time `json:"sent_at" db:"sent_at"`
	UserID    int       `json:"user_id" db:"user_id"`
	EventID   nulls.Int `json:"event_id" db:"event_id"`
}

// String is not required by pop and may be deleted
func (l LifecycleEmailLog) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// LifecycleEmailLogs is not required by pop and may be deleted
type LifecycleEmailLogs []LifecycleEmailLog

// String is not required by pop and may be deleted
func (l LifecycleEmailLogs) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *LifecycleEmailLog) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: l.ID, Name: "ID"},
		&validators.StringIsPresent{Field: l.Type, Name: "Type"},
		&validators.TimeIsPresent{Field: l.SentAt, Name: "SentAt"},
		&validators.IntIsPresent{Field: l.UserID, Name: "UserID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *LifecycleEmailLog) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *LifecycleEmailLog) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x LifecycleEmailLog) TableName() string { return "lifecycle_email_logs" }
