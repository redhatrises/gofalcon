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

// ModelsResultLocation models result location
//
// swagger:model models.ResultLocation
type ModelsResultLocation struct {

	// physical location
	// Required: true
	PhysicalLocation *ModelsResultPhysicalLocation `json:"physicalLocation"`

	// properties
	// Required: true
	Properties *ModelsResultLocationProperties `json:"properties"`
}

// Validate validates this models result location
func (m *ModelsResultLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhysicalLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsResultLocation) validatePhysicalLocation(formats strfmt.Registry) error {

	if err := validate.Required("physicalLocation", "body", m.PhysicalLocation); err != nil {
		return err
	}

	if m.PhysicalLocation != nil {
		if err := m.PhysicalLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("physicalLocation")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsResultLocation) validateProperties(formats strfmt.Registry) error {

	if err := validate.Required("properties", "body", m.Properties); err != nil {
		return err
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this models result location based on the context it is used
func (m *ModelsResultLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePhysicalLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsResultLocation) contextValidatePhysicalLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.PhysicalLocation != nil {

		if err := m.PhysicalLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("physicalLocation")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsResultLocation) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	if m.Properties != nil {

		if err := m.Properties.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsResultLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsResultLocation) UnmarshalBinary(b []byte) error {
	var res ModelsResultLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}