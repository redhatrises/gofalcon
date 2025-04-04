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

// V2Model v2 model
//
// swagger:model v2.Model
type V2Model struct {

	// actions
	Actions map[string]V2Activity `json:"actions,omitempty"`

	// conditions
	Conditions map[string]V2Condition `json:"conditions,omitempty"`

	// loops
	Loops map[string]V2Loop `json:"loops,omitempty"`

	// node registry
	// Required: true
	NodeRegistry map[string]string `json:"nodeRegistry"`

	// output fields
	OutputFields []string `json:"output_fields"`

	// parent
	// Required: true
	Parent *V2Model `json:"parent"`

	// summary
	Summary string `json:"summary,omitempty"`

	// trigger
	// Required: true
	Trigger *V2Trigger `json:"trigger"`

	// uniq node seen
	// Required: true
	UniqNodeSeen map[string]bool `json:"uniqNodeSeen"`
}

// Validate validates this v2 model
func (m *V2Model) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniqNodeSeen(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2Model) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for k := range m.Actions {

		if err := validate.Required("actions"+"."+k, "body", m.Actions[k]); err != nil {
			return err
		}
		if val, ok := m.Actions[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2Model) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for k := range m.Conditions {

		if err := validate.Required("conditions"+"."+k, "body", m.Conditions[k]); err != nil {
			return err
		}
		if val, ok := m.Conditions[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2Model) validateLoops(formats strfmt.Registry) error {
	if swag.IsZero(m.Loops) { // not required
		return nil
	}

	for k := range m.Loops {

		if err := validate.Required("loops"+"."+k, "body", m.Loops[k]); err != nil {
			return err
		}
		if val, ok := m.Loops[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("loops" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("loops" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2Model) validateNodeRegistry(formats strfmt.Registry) error {

	if err := validate.Required("nodeRegistry", "body", m.NodeRegistry); err != nil {
		return err
	}

	return nil
}

func (m *V2Model) validateParent(formats strfmt.Registry) error {

	if err := validate.Required("parent", "body", m.Parent); err != nil {
		return err
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *V2Model) validateTrigger(formats strfmt.Registry) error {

	if err := validate.Required("trigger", "body", m.Trigger); err != nil {
		return err
	}

	if m.Trigger != nil {
		if err := m.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

func (m *V2Model) validateUniqNodeSeen(formats strfmt.Registry) error {

	if err := validate.Required("uniqNodeSeen", "body", m.UniqNodeSeen); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v2 model based on the context it is used
func (m *V2Model) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrigger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2Model) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Actions {

		if val, ok := m.Actions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V2Model) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Conditions {

		if val, ok := m.Conditions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V2Model) contextValidateLoops(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Loops {

		if val, ok := m.Loops[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V2Model) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {

		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *V2Model) contextValidateTrigger(ctx context.Context, formats strfmt.Registry) error {

	if m.Trigger != nil {

		if err := m.Trigger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2Model) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2Model) UnmarshalBinary(b []byte) error {
	var res V2Model
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
