// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// System system
//
// swagger:model System
type System struct {

	// The host available Processor units
	// Required: true
	Cores *float64 `json:"cores"`

	// The host identifier
	ID int64 `json:"id,omitempty"`

	// The host available RAM memory in GiB
	// Required: true
	Memory *int64 `json:"memory"`
}

// Validate validates this system
func (m *System) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *System) validateCores(formats strfmt.Registry) error {

	if err := validate.Required("cores", "body", m.Cores); err != nil {
		return err
	}

	return nil
}

func (m *System) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this system based on context it is used
func (m *System) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *System) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *System) UnmarshalBinary(b []byte) error {
	var res System
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
