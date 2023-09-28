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

// RulegroupsRule rulegroups rule
//
// swagger:model rulegroups.Rule
type RulegroupsRule struct {

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// depth
	// Required: true
	Depth *string `json:"depth"`

	// description
	Description string `json:"description,omitempty"`

	// exclude
	Exclude string `json:"exclude,omitempty"`

	// exclude processes
	ExcludeProcesses string `json:"exclude_processes,omitempty"`

	// exclude users
	ExcludeUsers string `json:"exclude_users,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// include
	// Required: true
	Include *string `json:"include"`

	// include processes
	IncludeProcesses string `json:"include_processes,omitempty"`

	// include users
	IncludeUsers string `json:"include_users,omitempty"`

	// modified timestamp
	ModifiedTimestamp string `json:"modified_timestamp,omitempty"`

	// path
	// Required: true
	Path *string `json:"path"`

	// precedence
	Precedence int32 `json:"precedence,omitempty"`

	// rule group id
	// Required: true
	RuleGroupID *string `json:"rule_group_id"`

	// severity
	// Required: true
	Severity *string `json:"severity"`

	// type
	// Required: true
	Type *string `json:"type"`

	// watch attributes directory changes
	WatchAttributesDirectoryChanges bool `json:"watch_attributes_directory_changes,omitempty"`

	// watch attributes file changes
	WatchAttributesFileChanges bool `json:"watch_attributes_file_changes,omitempty"`

	// watch create directory changes
	WatchCreateDirectoryChanges bool `json:"watch_create_directory_changes,omitempty"`

	// watch create file changes
	WatchCreateFileChanges bool `json:"watch_create_file_changes,omitempty"`

	// watch create key changes
	WatchCreateKeyChanges bool `json:"watch_create_key_changes,omitempty"`

	// watch delete directory changes
	WatchDeleteDirectoryChanges bool `json:"watch_delete_directory_changes,omitempty"`

	// watch delete file changes
	WatchDeleteFileChanges bool `json:"watch_delete_file_changes,omitempty"`

	// watch delete key changes
	WatchDeleteKeyChanges bool `json:"watch_delete_key_changes,omitempty"`

	// watch delete value changes
	WatchDeleteValueChanges bool `json:"watch_delete_value_changes,omitempty"`

	// watch permissions directory changes
	WatchPermissionsDirectoryChanges bool `json:"watch_permissions_directory_changes,omitempty"`

	// watch permissions file changes
	WatchPermissionsFileChanges bool `json:"watch_permissions_file_changes,omitempty"`

	// watch rename directory changes
	WatchRenameDirectoryChanges bool `json:"watch_rename_directory_changes,omitempty"`

	// watch rename file changes
	WatchRenameFileChanges bool `json:"watch_rename_file_changes,omitempty"`

	// watch rename key changes
	WatchRenameKeyChanges bool `json:"watch_rename_key_changes,omitempty"`

	// watch set value changes
	WatchSetValueChanges bool `json:"watch_set_value_changes,omitempty"`

	// watch write file changes
	WatchWriteFileChanges bool `json:"watch_write_file_changes,omitempty"`
}

// Validate validates this rulegroups rule
func (m *RulegroupsRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInclude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
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

func (m *RulegroupsRule) validateDepth(formats strfmt.Registry) error {

	if err := validate.Required("depth", "body", m.Depth); err != nil {
		return err
	}

	return nil
}

func (m *RulegroupsRule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RulegroupsRule) validateInclude(formats strfmt.Registry) error {

	if err := validate.Required("include", "body", m.Include); err != nil {
		return err
	}

	return nil
}

func (m *RulegroupsRule) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *RulegroupsRule) validateRuleGroupID(formats strfmt.Registry) error {

	if err := validate.Required("rule_group_id", "body", m.RuleGroupID); err != nil {
		return err
	}

	return nil
}

func (m *RulegroupsRule) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *RulegroupsRule) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rulegroups rule based on context it is used
func (m *RulegroupsRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RulegroupsRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RulegroupsRule) UnmarshalBinary(b []byte) error {
	var res RulegroupsRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
