package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type OrgMember struct {
	ID             int       `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	OrganizationID nulls.Int `json:"organization_id" db:"organization_id"`
	UserID         nulls.Int `json:"user_id" db:"user_id"`
	RoleID         nulls.Int `json:"role_id" db:"role_id"`
}

// String is not required by pop and may be deleted
func (o OrgMember) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OrgMembers is not required by pop and may be deleted
type OrgMembers []OrgMember

// String is not required by pop and may be deleted
func (o OrgMembers) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OrgMember) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: o.ID, Name: "ID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OrgMember) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OrgMember) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x OrgMember) TableName() string { return "org_members" }
