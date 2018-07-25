package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type User struct {
	ID        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Name      string       `json:"name" db:"name"`
	Email     nulls.String `json:"email" db:"email"`
	Token     string       `json:"token" db:"token"`
	SuperUser bool         `json:"super_user" db:"super_user"`
	Location  string       `json:"location" db:"location"`
	Slug      nulls.String `json:"slug" db:"slug"`
	Avatar    nulls.String `json:"avatar" db:"avatar"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: u.ID, Name: "ID"},
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		&validators.StringIsPresent{Field: u.Token, Name: "Token"},
		&validators.StringIsPresent{Field: u.Location, Name: "Location"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x User) TableName() string { return "users" }
