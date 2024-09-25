// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesSnykMetadata types snyk metadata
//
// swagger:model types.SnykMetadata
type TypesSnykMetadata struct {

	// api endpoint Url
	APIEndpointURL string `json:"apiEndpointUrl,omitempty"`

	// app endpoint Url
	AppEndpointURL string `json:"appEndpointUrl,omitempty"`

	// group Id
	GroupID string `json:"groupId,omitempty"`
}

// Validate validates this types snyk metadata
func (m *TypesSnykMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this types snyk metadata based on context it is used
func (m *TypesSnykMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesSnykMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesSnykMetadata) UnmarshalBinary(b []byte) error {
	var res TypesSnykMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}