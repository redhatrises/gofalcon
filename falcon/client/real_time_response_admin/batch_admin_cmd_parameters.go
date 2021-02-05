// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewBatchAdminCmdParams creates a new BatchAdminCmdParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBatchAdminCmdParams() *BatchAdminCmdParams {
	return &BatchAdminCmdParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewBatchAdminCmdParamsWithTimeout creates a new BatchAdminCmdParams object
// with the ability to set a timeout on a request.
func NewBatchAdminCmdParamsWithTimeout(timeout time.Duration) *BatchAdminCmdParams {
	return &BatchAdminCmdParams{
		requestTimeout: timeout,
	}
}

// NewBatchAdminCmdParamsWithContext creates a new BatchAdminCmdParams object
// with the ability to set a context for a request.
func NewBatchAdminCmdParamsWithContext(ctx context.Context) *BatchAdminCmdParams {
	return &BatchAdminCmdParams{
		Context: ctx,
	}
}

// NewBatchAdminCmdParamsWithHTTPClient creates a new BatchAdminCmdParams object
// with the ability to set a custom HTTPClient for a request.
func NewBatchAdminCmdParamsWithHTTPClient(client *http.Client) *BatchAdminCmdParams {
	return &BatchAdminCmdParams{
		HTTPClient: client,
	}
}

/* BatchAdminCmdParams contains all the parameters to send to the API endpoint
   for the batch admin cmd operation.

   Typically these are written to a http.Request.
*/
type BatchAdminCmdParams struct {

	/* Body.

	     Use this endpoint to run these [real time response commands](https://falcon.crowdstrike.com/support/documentation/11/getting-started-guide#rtr_commands):
	- `cat`
	- `cd`
	- `clear`
	- `cp`
	- `encrypt`
	- `env`
	- `eventlog`
	- `filehash`
	- `get`
	- `getsid`
	- `help`
	- `history`
	- `ipconfig`
	- `kill`
	- `ls`
	- `map`
	- `memdump`
	- `mkdir`
	- `mount`
	- `mv`
	- `netstat`
	- `ps`
	- `put`
	- `reg query`
	- `reg set`
	- `reg delete`
	- `reg load`
	- `reg unload`
	- `restart`
	- `rm`
	- `run`
	- `runscript`
	- `shutdown`
	- `unmap`
	- `update history`
	- `update install`
	- `update list`
	- `update query`
	- `xmemdump`
	- `zip`

	**`base_command`** Active-Responder command type we are going to execute, for example: `get` or `cp`.  Refer to the RTR documentation for the full list of commands.
	**`batch_id`** Batch ID to execute the command on.  Received from `/real-time-response/combined/init-sessions/v1`.
	**`command_string`** Full command string for the command. For example  `get some_file.txt`
	**`optional_hosts`** List of a subset of hosts we want to run the command on.  If this list is supplied, only these hosts will receive the command.
	*/
	Body *models.DomainBatchExecuteCommandRequest

	/* Timeout.

	   Timeout for how long to wait for the request in seconds, default timeout is 30 seconds. Maximum is 10 minutes.

	   Default: 30
	*/
	Timeout *int64

	/* TimeoutDuration.

	   Timeout duration for for how long to wait for the request in duration syntax. Example, `10s`. Valid units: `ns, us, ms, s, m, h`. Maximum is 10 minutes.

	   Default: "30s"
	*/
	TimeoutDuration *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the batch admin cmd params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchAdminCmdParams) WithDefaults() *BatchAdminCmdParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the batch admin cmd params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchAdminCmdParams) SetDefaults() {
	var (
		timeoutDefault = int64(30)

		timeoutDurationDefault = string("30s")
	)

	val := BatchAdminCmdParams{
		Timeout:         &timeoutDefault,
		TimeoutDuration: &timeoutDurationDefault,
	}

	val.requestTimeout = o.requestTimeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithRequestTimeout adds the timeout to the batch admin cmd params
func (o *BatchAdminCmdParams) WithRequestTimeout(timeout time.Duration) *BatchAdminCmdParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the batch admin cmd params
func (o *BatchAdminCmdParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the batch admin cmd params
func (o *BatchAdminCmdParams) WithContext(ctx context.Context) *BatchAdminCmdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch admin cmd params
func (o *BatchAdminCmdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch admin cmd params
func (o *BatchAdminCmdParams) WithHTTPClient(client *http.Client) *BatchAdminCmdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch admin cmd params
func (o *BatchAdminCmdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batch admin cmd params
func (o *BatchAdminCmdParams) WithBody(body *models.DomainBatchExecuteCommandRequest) *BatchAdminCmdParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batch admin cmd params
func (o *BatchAdminCmdParams) SetBody(body *models.DomainBatchExecuteCommandRequest) {
	o.Body = body
}

// WithTimeout adds the timeout to the batch admin cmd params
func (o *BatchAdminCmdParams) WithTimeout(timeout *int64) *BatchAdminCmdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch admin cmd params
func (o *BatchAdminCmdParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WithTimeoutDuration adds the timeoutDuration to the batch admin cmd params
func (o *BatchAdminCmdParams) WithTimeoutDuration(timeoutDuration *string) *BatchAdminCmdParams {
	o.SetTimeoutDuration(timeoutDuration)
	return o
}

// SetTimeoutDuration adds the timeoutDuration to the batch admin cmd params
func (o *BatchAdminCmdParams) SetTimeoutDuration(timeoutDuration *string) {
	o.TimeoutDuration = timeoutDuration
}

// WriteToRequest writes these params to a swagger request
func (o *BatchAdminCmdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64

		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {

			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}
	}

	if o.TimeoutDuration != nil {

		// query param timeout_duration
		var qrTimeoutDuration string

		if o.TimeoutDuration != nil {
			qrTimeoutDuration = *o.TimeoutDuration
		}
		qTimeoutDuration := qrTimeoutDuration
		if qTimeoutDuration != "" {

			if err := r.SetQueryParam("timeout_duration", qTimeoutDuration); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}