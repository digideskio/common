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

type Cfp struct {
	ID               int           `json:"id" db:"id"`
	CreatedAt        time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at" db:"updated_at"`
	EventID          int           `json:"event_id" db:"event_id"`
	Description      nulls.String  `json:"description" db:"description"`
	Notes            nulls.String  `json:"notes" db:"notes"`
	OpensAt          time.Time     `json:"opens_at" db:"opens_at"`
	ClosesAt         time.Time     `json:"closes_at" db:"closes_at"`
	Blind            bool          `json:"blind" db:"blind"`
	BlindRevealedAt  time.Time     `json:"blind_revealed_at" db:"blind_revealed_at"`
	TravelPolicy     nulls.String  `json:"travel_policy" db:"travel_policy"`
	TravelAssistance bool          `json:"travel_assistance" db:"travel_assistance"`
	TalkFormatTypes  slices.Int    `json:"talk_format_types" db:"talk_format_types"`
	AdditionalInfo   nulls.String  `json:"additional_info" db:"additional_info"`
	TalkFormatNames  slices.Map    `json:"talk_format_names" db:"talk_format_names"`
	RatingStyle      string        `json:"rating_style" db:"rating_style"`
	MaxRatingCount   int           `json:"max_rating_count" db:"max_rating_count"`
	State            nulls.String  `json:"state" db:"state"`
	Tags             slices.String `json:"tags" db:"tags"`
}

// String is not required by pop and may be deleted
func (c Cfp) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Cfps is not required by pop and may be deleted
type Cfps []Cfp

// String is not required by pop and may be deleted
func (c Cfps) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Cfp) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: c.ID, Name: "ID"},
		&validators.IntIsPresent{Field: c.EventID, Name: "EventID"},
		&validators.TimeIsPresent{Field: c.OpensAt, Name: "OpensAt"},
		&validators.TimeIsPresent{Field: c.ClosesAt, Name: "ClosesAt"},
		&validators.TimeIsPresent{Field: c.BlindRevealedAt, Name: "BlindRevealedAt"},
		&validators.StringIsPresent{Field: c.RatingStyle, Name: "RatingStyle"},
		&validators.IntIsPresent{Field: c.MaxRatingCount, Name: "MaxRatingCount"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Cfp) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Cfp) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Cfp) TableName() string { return "cfps" }
