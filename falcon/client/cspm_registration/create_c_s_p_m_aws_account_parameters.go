// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// NewCreateCSPMAwsAccountParams creates a new CreateCSPMAwsAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCSPMAwsAccountParams() *CreateCSPMAwsAccountParams {
	return &CreateCSPMAwsAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCSPMAwsAccountParamsWithTimeout creates a new CreateCSPMAwsAccountParams object
// with the ability to set a timeout on a request.
func NewCreateCSPMAwsAccountParamsWithTimeout(timeout time.Duration) *CreateCSPMAwsAccountParams {
	return &CreateCSPMAwsAccountParams{
		timeout: timeout,
	}
}

// NewCreateCSPMAwsAccountParamsWithContext creates a new CreateCSPMAwsAccountParams object
// with the ability to set a context for a request.
func NewCreateCSPMAwsAccountParamsWithContext(ctx context.Context) *CreateCSPMAwsAccountParams {
	return &CreateCSPMAwsAccountParams{
		Context: ctx,
	}
}

// NewCreateCSPMAwsAccountParamsWithHTTPClient creates a new CreateCSPMAwsAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCSPMAwsAccountParamsWithHTTPClient(client *http.Client) *CreateCSPMAwsAccountParams {
	return &CreateCSPMAwsAccountParams{
		HTTPClient: client,
	}
}

/* CreateCSPMAwsAccountParams contains all the parameters to send to the API endpoint
   for the create c s p m aws account operation.

   Typically these are written to a http.Request.
*/
type CreateCSPMAwsAccountParams struct {

	// Body.
	Body *models.RegistrationAWSAccountCreateRequestExtV2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create c s p m aws account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCSPMAwsAccountParams) WithDefaults() *CreateCSPMAwsAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create c s p m aws account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCSPMAwsAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create c s p m aws account params
func (o *CreateCSPMAwsAccountParams) WithTimeout(timeout time.Duration) *CreateCSPMAwsAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create c s p m aws account params
func (o *CreateCSPMAwsAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create c s p m aws account params
func (o *CreateCSPMAwsAccountParams) WithContext(ctx context.Context) *CreateCSPMAwsAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create c s p m aws account params
func (o *CreateCSPMAwsAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create c s p m aws account params
func (o *CreateCSPMAwsAccountParams) WithHTTPClient(client *http.Client) *CreateCSPMAwsAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create c s p m aws account params
func (o *CreateCSPMAwsAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create c s p m aws account params
func (o *CreateCSPMAwsAccountParams) WithBody(body *models.RegistrationAWSAccountCreateRequestExtV2) *CreateCSPMAwsAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create c s p m aws account params
func (o *CreateCSPMAwsAccountParams) SetBody(body *models.RegistrationAWSAccountCreateRequestExtV2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCSPMAwsAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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