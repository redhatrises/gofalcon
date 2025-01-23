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

// ModelsRunTool models run tool
//
// swagger:model models.RunTool
type ModelsRunTool struct {

	// driver
	// Required: true
	Driver *ModelsRunToolDriver `json:"driver"`
}

// Validate validates this models run tool
func (m *ModelsRunTool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsRunTool) validateDriver(formats strfmt.Registry) error {

	if err := validate.Required("driver", "body", m.Driver); err != nil {
		return err
	}

	if m.Driver != nil {
		if err := m.Driver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("driver")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this models run tool based on the context it is used
func (m *ModelsRunTool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDriver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsRunTool) contextValidateDriver(ctx context.Context, formats strfmt.Registry) error {

	if m.Driver != nil {

		if err := m.Driver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("driver")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsRunTool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsRunTool) UnmarshalBinary(b []byte) error {
	var res ModelsRunTool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}