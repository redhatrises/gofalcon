// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewSetSensorUpdatePoliciesPrecedenceParams creates a new SetSensorUpdatePoliciesPrecedenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetSensorUpdatePoliciesPrecedenceParams() *SetSensorUpdatePoliciesPrecedenceParams {
	return &SetSensorUpdatePoliciesPrecedenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetSensorUpdatePoliciesPrecedenceParamsWithTimeout creates a new SetSensorUpdatePoliciesPrecedenceParams object
// with the ability to set a timeout on a request.
func NewSetSensorUpdatePoliciesPrecedenceParamsWithTimeout(timeout time.Duration) *SetSensorUpdatePoliciesPrecedenceParams {
	return &SetSensorUpdatePoliciesPrecedenceParams{
		timeout: timeout,
	}
}

// NewSetSensorUpdatePoliciesPrecedenceParamsWithContext creates a new SetSensorUpdatePoliciesPrecedenceParams object
// with the ability to set a context for a request.
func NewSetSensorUpdatePoliciesPrecedenceParamsWithContext(ctx context.Context) *SetSensorUpdatePoliciesPrecedenceParams {
	return &SetSensorUpdatePoliciesPrecedenceParams{
		Context: ctx,
	}
}

// NewSetSensorUpdatePoliciesPrecedenceParamsWithHTTPClient creates a new SetSensorUpdatePoliciesPrecedenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetSensorUpdatePoliciesPrecedenceParamsWithHTTPClient(client *http.Client) *SetSensorUpdatePoliciesPrecedenceParams {
	return &SetSensorUpdatePoliciesPrecedenceParams{
		HTTPClient: client,
	}
}

/* SetSensorUpdatePoliciesPrecedenceParams contains all the parameters to send to the API endpoint
   for the set sensor update policies precedence operation.

   Typically these are written to a http.Request.
*/
type SetSensorUpdatePoliciesPrecedenceParams struct {

	// Body.
	Body *models.RequestsSetPolicyPrecedenceReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set sensor update policies precedence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetSensorUpdatePoliciesPrecedenceParams) WithDefaults() *SetSensorUpdatePoliciesPrecedenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set sensor update policies precedence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetSensorUpdatePoliciesPrecedenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set sensor update policies precedence params
func (o *SetSensorUpdatePoliciesPrecedenceParams) WithTimeout(timeout time.Duration) *SetSensorUpdatePoliciesPrecedenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set sensor update policies precedence params
func (o *SetSensorUpdatePoliciesPrecedenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set sensor update policies precedence params
func (o *SetSensorUpdatePoliciesPrecedenceParams) WithContext(ctx context.Context) *SetSensorUpdatePoliciesPrecedenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set sensor update policies precedence params
func (o *SetSensorUpdatePoliciesPrecedenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set sensor update policies precedence params
func (o *SetSensorUpdatePoliciesPrecedenceParams) WithHTTPClient(client *http.Client) *SetSensorUpdatePoliciesPrecedenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set sensor update policies precedence params
func (o *SetSensorUpdatePoliciesPrecedenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set sensor update policies precedence params
func (o *SetSensorUpdatePoliciesPrecedenceParams) WithBody(body *models.RequestsSetPolicyPrecedenceReqV1) *SetSensorUpdatePoliciesPrecedenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set sensor update policies precedence params
func (o *SetSensorUpdatePoliciesPrecedenceParams) SetBody(body *models.RequestsSetPolicyPrecedenceReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetSensorUpdatePoliciesPrecedenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}