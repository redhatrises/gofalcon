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

// APIReportExecutionV1 api report execution v1
//
// swagger:model api.ReportExecutionV1
type APIReportExecutionV1 struct {

	// can write
	// Required: true
	CanWrite *bool `json:"can_write"`

	// created on
	// Required: true
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on"`

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// expiration on
	// Required: true
	// Format: date-time
	ExpirationOn *strfmt.DateTime `json:"expiration_on"`

	// id
	// Required: true
	ID *string `json:"id"`

	// job reference
	// Required: true
	JobReference *string `json:"job_reference"`

	// last updated on
	// Required: true
	// Format: date-time
	LastUpdatedOn *strfmt.DateTime `json:"last_updated_on"`

	// report file reference
	// Required: true
	ReportFileReference *string `json:"report_file_reference"`

	// result metadata
	ResultMetadata *DomainResultMetadata `json:"result_metadata,omitempty"`

	// scheduled report id
	// Required: true
	ScheduledReportID *string `json:"scheduled_report_id"`

	// shared with
	// Required: true
	SharedWith []string `json:"shared_with"`

	// status
	// Required: true
	Status *string `json:"status"`

	// status msg
	// Required: true
	StatusMsg *string `json:"status_msg"`

	// tracking
	Tracking string `json:"tracking,omitempty"`

	// trigger reference
	// Required: true
	TriggerReference *string `json:"trigger_reference"`

	// type
	// Required: true
	Type *string `json:"type"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`

	// user uuid
	// Required: true
	UserUUID *string `json:"user_uuid"`
}

// Validate validates this api report execution v1
func (m *APIReportExecutionV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanWrite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportFileReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledReportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusMsg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggerReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIReportExecutionV1) validateCanWrite(formats strfmt.Registry) error {

	if err := validate.Required("can_write", "body", m.CanWrite); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateCreatedOn(formats strfmt.Registry) error {

	if err := validate.Required("created_on", "body", m.CreatedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateExpirationOn(formats strfmt.Registry) error {

	if err := validate.Required("expiration_on", "body", m.ExpirationOn); err != nil {
		return err
	}

	if err := validate.FormatOf("expiration_on", "body", "date-time", m.ExpirationOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateJobReference(formats strfmt.Registry) error {

	if err := validate.Required("job_reference", "body", m.JobReference); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateLastUpdatedOn(formats strfmt.Registry) error {

	if err := validate.Required("last_updated_on", "body", m.LastUpdatedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("last_updated_on", "body", "date-time", m.LastUpdatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateReportFileReference(formats strfmt.Registry) error {

	if err := validate.Required("report_file_reference", "body", m.ReportFileReference); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateResultMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.ResultMetadata) { // not required
		return nil
	}

	if m.ResultMetadata != nil {
		if err := m.ResultMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *APIReportExecutionV1) validateScheduledReportID(formats strfmt.Registry) error {

	if err := validate.Required("scheduled_report_id", "body", m.ScheduledReportID); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateSharedWith(formats strfmt.Registry) error {

	if err := validate.Required("shared_with", "body", m.SharedWith); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateStatusMsg(formats strfmt.Registry) error {

	if err := validate.Required("status_msg", "body", m.StatusMsg); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateTriggerReference(formats strfmt.Registry) error {

	if err := validate.Required("trigger_reference", "body", m.TriggerReference); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *APIReportExecutionV1) validateUserUUID(formats strfmt.Registry) error {

	if err := validate.Required("user_uuid", "body", m.UserUUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api report execution v1 based on the context it is used
func (m *APIReportExecutionV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResultMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIReportExecutionV1) contextValidateResultMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.ResultMetadata != nil {
		if err := m.ResultMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result_metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIReportExecutionV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIReportExecutionV1) UnmarshalBinary(b []byte) error {
	var res APIReportExecutionV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}