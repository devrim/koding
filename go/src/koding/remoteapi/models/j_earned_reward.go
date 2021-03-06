package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JEarnedReward j earned reward
// swagger:model JEarnedReward
type JEarnedReward struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// origin Id
	// Required: true
	OriginID *string `json:"originId"`

	// type
	Type string `json:"type,omitempty"`

	// unit
	Unit string `json:"unit,omitempty"`
}

// Validate validates this j earned reward
func (m *JEarnedReward) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOriginID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JEarnedReward) validateOriginID(formats strfmt.Registry) error {

	if err := validate.Required("originId", "body", m.OriginID); err != nil {
		return err
	}

	return nil
}
