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
	"github.com/go-openapi/validate"
)

// MlscannerSamplesScanResult mlscanner samples scan result
//
// swagger:model mlscanner.SamplesScanResult
type MlscannerSamplesScanResult struct {

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// samples
	// Required: true
	Samples []*MlscannerScannedSample `json:"samples"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this mlscanner samples scan result
func (m *MlscannerSamplesScanResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamples(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MlscannerSamplesScanResult) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *MlscannerSamplesScanResult) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *MlscannerSamplesScanResult) validateSamples(formats strfmt.Registry) error {

	if err := validate.Required("samples", "body", m.Samples); err != nil {
		return err
	}

	for i := 0; i < len(m.Samples); i++ {
		if swag.IsZero(m.Samples[i]) { // not required
			continue
		}

		if m.Samples[i] != nil {
			if err := m.Samples[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("samples" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MlscannerSamplesScanResult) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mlscanner samples scan result based on the context it is used
func (m *MlscannerSamplesScanResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSamples(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MlscannerSamplesScanResult) contextValidateSamples(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Samples); i++ {

		if m.Samples[i] != nil {
			if err := m.Samples[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("samples" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MlscannerSamplesScanResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MlscannerSamplesScanResult) UnmarshalBinary(b []byte) error {
	var res MlscannerSamplesScanResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}