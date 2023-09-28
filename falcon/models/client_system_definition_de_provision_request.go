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

// ClientSystemDefinitionDeProvisionRequest client system definition de provision request
//
// swagger:model client.SystemDefinitionDeProvisionRequest
type ClientSystemDefinitionDeProvisionRequest struct {

	// Customer scoped definition ID that is being deprovisioned. This is required when the template is defined as multi-instance
	// Required: true
	DefinitionID *string `json:"definition_id"`

	// When enabled, the CustomerDefinitionID property is ignored and all template workflows are deprovisioned
	// Required: true
	DeprovisionAll *bool `json:"deprovision_all"`

	// ID of the system definition template that is to be deprovisioned
	// Required: true
	TemplateID *string `json:"template_id"`

	// Name of the system definition template to deprovision
	// Required: true
	TemplateName *string `json:"template_name"`
}

// Validate validates this client system definition de provision request
func (m *ClientSystemDefinitionDeProvisionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefinitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeprovisionAll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientSystemDefinitionDeProvisionRequest) validateDefinitionID(formats strfmt.Registry) error {

	if err := validate.Required("definition_id", "body", m.DefinitionID); err != nil {
		return err
	}

	return nil
}

func (m *ClientSystemDefinitionDeProvisionRequest) validateDeprovisionAll(formats strfmt.Registry) error {

	if err := validate.Required("deprovision_all", "body", m.DeprovisionAll); err != nil {
		return err
	}

	return nil
}

func (m *ClientSystemDefinitionDeProvisionRequest) validateTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("template_id", "body", m.TemplateID); err != nil {
		return err
	}

	return nil
}

func (m *ClientSystemDefinitionDeProvisionRequest) validateTemplateName(formats strfmt.Registry) error {

	if err := validate.Required("template_name", "body", m.TemplateName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this client system definition de provision request based on context it is used
func (m *ClientSystemDefinitionDeProvisionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientSystemDefinitionDeProvisionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientSystemDefinitionDeProvisionRequest) UnmarshalBinary(b []byte) error {
	var res ClientSystemDefinitionDeProvisionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
