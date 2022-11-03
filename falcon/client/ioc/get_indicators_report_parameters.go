// Code generated by go-swagger; DO NOT EDIT.

package ioc

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

// NewGetIndicatorsReportParams creates a new GetIndicatorsReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIndicatorsReportParams() *GetIndicatorsReportParams {
	return &GetIndicatorsReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIndicatorsReportParamsWithTimeout creates a new GetIndicatorsReportParams object
// with the ability to set a timeout on a request.
func NewGetIndicatorsReportParamsWithTimeout(timeout time.Duration) *GetIndicatorsReportParams {
	return &GetIndicatorsReportParams{
		timeout: timeout,
	}
}

// NewGetIndicatorsReportParamsWithContext creates a new GetIndicatorsReportParams object
// with the ability to set a context for a request.
func NewGetIndicatorsReportParamsWithContext(ctx context.Context) *GetIndicatorsReportParams {
	return &GetIndicatorsReportParams{
		Context: ctx,
	}
}

// NewGetIndicatorsReportParamsWithHTTPClient creates a new GetIndicatorsReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIndicatorsReportParamsWithHTTPClient(client *http.Client) *GetIndicatorsReportParams {
	return &GetIndicatorsReportParams{
		HTTPClient: client,
	}
}

/*
GetIndicatorsReportParams contains all the parameters to send to the API endpoint

	for the get indicators report operation.

	Typically these are written to a http.Request.
*/
type GetIndicatorsReportParams struct {

	// Body.
	Body *models.APIIndicatorsReportRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get indicators report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIndicatorsReportParams) WithDefaults() *GetIndicatorsReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get indicators report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIndicatorsReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get indicators report params
func (o *GetIndicatorsReportParams) WithTimeout(timeout time.Duration) *GetIndicatorsReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get indicators report params
func (o *GetIndicatorsReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get indicators report params
func (o *GetIndicatorsReportParams) WithContext(ctx context.Context) *GetIndicatorsReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get indicators report params
func (o *GetIndicatorsReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get indicators report params
func (o *GetIndicatorsReportParams) WithHTTPClient(client *http.Client) *GetIndicatorsReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get indicators report params
func (o *GetIndicatorsReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get indicators report params
func (o *GetIndicatorsReportParams) WithBody(body *models.APIIndicatorsReportRequest) *GetIndicatorsReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get indicators report params
func (o *GetIndicatorsReportParams) SetBody(body *models.APIIndicatorsReportRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetIndicatorsReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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