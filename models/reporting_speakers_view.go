package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type ReportingSpeakersView struct {
	ID              uuid.UUID `json:"id" db:"id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	Users           nulls.Int `json:"users" db:"users"`
	Speakers        nulls.Int `json:"speakers" db:"speakers"`
	ActiveSpeakers  nulls.Int `json:"active_speakers" db:"active_speakers"`
	CrossSubmitters nulls.Int `json:"cross_submitters" db:"cross_submitters"`
}

// String is not required by pop and may be deleted
func (r ReportingSpeakersView) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// ReportingSpeakersViews is not required by pop and may be deleted
type ReportingSpeakersViews []ReportingSpeakersView

// String is not required by pop and may be deleted
func (r ReportingSpeakersViews) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *ReportingSpeakersView) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *ReportingSpeakersView) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *ReportingSpeakersView) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x ReportingSpeakersView) TableName() string { return "reporting_speakers_view" }
