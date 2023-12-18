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

// ModelsAPIImageAssessment models API image assessment
//
// swagger:model models.APIImageAssessment
type ModelsAPIImageAssessment struct {

	// assessed
	// Required: true
	Assessed *int64 `json:"assessed"`

	// error image pull
	// Required: true
	ErrorImagePull *int64 `json:"error_image_pull"`

	// error image push
	// Required: true
	ErrorImagePush *int64 `json:"error_image_push"`

	// error missing config
	// Required: true
	ErrorMissingConfig *int64 `json:"error_missing_config"`

	// error unsupported schema version
	// Required: true
	ErrorUnsupportedSchemaVersion *int64 `json:"error_unsupported_schema_version"`

	// time stamp
	// Required: true
	TimeStamp *string `json:"time_stamp"`
}

// Validate validates this models API image assessment
func (m *ModelsAPIImageAssessment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssessed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorImagePull(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorImagePush(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorMissingConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorUnsupportedSchemaVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeStamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsAPIImageAssessment) validateAssessed(formats strfmt.Registry) error {

	if err := validate.Required("assessed", "body", m.Assessed); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIImageAssessment) validateErrorImagePull(formats strfmt.Registry) error {

	if err := validate.Required("error_image_pull", "body", m.ErrorImagePull); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIImageAssessment) validateErrorImagePush(formats strfmt.Registry) error {

	if err := validate.Required("error_image_push", "body", m.ErrorImagePush); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIImageAssessment) validateErrorMissingConfig(formats strfmt.Registry) error {

	if err := validate.Required("error_missing_config", "body", m.ErrorMissingConfig); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIImageAssessment) validateErrorUnsupportedSchemaVersion(formats strfmt.Registry) error {

	if err := validate.Required("error_unsupported_schema_version", "body", m.ErrorUnsupportedSchemaVersion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAPIImageAssessment) validateTimeStamp(formats strfmt.Registry) error {

	if err := validate.Required("time_stamp", "body", m.TimeStamp); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models API image assessment based on context it is used
func (m *ModelsAPIImageAssessment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAPIImageAssessment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAPIImageAssessment) UnmarshalBinary(b []byte) error {
	var res ModelsAPIImageAssessment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}