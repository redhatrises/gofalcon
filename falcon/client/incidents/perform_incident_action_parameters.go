// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewPerformIncidentActionParams creates a new PerformIncidentActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformIncidentActionParams() *PerformIncidentActionParams {
	return &PerformIncidentActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformIncidentActionParamsWithTimeout creates a new PerformIncidentActionParams object
// with the ability to set a timeout on a request.
func NewPerformIncidentActionParamsWithTimeout(timeout time.Duration) *PerformIncidentActionParams {
	return &PerformIncidentActionParams{
		timeout: timeout,
	}
}

// NewPerformIncidentActionParamsWithContext creates a new PerformIncidentActionParams object
// with the ability to set a context for a request.
func NewPerformIncidentActionParamsWithContext(ctx context.Context) *PerformIncidentActionParams {
	return &PerformIncidentActionParams{
		Context: ctx,
	}
}

// NewPerformIncidentActionParamsWithHTTPClient creates a new PerformIncidentActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformIncidentActionParamsWithHTTPClient(client *http.Client) *PerformIncidentActionParams {
	return &PerformIncidentActionParams{
		HTTPClient: client,
	}
}

/* PerformIncidentActionParams contains all the parameters to send to the API endpoint
   for the perform incident action operation.

   Typically these are written to a http.Request.
*/
type PerformIncidentActionParams struct {

	// Body.
	Body *models.MsaEntityActionRequestV2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform incident action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformIncidentActionParams) WithDefaults() *PerformIncidentActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform incident action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformIncidentActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform incident action params
func (o *PerformIncidentActionParams) WithTimeout(timeout time.Duration) *PerformIncidentActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform incident action params
func (o *PerformIncidentActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform incident action params
func (o *PerformIncidentActionParams) WithContext(ctx context.Context) *PerformIncidentActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform incident action params
func (o *PerformIncidentActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform incident action params
func (o *PerformIncidentActionParams) WithHTTPClient(client *http.Client) *PerformIncidentActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform incident action params
func (o *PerformIncidentActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the perform incident action params
func (o *PerformIncidentActionParams) WithBody(body *models.MsaEntityActionRequestV2) *PerformIncidentActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform incident action params
func (o *PerformIncidentActionParams) SetBody(body *models.MsaEntityActionRequestV2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PerformIncidentActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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