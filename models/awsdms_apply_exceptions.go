package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type AwsdmsApplyException struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	TASKNAME   string    `json:"t_a_s_k_n_a_m_e" db:"t_a_s_k_n_a_m_e"`
	TABLEOWNER string    `json:"t_a_b_l_e_o_w_n_e_r" db:"t_a_b_l_e_o_w_n_e_r"`
	TABLENAME  string    `json:"t_a_b_l_e_n_a_m_e" db:"t_a_b_l_e_n_a_m_e"`
	ERRORTIME  time.Time `json:"e_r_r_o_r_t_i_m_e" db:"e_r_r_o_r_t_i_m_e"`
	STATEMENT  string    `json:"s_t_a_t_e_m_e_n_t" db:"s_t_a_t_e_m_e_n_t"`
	ERROR      string    `json:"e_r_r_o_r" db:"e_r_r_o_r"`
}

// String is not required by pop and may be deleted
func (a AwsdmsApplyException) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// AwsdmsApplyExceptions is not required by pop and may be deleted
type AwsdmsApplyExceptions []AwsdmsApplyException

// String is not required by pop and may be deleted
func (a AwsdmsApplyExceptions) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *AwsdmsApplyException) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.TASKNAME, Name: "TASKNAME"},
		&validators.StringIsPresent{Field: a.TABLEOWNER, Name: "TABLEOWNER"},
		&validators.StringIsPresent{Field: a.TABLENAME, Name: "TABLENAME"},
		&validators.TimeIsPresent{Field: a.ERRORTIME, Name: "ERRORTIME"},
		&validators.StringIsPresent{Field: a.STATEMENT, Name: "STATEMENT"},
		&validators.StringIsPresent{Field: a.ERROR, Name: "ERROR"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *AwsdmsApplyException) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *AwsdmsApplyException) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (x AwsdmsApplyException) TableName() string { return "awsdms_apply_exceptions" }
