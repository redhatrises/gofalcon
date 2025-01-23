// Code generated by go-swagger; DO NOT EDIT.

package release_notes

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

// NewGetEntityIDsByQueryPOSTParams creates a new GetEntityIDsByQueryPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEntityIDsByQueryPOSTParams() *GetEntityIDsByQueryPOSTParams {
	return &GetEntityIDsByQueryPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEntityIDsByQueryPOSTParamsWithTimeout creates a new GetEntityIDsByQueryPOSTParams object
// with the ability to set a timeout on a request.
func NewGetEntityIDsByQueryPOSTParamsWithTimeout(timeout time.Duration) *GetEntityIDsByQueryPOSTParams {
	return &GetEntityIDsByQueryPOSTParams{
		timeout: timeout,
	}
}

// NewGetEntityIDsByQueryPOSTParamsWithContext creates a new GetEntityIDsByQueryPOSTParams object
// with the ability to set a context for a request.
func NewGetEntityIDsByQueryPOSTParamsWithContext(ctx context.Context) *GetEntityIDsByQueryPOSTParams {
	return &GetEntityIDsByQueryPOSTParams{
		Context: ctx,
	}
}

// NewGetEntityIDsByQueryPOSTParamsWithHTTPClient creates a new GetEntityIDsByQueryPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEntityIDsByQueryPOSTParamsWithHTTPClient(client *http.Client) *GetEntityIDsByQueryPOSTParams {
	return &GetEntityIDsByQueryPOSTParams{
		HTTPClient: client,
	}
}

/*
GetEntityIDsByQueryPOSTParams contains all the parameters to send to the API endpoint

	for the get entity i ds by query p o s t operation.

	Typically these are written to a http.Request.
*/
type GetEntityIDsByQueryPOSTParams struct {

	/* Authorization.

	   authorization header
	*/
	Authorization string

	/* XCSUSERNAME.

	   user name
	*/
	XCSUSERNAME *string

	// Body.
	Body *models.ReleasenotesEntitiesGetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get entity i ds by query p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEntityIDsByQueryPOSTParams) WithDefaults() *GetEntityIDsByQueryPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get entity i ds by query p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEntityIDsByQueryPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) WithTimeout(timeout time.Duration) *GetEntityIDsByQueryPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) WithContext(ctx context.Context) *GetEntityIDsByQueryPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) WithHTTPClient(client *http.Client) *GetEntityIDsByQueryPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) WithAuthorization(authorization string) *GetEntityIDsByQueryPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXCSUSERNAME adds the xCSUSERNAME to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) WithXCSUSERNAME(xCSUSERNAME *string) *GetEntityIDsByQueryPOSTParams {
	o.SetXCSUSERNAME(xCSUSERNAME)
	return o
}

// SetXCSUSERNAME adds the xCSUSERNAME to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) SetXCSUSERNAME(xCSUSERNAME *string) {
	o.XCSUSERNAME = xCSUSERNAME
}

// WithBody adds the body to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) WithBody(body *models.ReleasenotesEntitiesGetRequest) *GetEntityIDsByQueryPOSTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get entity i ds by query p o s t params
func (o *GetEntityIDsByQueryPOSTParams) SetBody(body *models.ReleasenotesEntitiesGetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetEntityIDsByQueryPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.XCSUSERNAME != nil {

		// header param X-CS-USERNAME
		if err := r.SetHeaderParam("X-CS-USERNAME", *o.XCSUSERNAME); err != nil {
			return err
		}
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