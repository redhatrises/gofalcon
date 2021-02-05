// Code generated by go-swagger; DO NOT EDIT.

package installation_tokens

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
)

// NewCustomerSettingsReadParams creates a new CustomerSettingsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerSettingsReadParams() *CustomerSettingsReadParams {
	return &CustomerSettingsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerSettingsReadParamsWithTimeout creates a new CustomerSettingsReadParams object
// with the ability to set a timeout on a request.
func NewCustomerSettingsReadParamsWithTimeout(timeout time.Duration) *CustomerSettingsReadParams {
	return &CustomerSettingsReadParams{
		timeout: timeout,
	}
}

// NewCustomerSettingsReadParamsWithContext creates a new CustomerSettingsReadParams object
// with the ability to set a context for a request.
func NewCustomerSettingsReadParamsWithContext(ctx context.Context) *CustomerSettingsReadParams {
	return &CustomerSettingsReadParams{
		Context: ctx,
	}
}

// NewCustomerSettingsReadParamsWithHTTPClient creates a new CustomerSettingsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerSettingsReadParamsWithHTTPClient(client *http.Client) *CustomerSettingsReadParams {
	return &CustomerSettingsReadParams{
		HTTPClient: client,
	}
}

/* CustomerSettingsReadParams contains all the parameters to send to the API endpoint
   for the customer settings read operation.

   Typically these are written to a http.Request.
*/
type CustomerSettingsReadParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer settings read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerSettingsReadParams) WithDefaults() *CustomerSettingsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer settings read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerSettingsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer settings read params
func (o *CustomerSettingsReadParams) WithTimeout(timeout time.Duration) *CustomerSettingsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer settings read params
func (o *CustomerSettingsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer settings read params
func (o *CustomerSettingsReadParams) WithContext(ctx context.Context) *CustomerSettingsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer settings read params
func (o *CustomerSettingsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer settings read params
func (o *CustomerSettingsReadParams) WithHTTPClient(client *http.Client) *CustomerSettingsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer settings read params
func (o *CustomerSettingsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerSettingsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}