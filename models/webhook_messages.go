package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type WebhookMessage struct {
	ID        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	WebhookID nulls.Int    `json:"webhook_id" db:"webhook_id"`
	Payload   nulls.String `json:"payload" db:"payload"`
	State     nulls.String `json:"state" db:"state"`
}

// String is not required by pop and may be deleted
func (w WebhookMessage) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// WebhookMessages is not required by pop and may be deleted
type WebhookMessages []WebhookMessage

// String is not required by pop and may be deleted
func (w WebhookMessages) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *WebhookMessage) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: w.ID, Name: "ID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *WebhookMessage) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *WebhookMessage) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x WebhookMessage) TableName() string { return "webhook_messages" }
