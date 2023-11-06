// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Datacenter datacenter
//
// swagger:model Datacenter
type Datacenter struct {

	// Datacenter Capabilities
	// Required: true
	Capabilities map[string]bool `json:"capabilities"`

	// The Datacenter location
	// Required: true
	Location *DatacenterLocation `json:"location"`

	// The Datacenter status
	// Required: true
	// Enum: [ACTIVE MAINTENENCE DOWN]
	Status *string `json:"status"`

	// The Datacenter type
	// Required: true
	// Enum: [Public Cloud Private Cloud]
	Type *string `json:"type"`
}

// Validate validates this datacenter
func (m *Datacenter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Datacenter) validateCapabilities(formats strfmt.Registry) error {

	if err := validate.Required("capabilities", "body", m.Capabilities); err != nil {
		return err
	}

	return nil
}

func (m *Datacenter) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

var datacenterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","MAINTENENCE","DOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datacenterTypeStatusPropEnum = append(datacenterTypeStatusPropEnum, v)
	}
}

const (

	// DatacenterStatusACTIVE captures enum value "ACTIVE"
	DatacenterStatusACTIVE string = "ACTIVE"

	// DatacenterStatusMAINTENENCE captures enum value "MAINTENENCE"
	DatacenterStatusMAINTENENCE string = "MAINTENENCE"

	// DatacenterStatusDOWN captures enum value "DOWN"
	DatacenterStatusDOWN string = "DOWN"
)

// prop value enum
func (m *Datacenter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datacenterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Datacenter) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var datacenterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Public Cloud","Private Cloud"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datacenterTypeTypePropEnum = append(datacenterTypeTypePropEnum, v)
	}
}

const (

	// DatacenterTypePublicCloud captures enum value "Public Cloud"
	DatacenterTypePublicCloud string = "Public Cloud"

	// DatacenterTypePrivateCloud captures enum value "Private Cloud"
	DatacenterTypePrivateCloud string = "Private Cloud"
)

// prop value enum
func (m *Datacenter) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datacenterTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Datacenter) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datacenter based on the context it is used
func (m *Datacenter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Datacenter) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Datacenter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Datacenter) UnmarshalBinary(b []byte) error {
	var res Datacenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}