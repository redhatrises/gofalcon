// Code generated by go-swagger; DO NOT EDIT.

package filevantage

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

// NewDeletePoliciesParams creates a new DeletePoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePoliciesParams() *DeletePoliciesParams {
	return &DeletePoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePoliciesParamsWithTimeout creates a new DeletePoliciesParams object
// with the ability to set a timeout on a request.
func NewDeletePoliciesParamsWithTimeout(timeout time.Duration) *DeletePoliciesParams {
	return &DeletePoliciesParams{
		timeout: timeout,
	}
}

// NewDeletePoliciesParamsWithContext creates a new DeletePoliciesParams object
// with the ability to set a context for a request.
func NewDeletePoliciesParamsWithContext(ctx context.Context) *DeletePoliciesParams {
	return &DeletePoliciesParams{
		Context: ctx,
	}
}

// NewDeletePoliciesParamsWithHTTPClient creates a new DeletePoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePoliciesParamsWithHTTPClient(client *http.Client) *DeletePoliciesParams {
	return &DeletePoliciesParams{
		HTTPClient: client,
	}
}

/*
DeletePoliciesParams contains all the parameters to send to the API endpoint

	for the delete policies operation.

	Typically these are written to a http.Request.
*/
type DeletePoliciesParams struct {

	/* Ids.

	   One or more (up to 500) policy ids in the form of `ids=ID1&ids=ID2`
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePoliciesParams) WithDefaults() *DeletePoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete policies params
func (o *DeletePoliciesParams) WithTimeout(timeout time.Duration) *DeletePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete policies params
func (o *DeletePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete policies params
func (o *DeletePoliciesParams) WithContext(ctx context.Context) *DeletePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete policies params
func (o *DeletePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete policies params
func (o *DeletePoliciesParams) WithHTTPClient(client *http.Client) *DeletePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete policies params
func (o *DeletePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the delete policies params
func (o *DeletePoliciesParams) WithIds(ids []string) *DeletePoliciesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete policies params
func (o *DeletePoliciesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamDeletePolicies binds the parameter ids
func (o *DeletePoliciesParams) bindParamIds(formats strfmt.Registry) []string {
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
