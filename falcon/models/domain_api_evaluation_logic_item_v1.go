// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	jsonext "encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainAPIEvaluationLogicItemV1 domain API evaluation logic item v1
//
// swagger:model domain.APIEvaluationLogicItemV1
type DomainAPIEvaluationLogicItemV1 struct {

	// comparison check
	ComparisonCheck string `json:"comparison_check,omitempty"`

	// comparisons
	Comparisons *DomainAPIEvaluationLogicComparisonsV1 `json:"comparisons,omitempty"`

	// determined by comparison
	DeterminedByComparison bool `json:"determined_by_comparison,omitempty"`

	// existence check
	ExistenceCheck string `json:"existence_check,omitempty"`

	// id
	ID jsonext.Number `json:"id,omitempty"`

	// items
	Items []DomainAPIEvaluationLogicSystemCharacteristicV1 `json:"items"`

	// negate
	Negate bool `json:"negate,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this domain API evaluation logic item v1
func (m *DomainAPIEvaluationLogicItemV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparisons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
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

func (m *DomainAPIEvaluationLogicItemV1) validateComparisons(formats strfmt.Registry) error {
	if swag.IsZero(m.Comparisons) { // not required
		return nil
	}

	if m.Comparisons != nil {
		if err := m.Comparisons.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("comparisons")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("comparisons")
			}
			return err
		}
	}

	return nil
}

func (m *DomainAPIEvaluationLogicItemV1) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIEvaluationLogicItemV1) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain API evaluation logic item v1 based on the context it is used
func (m *DomainAPIEvaluationLogicItemV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComparisons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIEvaluationLogicItemV1) contextValidateComparisons(ctx context.Context, formats strfmt.Registry) error {

	if m.Comparisons != nil {
		if err := m.Comparisons.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("comparisons")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("comparisons")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIEvaluationLogicItemV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIEvaluationLogicItemV1) UnmarshalBinary(b []byte) error {
	var res DomainAPIEvaluationLogicItemV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}