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

type EventView struct {
	ID                nulls.Int     `json:"id" db:"id"`
	CreatedAt         time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at" db:"updated_at"`
	Name              nulls.String  `json:"name" db:"name"`
	Slug              nulls.String  `json:"slug" db:"slug"`
	Description       nulls.String  `json:"description" db:"description"`
	CustomDomain      nulls.String  `json:"custom_domain" db:"custom_domain"`
	Logo              nulls.String  `json:"logo" db:"logo"`
	Location          nulls.String  `json:"location" db:"location"`
	PlanID            nulls.Int     `json:"plan_id" db:"plan_id"`
	Website           nulls.String  `json:"website" db:"website"`
	CodeOfConduct     nulls.String  `json:"code_of_conduct" db:"code_of_conduct"`
	FeaturedPastEvent bool          `json:"featured_past_event" db:"featured_past_event"`
	GaCode            nulls.String  `json:"ga_code" db:"ga_code"`
	Dates             []time.Time   `json:"dates" db:"dates"`
	DisplayFrom       time.Time     `json:"display_from" db:"display_from"`
	DisplayUntil      time.Time     `json:"display_until" db:"display_until"`
	CfpID             nulls.Int     `json:"cfp_id" db:"cfp_id"`
	AdditionalInfo    nulls.String  `json:"additional_info" db:"additional_info"`
	Blind             bool          `json:"blind" db:"blind"`
	BlindRevealedAt   time.Time     `json:"blind_revealed_at" db:"blind_revealed_at"`
	OpensAt           time.Time     `json:"opens_at" db:"opens_at"`
	ClosesAt          time.Time     `json:"closes_at" db:"closes_at"`
	CfpDescription    nulls.String  `json:"cfp_description" db:"cfp_description"`
	State             nulls.String  `json:"state" db:"state"`
	MaxRatingCount    nulls.Int     `json:"max_rating_count" db:"max_rating_count"`
	CfpNotes          nulls.String  `json:"cfp_notes" db:"cfp_notes"`
	RatingStyle       nulls.String  `json:"rating_style" db:"rating_style"`
	TalkFormatNames   slices.Map    `json:"talk_format_names" db:"talk_format_names"`
	TalkFormatTypes   slices.Int    `json:"talk_format_types" db:"talk_format_types"`
	TravelAssistance  bool          `json:"travel_assistance" db:"travel_assistance"`
	TravelPolicy      nulls.String  `json:"travel_policy" db:"travel_policy"`
	PlanName          nulls.String  `json:"plan_name" db:"plan_name"`
	PlanKey           nulls.String  `json:"plan_key" db:"plan_key"`
	Entitlements      slices.Map    `json:"entitlements" db:"entitlements"`
	Tags              slices.String `json:"tags" db:"tags"`
	Paid              bool          `json:"paid" db:"paid"`
}

// String is not required by pop and may be deleted
func (e EventView) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// EventViews is not required by pop and may be deleted
type EventViews []EventView

// String is not required by pop and may be deleted
func (e EventViews) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *EventView) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.TimeIsPresent{Field: e.DisplayFrom, Name: "DisplayFrom"},
		&validators.TimeIsPresent{Field: e.DisplayUntil, Name: "DisplayUntil"},
		&validators.TimeIsPresent{Field: e.BlindRevealedAt, Name: "BlindRevealedAt"},
		&validators.TimeIsPresent{Field: e.OpensAt, Name: "OpensAt"},
		&validators.TimeIsPresent{Field: e.ClosesAt, Name: "ClosesAt"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *EventView) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *EventView) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x EventView) TableName() string { return "event_view" }
