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

type Talk struct {
	ID            int           `json:"id" db:"id"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`
	UserID        int           `json:"user_id" db:"user_id"`
	Title         string        `json:"title" db:"title"`
	Description   nulls.String  `json:"description" db:"description"`
	Notes         nulls.String  `json:"notes" db:"notes"`
	ParentID      nulls.Int     `json:"parent_id" db:"parent_id"`
	Abstract      string        `json:"abstract" db:"abstract"`
	AudienceLevel string        `json:"audience_level" db:"audience_level"`
	IsPublic      bool          `json:"is_public" db:"is_public"`
	TalkFormatID  nulls.Int     `json:"talk_format_id" db:"talk_format_id"`
	Tags          slices.String `json:"tags" db:"tags"`
	State         nulls.String  `json:"state" db:"state"`
}

// String is not required by pop and may be deleted
func (t Talk) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Talks is not required by pop and may be deleted
type Talks []Talk

// String is not required by pop and may be deleted
func (t Talks) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Talk) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: t.ID, Name: "ID"},
		&validators.IntIsPresent{Field: t.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: t.Title, Name: "Title"},
		&validators.StringIsPresent{Field: t.Abstract, Name: "Abstract"},
		&validators.StringIsPresent{Field: t.AudienceLevel, Name: "AudienceLevel"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Talk) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Talk) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Talk) TableName() string { return "talks" }
