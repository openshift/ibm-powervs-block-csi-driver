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

// VolumeOnboardingCommon volume onboarding common
//
// swagger:model VolumeOnboardingCommon
type VolumeOnboardingCommon struct {

	// Description of the volume onboarding operation
	Description string `json:"description,omitempty"`

	// Indicates the volume onboarding operation id
	// Required: true
	ID *string `json:"id"`

	// List of volumes requested to be onboarded
	InputVolumes []string `json:"inputVolumes"`

	// Indicates the status of volume onboarding operation
	Status string `json:"status,omitempty"`
}

// Validate validates this volume onboarding common
func (m *VolumeOnboardingCommon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeOnboardingCommon) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this volume onboarding common based on context it is used
func (m *VolumeOnboardingCommon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeOnboardingCommon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeOnboardingCommon) UnmarshalBinary(b []byte) error {
	var res VolumeOnboardingCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}