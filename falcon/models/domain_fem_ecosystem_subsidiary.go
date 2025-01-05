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

// DomainFemEcosystemSubsidiary domain fem ecosystem subsidiary
//
// swagger:model domain.FemEcosystemSubsidiary
type DomainFemEcosystemSubsidiary struct {

	// The number of assets associated with the subsidiary
	AssetCount int32 `json:"asset_count,omitempty"`

	// The number of subsidiary children
	// Required: true
	ChildrenCount *int32 `json:"children_count"`

	// The customer ID
	// Required: true
	Cid *string `json:"cid"`

	// An object that shows how this subsidiary has been linked to the customer inventory
	Discovery *DomainExternalAssetDiscoveryAttributes `json:"discovery,omitempty"`

	// The ID of the subsidiary
	// Required: true
	ID *string `json:"id"`

	// The name of the subsidiary
	// Required: true
	Name *string `json:"name"`

	// The ID of the parent subsidiary
	// Required: true
	ParentID *string `json:"parent_id"`

	// The primary domain of the subsidiary (Originally was called apex_domain)
	// Required: true
	PrimaryDomain *string `json:"primary_domain"`

	// The risk score of the subsidiary
	RiskScore float64 `json:"risk_score,omitempty"`
}

// Validate validates this domain fem ecosystem subsidiary
func (m *DomainFemEcosystemSubsidiary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildrenCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscovery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainFemEcosystemSubsidiary) validateChildrenCount(formats strfmt.Registry) error {

	if err := validate.Required("children_count", "body", m.ChildrenCount); err != nil {
		return err
	}

	return nil
}

func (m *DomainFemEcosystemSubsidiary) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *DomainFemEcosystemSubsidiary) validateDiscovery(formats strfmt.Registry) error {
	if swag.IsZero(m.Discovery) { // not required
		return nil
	}

	if m.Discovery != nil {
		if err := m.Discovery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discovery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discovery")
			}
			return err
		}
	}

	return nil
}

func (m *DomainFemEcosystemSubsidiary) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainFemEcosystemSubsidiary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainFemEcosystemSubsidiary) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

func (m *DomainFemEcosystemSubsidiary) validatePrimaryDomain(formats strfmt.Registry) error {

	if err := validate.Required("primary_domain", "body", m.PrimaryDomain); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain fem ecosystem subsidiary based on the context it is used
func (m *DomainFemEcosystemSubsidiary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiscovery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainFemEcosystemSubsidiary) contextValidateDiscovery(ctx context.Context, formats strfmt.Registry) error {

	if m.Discovery != nil {

		if swag.IsZero(m.Discovery) { // not required
			return nil
		}

		if err := m.Discovery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discovery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discovery")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainFemEcosystemSubsidiary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainFemEcosystemSubsidiary) UnmarshalBinary(b []byte) error {
	var res DomainFemEcosystemSubsidiary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}