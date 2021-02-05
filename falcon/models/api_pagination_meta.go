// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPaginationMeta api pagination meta
//
// swagger:model api.paginationMeta
type APIPaginationMeta struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// next page
	NextPage string `json:"next_page,omitempty"`

	// offset
	Offset APIPaginationMetaOffset `json:"offset,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this api pagination meta
func (m *APIPaginationMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api pagination meta based on context it is used
func (m *APIPaginationMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPaginationMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPaginationMeta) UnmarshalBinary(b []byte) error {
	var res APIPaginationMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}