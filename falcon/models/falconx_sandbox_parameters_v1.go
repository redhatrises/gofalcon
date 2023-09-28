// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxSandboxParametersV1 falconx sandbox parameters v1
//
// swagger:model falconx.SandboxParametersV1
type FalconxSandboxParametersV1 struct {

	// action script
	ActionScript string `json:"action_script,omitempty"`

	// command line
	CommandLine string `json:"command_line,omitempty"`

	// document password
	DocumentPassword string `json:"document_password,omitempty"`

	// enable tor
	EnableTor bool `json:"enable_tor,omitempty"`

	// environment id
	EnvironmentID int32 `json:"environment_id,omitempty"`

	// interactivity
	Interactivity bool `json:"interactivity,omitempty"`

	// network settings
	NetworkSettings string `json:"network_settings,omitempty"`

	// sha256
	Sha256 string `json:"sha256,omitempty"`

	// submit name
	SubmitName string `json:"submit_name,omitempty"`

	// system date
	SystemDate string `json:"system_date,omitempty"`

	// system time
	SystemTime string `json:"system_time,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this falconx sandbox parameters v1
func (m *FalconxSandboxParametersV1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this falconx sandbox parameters v1 based on context it is used
func (m *FalconxSandboxParametersV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FalconxSandboxParametersV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxSandboxParametersV1) UnmarshalBinary(b []byte) error {
	var res FalconxSandboxParametersV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
