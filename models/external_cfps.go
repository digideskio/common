package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type ExternalCfp struct {
	ID                int          `json:"id" db:"id"`
	CreatedAt         time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at" db:"updated_at"`
	Name              string       `json:"name" db:"name"`
	Location          string       `json:"location" db:"location"`
	Link              string       `json:"link" db:"link"`
	DueDate           time.Time    `json:"due_date" db:"due_date"`
	StartDate         time.Time    `json:"start_date" db:"start_date"`
	EndDate           time.Time    `json:"end_date" db:"end_date"`
	TravelDescription nulls.String `json:"travel_description" db:"travel_description"`
	OrganizerEmail    nulls.String `json:"organizer_email" db:"organizer_email"`
	CodeOfConductLink nulls.String `json:"code_of_conduct_link" db:"code_of_conduct_link"`
	ConferenceWebsite nulls.String `json:"conference_website" db:"conference_website"`
}

// String is not required by pop and may be deleted
func (e ExternalCfp) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// ExternalCfps is not required by pop and may be deleted
type ExternalCfps []ExternalCfp

// String is not required by pop and may be deleted
func (e ExternalCfps) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *ExternalCfp) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: e.ID, Name: "ID"},
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.StringIsPresent{Field: e.Location, Name: "Location"},
		&validators.StringIsPresent{Field: e.Link, Name: "Link"},
		&validators.TimeIsPresent{Field: e.DueDate, Name: "DueDate"},
		&validators.TimeIsPresent{Field: e.StartDate, Name: "StartDate"},
		&validators.TimeIsPresent{Field: e.EndDate, Name: "EndDate"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *ExternalCfp) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *ExternalCfp) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x ExternalCfp) TableName() string { return "external_cfps" }
