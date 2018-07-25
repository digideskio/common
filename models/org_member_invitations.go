package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type OrgMemberInvitation struct {
	ID             int       `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	OrganizationID nulls.Int `json:"organization_id" db:"organization_id"`
	InviterID      nulls.Int `json:"inviter_id" db:"inviter_id"`
	RoleID         nulls.Int `json:"role_id" db:"role_id"`
	Email          string    `json:"email" db:"email"`
	Token          string    `json:"token" db:"token"`
	State          string    `json:"state" db:"state"`
}

// String is not required by pop and may be deleted
func (o OrgMemberInvitation) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OrgMemberInvitations is not required by pop and may be deleted
type OrgMemberInvitations []OrgMemberInvitation

// String is not required by pop and may be deleted
func (o OrgMemberInvitations) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OrgMemberInvitation) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: o.ID, Name: "ID"},
		&validators.StringIsPresent{Field: o.Email, Name: "Email"},
		&validators.StringIsPresent{Field: o.Token, Name: "Token"},
		&validators.StringIsPresent{Field: o.State, Name: "State"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OrgMemberInvitation) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OrgMemberInvitation) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x OrgMemberInvitation) TableName() string { return "org_member_invitations" }
