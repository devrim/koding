package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// JGroupLimit j group limit
// swagger:model JGroupLimit
type JGroupLimit struct {

	// allowed instances
	AllowedInstances []string `json:"allowedInstances"`

	// instance per member
	InstancePerMember float64 `json:"instancePerMember,omitempty"`

	// max instance
	MaxInstance float64 `json:"maxInstance,omitempty"`

	// member
	Member float64 `json:"member,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// restrictions
	Restrictions *JGroupLimitRestrictions `json:"restrictions,omitempty"`

	// storage per instance
	StoragePerInstance float64 `json:"storagePerInstance,omitempty"`

	// valid for
	ValidFor float64 `json:"validFor,omitempty"`
}

// Validate validates this j group limit
func (m *JGroupLimit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedInstances(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRestrictions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JGroupLimit) validateAllowedInstances(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedInstances) { // not required
		return nil
	}

	return nil
}

func (m *JGroupLimit) validateRestrictions(formats strfmt.Registry) error {

	if swag.IsZero(m.Restrictions) { // not required
		return nil
	}

	if m.Restrictions != nil {

		if err := m.Restrictions.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// JGroupLimitRestrictions j group limit restrictions
// swagger:model JGroupLimitRestrictions
type JGroupLimitRestrictions struct {

	// custom
	Custom interface{} `json:"custom,omitempty"`

	// provider
	Provider []string `json:"provider"`

	// resource
	Resource []string `json:"resource"`

	// supports
	Supports []string `json:"supports"`
}

// Validate validates this j group limit restrictions
func (m *JGroupLimitRestrictions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupports(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JGroupLimitRestrictions) validateProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	return nil
}

func (m *JGroupLimitRestrictions) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	return nil
}

func (m *JGroupLimitRestrictions) validateSupports(formats strfmt.Registry) error {

	if swag.IsZero(m.Supports) { // not required
		return nil
	}

	return nil
}
