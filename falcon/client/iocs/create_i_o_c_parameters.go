// Code generated by go-swagger; DO NOT EDIT.

package iocs

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

// NewCreateIOCParams creates a new CreateIOCParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIOCParams() *CreateIOCParams {
	return &CreateIOCParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIOCParamsWithTimeout creates a new CreateIOCParams object
// with the ability to set a timeout on a request.
func NewCreateIOCParamsWithTimeout(timeout time.Duration) *CreateIOCParams {
	return &CreateIOCParams{
		timeout: timeout,
	}
}

// NewCreateIOCParamsWithContext creates a new CreateIOCParams object
// with the ability to set a context for a request.
func NewCreateIOCParamsWithContext(ctx context.Context) *CreateIOCParams {
	return &CreateIOCParams{
		Context: ctx,
	}
}

// NewCreateIOCParamsWithHTTPClient creates a new CreateIOCParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIOCParamsWithHTTPClient(client *http.Client) *CreateIOCParams {
	return &CreateIOCParams{
		HTTPClient: client,
	}
}

/* CreateIOCParams contains all the parameters to send to the API endpoint
   for the create i o c operation.

   Typically these are written to a http.Request.
*/
type CreateIOCParams struct {

	/* Body.

	     Create a new IOC by providing a JSON object that includes these key/value pairs:

	**type** (required): The type of the indicator. Valid values:

	- sha256: A hex-encoded sha256 hash string. Length - min: 64, max: 64.

	- md5: A hex-encoded md5 hash string. Length - min 32, max: 32.

	- domain: A domain name. Length - min: 1, max: 200.

	- ipv4: An IPv4 address. Must be a valid IP address.

	- ipv6: An IPv6 address. Must be a valid IP address.

	**value** (required): The string representation of the indicator.

	**policy** (required): Action to take when a host observes the custom IOC. Values:

	- detect: Enable detections for this custom IOC

	- none: Disable detections for this custom IOC

	**share_level** (optional): Visibility of this custom IOC. All custom IOCs are visible only within your customer account, so only one value is valid:

	- red

	**expiration_days** (optional): Number of days this custom IOC is active. Only applies for the types `domain`, `ipv4`, and `ipv6`.

	**source** (optional): The source where this indicator originated. This can be used for tracking where this indicator was defined. Limit 200 characters.

	**description** (optional): Descriptive label for this custom IOC
	*/
	Body []*models.APIIOCViewRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create i o c params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIOCParams) WithDefaults() *CreateIOCParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create i o c params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIOCParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create i o c params
func (o *CreateIOCParams) WithTimeout(timeout time.Duration) *CreateIOCParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create i o c params
func (o *CreateIOCParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create i o c params
func (o *CreateIOCParams) WithContext(ctx context.Context) *CreateIOCParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create i o c params
func (o *CreateIOCParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create i o c params
func (o *CreateIOCParams) WithHTTPClient(client *http.Client) *CreateIOCParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create i o c params
func (o *CreateIOCParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create i o c params
func (o *CreateIOCParams) WithBody(body []*models.APIIOCViewRecord) *CreateIOCParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create i o c params
func (o *CreateIOCParams) SetBody(body []*models.APIIOCViewRecord) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIOCParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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