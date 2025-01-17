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
	"github.com/go-openapi/swag"
)

// NewGetScansByScanIdsV2Params creates a new GetScansByScanIdsV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScansByScanIdsV2Params() *GetScansByScanIdsV2Params {
	return &GetScansByScanIdsV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScansByScanIdsV2ParamsWithTimeout creates a new GetScansByScanIdsV2Params object
// with the ability to set a timeout on a request.
func NewGetScansByScanIdsV2ParamsWithTimeout(timeout time.Duration) *GetScansByScanIdsV2Params {
	return &GetScansByScanIdsV2Params{
		timeout: timeout,
	}
}

// NewGetScansByScanIdsV2ParamsWithContext creates a new GetScansByScanIdsV2Params object
// with the ability to set a context for a request.
func NewGetScansByScanIdsV2ParamsWithContext(ctx context.Context) *GetScansByScanIdsV2Params {
	return &GetScansByScanIdsV2Params{
		Context: ctx,
	}
}

// NewGetScansByScanIdsV2ParamsWithHTTPClient creates a new GetScansByScanIdsV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetScansByScanIdsV2ParamsWithHTTPClient(client *http.Client) *GetScansByScanIdsV2Params {
	return &GetScansByScanIdsV2Params{
		HTTPClient: client,
	}
}

/*
GetScansByScanIdsV2Params contains all the parameters to send to the API endpoint

	for the get scans by scan ids v2 operation.

	Typically these are written to a http.Request.
*/
type GetScansByScanIdsV2Params struct {

	/* XCSUSERUUID.

	   The user ID
	*/
	XCSUSERUUID string

	/* Ids.

	   The scan IDs to retrieve the scan entities
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scans by scan ids v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScansByScanIdsV2Params) WithDefaults() *GetScansByScanIdsV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scans by scan ids v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScansByScanIdsV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) WithTimeout(timeout time.Duration) *GetScansByScanIdsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) WithContext(ctx context.Context) *GetScansByScanIdsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) WithHTTPClient(client *http.Client) *GetScansByScanIdsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERUUID adds the xCSUSERUUID to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) WithXCSUSERUUID(xCSUSERUUID string) *GetScansByScanIdsV2Params {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) SetXCSUSERUUID(xCSUSERUUID string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithIds adds the ids to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) WithIds(ids []string) *GetScansByScanIdsV2Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get scans by scan ids v2 params
func (o *GetScansByScanIdsV2Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetScansByScanIdsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-CS-USERUUID
	if err := r.SetHeaderParam("X-CS-USERUUID", o.XCSUSERUUID); err != nil {
		return err
	}

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

// bindParamGetScansByScanIdsV2 binds the parameter ids
func (o *GetScansByScanIdsV2Params) bindParamIds(formats strfmt.Registry) []string {
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
