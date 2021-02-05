// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainUpdateUserFields domain update user fields
//
// swagger:model domain.UpdateUserFields
type DomainUpdateUserFields struct {

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`
}

// Validate validates this domain update user fields
func (m *DomainUpdateUserFields) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain update user fields based on context it is used
func (m *DomainUpdateUserFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainUpdateUserFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainUpdateUserFields) UnmarshalBinary(b []byte) error {
	var res DomainUpdateUserFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}