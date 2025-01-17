// Code generated by go-swagger; DO NOT EDIT.

package ods

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

// NewCreateScanParams creates a new CreateScanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateScanParams() *CreateScanParams {
	return &CreateScanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScanParamsWithTimeout creates a new CreateScanParams object
// with the ability to set a timeout on a request.
func NewCreateScanParamsWithTimeout(timeout time.Duration) *CreateScanParams {
	return &CreateScanParams{
		timeout: timeout,
	}
}

// NewCreateScanParamsWithContext creates a new CreateScanParams object
// with the ability to set a context for a request.
func NewCreateScanParamsWithContext(ctx context.Context) *CreateScanParams {
	return &CreateScanParams{
		Context: ctx,
	}
}

// NewCreateScanParamsWithHTTPClient creates a new CreateScanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateScanParamsWithHTTPClient(client *http.Client) *CreateScanParams {
	return &CreateScanParams{
		HTTPClient: client,
	}
}

/*
CreateScanParams contains all the parameters to send to the API endpoint

	for the create scan operation.

	Typically these are written to a http.Request.
*/
type CreateScanParams struct {

	/* XCSUSERUUID.

	   The user ID
	*/
	XCSUSERUUID string

	// Body.
	Body *models.EntitiesODSScanRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateScanParams) WithDefaults() *CreateScanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateScanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create scan params
func (o *CreateScanParams) WithTimeout(timeout time.Duration) *CreateScanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create scan params
func (o *CreateScanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create scan params
func (o *CreateScanParams) WithContext(ctx context.Context) *CreateScanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create scan params
func (o *CreateScanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create scan params
func (o *CreateScanParams) WithHTTPClient(client *http.Client) *CreateScanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create scan params
func (o *CreateScanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERUUID adds the xCSUSERUUID to the create scan params
func (o *CreateScanParams) WithXCSUSERUUID(xCSUSERUUID string) *CreateScanParams {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the create scan params
func (o *CreateScanParams) SetXCSUSERUUID(xCSUSERUUID string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithBody adds the body to the create scan params
func (o *CreateScanParams) WithBody(body *models.EntitiesODSScanRequest) *CreateScanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create scan params
func (o *CreateScanParams) SetBody(body *models.EntitiesODSScanRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-CS-USERUUID
	if err := r.SetHeaderParam("X-CS-USERUUID", o.XCSUSERUUID); err != nil {
		return err
	}
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
