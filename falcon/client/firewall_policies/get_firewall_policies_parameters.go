// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetFirewallPoliciesParams creates a new GetFirewallPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFirewallPoliciesParams() *GetFirewallPoliciesParams {
	return &GetFirewallPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFirewallPoliciesParamsWithTimeout creates a new GetFirewallPoliciesParams object
// with the ability to set a timeout on a request.
func NewGetFirewallPoliciesParamsWithTimeout(timeout time.Duration) *GetFirewallPoliciesParams {
	return &GetFirewallPoliciesParams{
		timeout: timeout,
	}
}

// NewGetFirewallPoliciesParamsWithContext creates a new GetFirewallPoliciesParams object
// with the ability to set a context for a request.
func NewGetFirewallPoliciesParamsWithContext(ctx context.Context) *GetFirewallPoliciesParams {
	return &GetFirewallPoliciesParams{
		Context: ctx,
	}
}

// NewGetFirewallPoliciesParamsWithHTTPClient creates a new GetFirewallPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFirewallPoliciesParamsWithHTTPClient(client *http.Client) *GetFirewallPoliciesParams {
	return &GetFirewallPoliciesParams{
		HTTPClient: client,
	}
}

/* GetFirewallPoliciesParams contains all the parameters to send to the API endpoint
   for the get firewall policies operation.

   Typically these are written to a http.Request.
*/
type GetFirewallPoliciesParams struct {

	/* Ids.

	   The IDs of the Firewall Policies to return
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get firewall policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFirewallPoliciesParams) WithDefaults() *GetFirewallPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get firewall policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFirewallPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get firewall policies params
func (o *GetFirewallPoliciesParams) WithTimeout(timeout time.Duration) *GetFirewallPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get firewall policies params
func (o *GetFirewallPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get firewall policies params
func (o *GetFirewallPoliciesParams) WithContext(ctx context.Context) *GetFirewallPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get firewall policies params
func (o *GetFirewallPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get firewall policies params
func (o *GetFirewallPoliciesParams) WithHTTPClient(client *http.Client) *GetFirewallPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get firewall policies params
func (o *GetFirewallPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get firewall policies params
func (o *GetFirewallPoliciesParams) WithIds(ids []string) *GetFirewallPoliciesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get firewall policies params
func (o *GetFirewallPoliciesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetFirewallPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetFirewallPolicies binds the parameter ids
func (o *GetFirewallPoliciesParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}