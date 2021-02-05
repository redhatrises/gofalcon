// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxContactedHost falconx contacted host
//
// swagger:model falconx.ContactedHost
type FalconxContactedHost struct {

	// address
	Address string `json:"address,omitempty"`

	// associated runtime
	AssociatedRuntime []*FalconxAssociatedRuntime `json:"associated_runtime"`

	// compromised
	Compromised bool `json:"compromised,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// port
	Port int32 `json:"port,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this falconx contacted host
func (m *FalconxContactedHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociatedRuntime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxContactedHost) validateAssociatedRuntime(formats strfmt.Registry) error {
	if swag.IsZero(m.AssociatedRuntime) { // not required
		return nil
	}

	for i := 0; i < len(m.AssociatedRuntime); i++ {
		if swag.IsZero(m.AssociatedRuntime[i]) { // not required
			continue
		}

		if m.AssociatedRuntime[i] != nil {
			if err := m.AssociatedRuntime[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associated_runtime" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this falconx contacted host based on the context it is used
func (m *FalconxContactedHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssociatedRuntime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxContactedHost) contextValidateAssociatedRuntime(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AssociatedRuntime); i++ {

		if m.AssociatedRuntime[i] != nil {
			if err := m.AssociatedRuntime[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associated_runtime" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxContactedHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxContactedHost) UnmarshalBinary(b []byte) error {
	var res FalconxContactedHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}