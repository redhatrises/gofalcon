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

// DomainCloudCondition domain cloud condition
//
// swagger:model domain.CloudCondition
type DomainCloudCondition struct {

	// feature
	// Required: true
	Feature *string `json:"feature"`

	// is visible
	// Required: true
	IsVisible *bool `json:"is_visible"`

	// last transition
	// Required: true
	// Format: date-time
	LastTransition *strfmt.DateTime `json:"last_transition"`

	// message
	Message string `json:"message,omitempty"`

	// product
	Product string `json:"product,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this domain cloud condition
func (m *DomainCloudCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsVisible(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTransition(formats); err != nil {
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

func (m *DomainCloudCondition) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *DomainCloudCondition) validateIsVisible(formats strfmt.Registry) error {

	if err := validate.Required("is_visible", "body", m.IsVisible); err != nil {
		return err
	}

	return nil
}

func (m *DomainCloudCondition) validateLastTransition(formats strfmt.Registry) error {

	if err := validate.Required("last_transition", "body", m.LastTransition); err != nil {
		return err
	}

	if err := validate.FormatOf("last_transition", "body", "date-time", m.LastTransition.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainCloudCondition) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *DomainCloudCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain cloud condition based on context it is used
func (m *DomainCloudCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainCloudCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCloudCondition) UnmarshalBinary(b []byte) error {
	var res DomainCloudCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}