// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// NewCreateD4CGcpAccountParams creates a new CreateD4CGcpAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateD4CGcpAccountParams() *CreateD4CGcpAccountParams {
	return &CreateD4CGcpAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateD4CGcpAccountParamsWithTimeout creates a new CreateD4CGcpAccountParams object
// with the ability to set a timeout on a request.
func NewCreateD4CGcpAccountParamsWithTimeout(timeout time.Duration) *CreateD4CGcpAccountParams {
	return &CreateD4CGcpAccountParams{
		timeout: timeout,
	}
}

// NewCreateD4CGcpAccountParamsWithContext creates a new CreateD4CGcpAccountParams object
// with the ability to set a context for a request.
func NewCreateD4CGcpAccountParamsWithContext(ctx context.Context) *CreateD4CGcpAccountParams {
	return &CreateD4CGcpAccountParams{
		Context: ctx,
	}
}

// NewCreateD4CGcpAccountParamsWithHTTPClient creates a new CreateD4CGcpAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateD4CGcpAccountParamsWithHTTPClient(client *http.Client) *CreateD4CGcpAccountParams {
	return &CreateD4CGcpAccountParams{
		HTTPClient: client,
	}
}

/*
CreateD4CGcpAccountParams contains all the parameters to send to the API endpoint

	for the create d4 c gcp account operation.

	Typically these are written to a http.Request.
*/
type CreateD4CGcpAccountParams struct {

	// Body.
	Body *models.RegistrationGCPAccountCreateRequestExtV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create d4 c gcp account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateD4CGcpAccountParams) WithDefaults() *CreateD4CGcpAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create d4 c gcp account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateD4CGcpAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create d4 c gcp account params
func (o *CreateD4CGcpAccountParams) WithTimeout(timeout time.Duration) *CreateD4CGcpAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create d4 c gcp account params
func (o *CreateD4CGcpAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create d4 c gcp account params
func (o *CreateD4CGcpAccountParams) WithContext(ctx context.Context) *CreateD4CGcpAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create d4 c gcp account params
func (o *CreateD4CGcpAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create d4 c gcp account params
func (o *CreateD4CGcpAccountParams) WithHTTPClient(client *http.Client) *CreateD4CGcpAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create d4 c gcp account params
func (o *CreateD4CGcpAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create d4 c gcp account params
func (o *CreateD4CGcpAccountParams) WithBody(body *models.RegistrationGCPAccountCreateRequestExtV1) *CreateD4CGcpAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create d4 c gcp account params
func (o *CreateD4CGcpAccountParams) SetBody(body *models.RegistrationGCPAccountCreateRequestExtV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateD4CGcpAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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