package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Event struct {
	ID                int          `json:"id" db:"id"`
	CreatedAt         time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at" db:"updated_at"`
	Name              string       `json:"name" db:"name"`
	Slug              string       `json:"slug" db:"slug"`
	Description       nulls.String `json:"description" db:"description"`
	CustomDomain      nulls.String `json:"custom_domain" db:"custom_domain"`
	Logo              nulls.String `json:"logo" db:"logo"`
	Location          nulls.String `json:"location" db:"location"`
	PlanID            int          `json:"plan_id" db:"plan_id"`
	Website           nulls.String `json:"website" db:"website"`
	CodeOfConduct     nulls.String `json:"code_of_conduct" db:"code_of_conduct"`
	FeaturedPastEvent bool         `json:"featured_past_event" db:"featured_past_event"`
	GaCode            nulls.String `json:"ga_code" db:"ga_code"`
	Dates             []time.Time  `json:"dates" db:"dates"`
	DisplayFrom       time.Time    `json:"display_from" db:"display_from"`
	DisplayUntil      time.Time    `json:"display_until" db:"display_until"`
}

// String is not required by pop and may be deleted
func (e Event) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Events is not required by pop and may be deleted
type Events []Event

// String is not required by pop and may be deleted
func (e Events) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Event) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: e.ID, Name: "ID"},
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.StringIsPresent{Field: e.Slug, Name: "Slug"},
		&validators.IntIsPresent{Field: e.PlanID, Name: "PlanID"},
		&validators.TimeIsPresent{Field: e.DisplayFrom, Name: "DisplayFrom"},
		&validators.TimeIsPresent{Field: e.DisplayUntil, Name: "DisplayUntil"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Event) TableName() string { return "events" }
