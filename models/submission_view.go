package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/validate"
)

type SubmissionView struct {
	ID             nulls.Int     `json:"id" db:"id"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
	TalkID         nulls.Int     `json:"talk_id" db:"talk_id"`
	CfpID          nulls.Int     `json:"cfp_id" db:"cfp_id"`
	ProfileID      nulls.Int     `json:"profile_id" db:"profile_id"`
	State          nulls.String  `json:"state" db:"state"`
	ParentID       nulls.Int     `json:"parent_id" db:"parent_id"`
	EventID        nulls.Int     `json:"event_id" db:"event_id"`
	Confirmed      bool          `json:"confirmed" db:"confirmed"`
	AdditionalInfo nulls.String  `json:"additional_info" db:"additional_info"`
	Title          nulls.String  `json:"title" db:"title"`
	TalkFormatID   nulls.Int     `json:"talk_format_id" db:"talk_format_id"`
	Abstract       nulls.String  `json:"abstract" db:"abstract"`
	Description    nulls.String  `json:"description" db:"description"`
	Notes          nulls.String  `json:"notes" db:"notes"`
	UserID         nulls.Int     `json:"user_id" db:"user_id"`
	AudienceLevel  nulls.String  `json:"audience_level" db:"audience_level"`
	ProfileName    nulls.String  `json:"profile_name" db:"profile_name"`
	Location       nulls.String  `json:"location" db:"location"`
	SummedRating   nulls.Int     `json:"summed_rating" db:"summed_rating"`
	RatingsCount   nulls.Int     `json:"ratings_count" db:"ratings_count"`
	MaxRatingCount nulls.Int     `json:"max_rating_count" db:"max_rating_count"`
	Confidence     nulls.Float32 `json:"confidence" db:"confidence"`
	AverageRating  nulls.Float32 `json:"average_rating" db:"average_rating"`
	NeedsRating    bool          `json:"needs_rating" db:"needs_rating"`
	RaterIds       slices.Int    `json:"rater_ids" db:"rater_ids"`
	Tags           slices.String `json:"tags" db:"tags"`
	Locked         bool          `json:"locked" db:"locked"`
}

// String is not required by pop and may be deleted
func (s SubmissionView) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// SubmissionViews is not required by pop and may be deleted
type SubmissionViews []SubmissionView

// String is not required by pop and may be deleted
func (s SubmissionViews) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *SubmissionView) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *SubmissionView) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *SubmissionView) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x SubmissionView) TableName() string { return "submission_view" }
