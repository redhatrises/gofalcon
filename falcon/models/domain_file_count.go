// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainFileCount domain file count
//
// swagger:model domain.FileCount
type DomainFileCount struct {

	// malicious
	Malicious int32 `json:"malicious,omitempty"`

	// quarantined
	Quarantined int32 `json:"quarantined,omitempty"`

	// scanned
	Scanned int32 `json:"scanned,omitempty"`

	// skipped
	Skipped int32 `json:"skipped,omitempty"`
}

// Validate validates this domain file count
func (m *DomainFileCount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain file count based on context it is used
func (m *DomainFileCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainFileCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainFileCount) UnmarshalBinary(b []byte) error {
	var res DomainFileCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
