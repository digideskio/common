package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type OrgEventLeaderView struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	EventID        nulls.Int `json:"event_id" db:"event_id"`
	Count          nulls.Int `json:"count" db:"count"`
	OrganizationID nulls.Int `json:"organization_id" db:"organization_id"`
}

// String is not required by pop and may be deleted
func (o OrgEventLeaderView) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OrgEventLeaderViews is not required by pop and may be deleted
type OrgEventLeaderViews []OrgEventLeaderView

// String is not required by pop and may be deleted
func (o OrgEventLeaderViews) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OrgEventLeaderView) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OrgEventLeaderView) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OrgEventLeaderView) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x OrgEventLeaderView) TableName() string { return "org_event_leader_view" }
