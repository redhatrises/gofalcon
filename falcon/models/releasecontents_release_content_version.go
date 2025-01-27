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

// ReleasecontentsReleaseContentVersion releasecontents release content version
//
// swagger:model releasecontents.ReleaseContentVersion
type ReleasecontentsReleaseContentVersion struct {

	// source id
	// Required: true
	SourceID *string `json:"source_id"`

	// source type
	// Required: true
	SourceType *string `json:"source_type"`

	// version
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this releasecontents release content version
func (m *ReleasecontentsReleaseContentVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleasecontentsReleaseContentVersion) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("source_id", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *ReleasecontentsReleaseContentVersion) validateSourceType(formats strfmt.Registry) error {

	if err := validate.Required("source_type", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

func (m *ReleasecontentsReleaseContentVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this releasecontents release content version based on context it is used
func (m *ReleasecontentsReleaseContentVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReleasecontentsReleaseContentVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleasecontentsReleaseContentVersion) UnmarshalBinary(b []byte) error {
	var res ReleasecontentsReleaseContentVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
