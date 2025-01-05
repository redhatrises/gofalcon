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

// DomainControl domain control
//
// swagger:model domain.Control
type DomainControl struct {

	// benchmarks
	Benchmarks []*DomainControlBenchmark `json:"benchmarks"`

	// description
	Description string `json:"description,omitempty"`

	// framework
	Framework string `json:"framework,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// section
	Section string `json:"section,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this domain control
func (m *DomainControl) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBenchmarks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainControl) validateBenchmarks(formats strfmt.Registry) error {
	if swag.IsZero(m.Benchmarks) { // not required
		return nil
	}

	for i := 0; i < len(m.Benchmarks); i++ {
		if swag.IsZero(m.Benchmarks[i]) { // not required
			continue
		}

		if m.Benchmarks[i] != nil {
			if err := m.Benchmarks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("benchmarks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("benchmarks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this domain control based on the context it is used
func (m *DomainControl) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBenchmarks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainControl) contextValidateBenchmarks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Benchmarks); i++ {

		if m.Benchmarks[i] != nil {

			if swag.IsZero(m.Benchmarks[i]) { // not required
				return nil
			}

			if err := m.Benchmarks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("benchmarks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("benchmarks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainControl) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainControl) UnmarshalBinary(b []byte) error {
	var res DomainControl
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}