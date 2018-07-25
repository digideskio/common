package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type OrganizerInvitation struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	EventID   nulls.Int `json:"event_id" db:"event_id"`
	InviterID nulls.Int `json:"inviter_id" db:"inviter_id"`
	Email     string    `json:"email" db:"email"`
	Token     string    `json:"token" db:"token"`
	State     string    `json:"state" db:"state"`
	RoleID    nulls.Int `json:"role_id" db:"role_id"`
}

// String is not required by pop and may be deleted
func (o OrganizerInvitation) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OrganizerInvitations is not required by pop and may be deleted
type OrganizerInvitations []OrganizerInvitation

// String is not required by pop and may be deleted
func (o OrganizerInvitations) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OrganizerInvitation) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: o.ID, Name: "ID"},
		&validators.StringIsPresent{Field: o.Email, Name: "Email"},
		&validators.StringIsPresent{Field: o.Token, Name: "Token"},
		&validators.StringIsPresent{Field: o.State, Name: "State"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OrganizerInvitation) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OrganizerInvitation) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x OrganizerInvitation) TableName() string { return "organizer_invitations" }
