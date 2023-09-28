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

// DomainScan domain scan
//
// swagger:model domain.Scan
type DomainScan struct {

	// affected hosts count
	AffectedHostsCount int32 `json:"affected_hosts_count,omitempty"`

	// cid
	Cid string `json:"cid,omitempty"`

	// cloud ml level detection
	CloudMlLevelDetection int32 `json:"cloud_ml_level_detection,omitempty"`

	// cloud ml level prevention
	CloudMlLevelPrevention int32 `json:"cloud_ml_level_prevention,omitempty"`

	// completed host count
	CompletedHostCount int32 `json:"completed_host_count,omitempty"`

	// cpu priority
	CPUPriority int32 `json:"cpu_priority,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// endpoint notification
	EndpointNotification bool `json:"endpoint_notification,omitempty"`

	// file paths
	FilePaths []string `json:"file_paths"`

	// filecount
	Filecount *DomainFileCount `json:"filecount,omitempty"`

	// host groups
	HostGroups []string `json:"host_groups"`

	// hosts
	Hosts []string `json:"hosts"`

	// id
	// Required: true
	ID *string `json:"id"`

	// incomplete host count
	IncompleteHostCount int32 `json:"incomplete_host_count,omitempty"`

	// initiated from
	InitiatedFrom string `json:"initiated_from,omitempty"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// max duration
	MaxDuration int32 `json:"max_duration,omitempty"`

	// max file size
	MaxFileSize int32 `json:"max_file_size,omitempty"`

	// metadata
	Metadata []*DomainScanMetadata `json:"metadata"`

	// missing host count
	MissingHostCount int32 `json:"missing_host_count,omitempty"`

	// not started host count
	NotStartedHostCount int32 `json:"not_started_host_count,omitempty"`

	// pause duration
	PauseDuration int32 `json:"pause_duration,omitempty"`

	// policy setting
	PolicySetting []int64 `json:"policy_setting"`

	// preemption priority
	PreemptionPriority int32 `json:"preemption_priority,omitempty"`

	// profile id
	ProfileID string `json:"profile_id,omitempty"`

	// quarantine
	Quarantine bool `json:"quarantine,omitempty"`

	// scan completed on
	// Format: date-time
	ScanCompletedOn strfmt.DateTime `json:"scan_completed_on,omitempty"`

	// scan exclusions
	ScanExclusions []string `json:"scan_exclusions"`

	// scan inclusions
	ScanInclusions []string `json:"scan_inclusions"`

	// scan scheduled on
	// Format: date-time
	ScanScheduledOn strfmt.DateTime `json:"scan_scheduled_on,omitempty"`

	// scan started on
	// Format: date-time
	ScanStartedOn strfmt.DateTime `json:"scan_started_on,omitempty"`

	// sensor ml level detection
	SensorMlLevelDetection int32 `json:"sensor_ml_level_detection,omitempty"`

	// sensor ml level prevention
	SensorMlLevelPrevention int32 `json:"sensor_ml_level_prevention,omitempty"`

	// severity
	Severity int64 `json:"severity,omitempty"`

	// started host count
	StartedHostCount int32 `json:"started_host_count,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// targeted host count
	TargetedHostCount int32 `json:"targeted_host_count,omitempty"`
}

// Validate validates this domain scan
func (m *DomainScan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilecount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanCompletedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanScheduledOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanStartedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainScan) validateCreatedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScan) validateFilecount(formats strfmt.Registry) error {
	if swag.IsZero(m.Filecount) { // not required
		return nil
	}

	if m.Filecount != nil {
		if err := m.Filecount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filecount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filecount")
			}
			return err
		}
	}

	return nil
}

func (m *DomainScan) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainScan) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScan) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainScan) validateScanCompletedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.ScanCompletedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("scan_completed_on", "body", "date-time", m.ScanCompletedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScan) validateScanScheduledOn(formats strfmt.Registry) error {
	if swag.IsZero(m.ScanScheduledOn) { // not required
		return nil
	}

	if err := validate.FormatOf("scan_scheduled_on", "body", "date-time", m.ScanScheduledOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScan) validateScanStartedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.ScanStartedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("scan_started_on", "body", "date-time", m.ScanStartedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain scan based on the context it is used
func (m *DomainScan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilecount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainScan) contextValidateFilecount(ctx context.Context, formats strfmt.Registry) error {

	if m.Filecount != nil {

		if swag.IsZero(m.Filecount) { // not required
			return nil
		}

		if err := m.Filecount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filecount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filecount")
			}
			return err
		}
	}

	return nil
}

func (m *DomainScan) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metadata); i++ {

		if m.Metadata[i] != nil {

			if swag.IsZero(m.Metadata[i]) { // not required
				return nil
			}

			if err := m.Metadata[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainScan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainScan) UnmarshalBinary(b []byte) error {
	var res DomainScan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
