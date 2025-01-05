// Code generated by go-swagger; DO NOT EDIT.

package humio_auth_proxy

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

// NewStopSearchV1Params creates a new StopSearchV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopSearchV1Params() *StopSearchV1Params {
	return &StopSearchV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopSearchV1ParamsWithTimeout creates a new StopSearchV1Params object
// with the ability to set a timeout on a request.
func NewStopSearchV1ParamsWithTimeout(timeout time.Duration) *StopSearchV1Params {
	return &StopSearchV1Params{
		timeout: timeout,
	}
}

// NewStopSearchV1ParamsWithContext creates a new StopSearchV1Params object
// with the ability to set a context for a request.
func NewStopSearchV1ParamsWithContext(ctx context.Context) *StopSearchV1Params {
	return &StopSearchV1Params{
		Context: ctx,
	}
}

// NewStopSearchV1ParamsWithHTTPClient creates a new StopSearchV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewStopSearchV1ParamsWithHTTPClient(client *http.Client) *StopSearchV1Params {
	return &StopSearchV1Params{
		HTTPClient: client,
	}
}

/*
StopSearchV1Params contains all the parameters to send to the API endpoint

	for the stop search v1 operation.

	Typically these are written to a http.Request.
*/
type StopSearchV1Params struct {

	/* ID.

	   id of query
	*/
	ID string

	/* Repository.

	   name of repository
	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop search v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopSearchV1Params) WithDefaults() *StopSearchV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop search v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopSearchV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop search v1 params
func (o *StopSearchV1Params) WithTimeout(timeout time.Duration) *StopSearchV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop search v1 params
func (o *StopSearchV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop search v1 params
func (o *StopSearchV1Params) WithContext(ctx context.Context) *StopSearchV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop search v1 params
func (o *StopSearchV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop search v1 params
func (o *StopSearchV1Params) WithHTTPClient(client *http.Client) *StopSearchV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop search v1 params
func (o *StopSearchV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the stop search v1 params
func (o *StopSearchV1Params) WithID(id string) *StopSearchV1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the stop search v1 params
func (o *StopSearchV1Params) SetID(id string) {
	o.ID = id
}

// WithRepository adds the repository to the stop search v1 params
func (o *StopSearchV1Params) WithRepository(repository string) *StopSearchV1Params {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the stop search v1 params
func (o *StopSearchV1Params) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *StopSearchV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}