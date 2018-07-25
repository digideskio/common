package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Authentication struct {
	ID        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	UserID    int          `json:"user_id" db:"user_id"`
	Provider  string       `json:"provider" db:"provider"`
	Token     string       `json:"token" db:"token"`
	Uid       string       `json:"uid" db:"uid"`
	Username  nulls.String `json:"username" db:"username"`
}

// String is not required by pop and may be deleted
func (a Authentication) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Authentications is not required by pop and may be deleted
type Authentications []Authentication

// String is not required by pop and may be deleted
func (a Authentications) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Authentication) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: a.ID, Name: "ID"},
		&validators.IntIsPresent{Field: a.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: a.Provider, Name: "Provider"},
		&validators.StringIsPresent{Field: a.Token, Name: "Token"},
		&validators.StringIsPresent{Field: a.Uid, Name: "Uid"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Authentication) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Authentication) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x Authentication) TableName() string { return "authentications" }
