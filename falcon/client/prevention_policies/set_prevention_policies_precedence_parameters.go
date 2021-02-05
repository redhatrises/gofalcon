// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// NewSetPreventionPoliciesPrecedenceParams creates a new SetPreventionPoliciesPrecedenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetPreventionPoliciesPrecedenceParams() *SetPreventionPoliciesPrecedenceParams {
	return &SetPreventionPoliciesPrecedenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetPreventionPoliciesPrecedenceParamsWithTimeout creates a new SetPreventionPoliciesPrecedenceParams object
// with the ability to set a timeout on a request.
func NewSetPreventionPoliciesPrecedenceParamsWithTimeout(timeout time.Duration) *SetPreventionPoliciesPrecedenceParams {
	return &SetPreventionPoliciesPrecedenceParams{
		timeout: timeout,
	}
}

// NewSetPreventionPoliciesPrecedenceParamsWithContext creates a new SetPreventionPoliciesPrecedenceParams object
// with the ability to set a context for a request.
func NewSetPreventionPoliciesPrecedenceParamsWithContext(ctx context.Context) *SetPreventionPoliciesPrecedenceParams {
	return &SetPreventionPoliciesPrecedenceParams{
		Context: ctx,
	}
}

// NewSetPreventionPoliciesPrecedenceParamsWithHTTPClient creates a new SetPreventionPoliciesPrecedenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetPreventionPoliciesPrecedenceParamsWithHTTPClient(client *http.Client) *SetPreventionPoliciesPrecedenceParams {
	return &SetPreventionPoliciesPrecedenceParams{
		HTTPClient: client,
	}
}

/* SetPreventionPoliciesPrecedenceParams contains all the parameters to send to the API endpoint
   for the set prevention policies precedence operation.

   Typically these are written to a http.Request.
*/
type SetPreventionPoliciesPrecedenceParams struct {

	// Body.
	Body *models.RequestsSetPolicyPrecedenceReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set prevention policies precedence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPreventionPoliciesPrecedenceParams) WithDefaults() *SetPreventionPoliciesPrecedenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set prevention policies precedence params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPreventionPoliciesPrecedenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set prevention policies precedence params
func (o *SetPreventionPoliciesPrecedenceParams) WithTimeout(timeout time.Duration) *SetPreventionPoliciesPrecedenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set prevention policies precedence params
func (o *SetPreventionPoliciesPrecedenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set prevention policies precedence params
func (o *SetPreventionPoliciesPrecedenceParams) WithContext(ctx context.Context) *SetPreventionPoliciesPrecedenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set prevention policies precedence params
func (o *SetPreventionPoliciesPrecedenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set prevention policies precedence params
func (o *SetPreventionPoliciesPrecedenceParams) WithHTTPClient(client *http.Client) *SetPreventionPoliciesPrecedenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set prevention policies precedence params
func (o *SetPreventionPoliciesPrecedenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set prevention policies precedence params
func (o *SetPreventionPoliciesPrecedenceParams) WithBody(body *models.RequestsSetPolicyPrecedenceReqV1) *SetPreventionPoliciesPrecedenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set prevention policies precedence params
func (o *SetPreventionPoliciesPrecedenceParams) SetBody(body *models.RequestsSetPolicyPrecedenceReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetPreventionPoliciesPrecedenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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