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

// FwmgrFirewallAddressRange fwmgr firewall address range
//
// swagger:model fwmgr.firewall.AddressRange
type FwmgrFirewallAddressRange struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// netmask
	Netmask int64 `json:"netmask,omitempty"`
}

// Validate validates this fwmgr firewall address range
func (m *FwmgrFirewallAddressRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrFirewallAddressRange) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fwmgr firewall address range based on context it is used
func (m *FwmgrFirewallAddressRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrFirewallAddressRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrFirewallAddressRange) UnmarshalBinary(b []byte) error {
	var res FwmgrFirewallAddressRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}