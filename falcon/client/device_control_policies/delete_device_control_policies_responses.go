// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// DeleteDeviceControlPoliciesReader is a Reader for the DeleteDeviceControlPolicies structure.
type DeleteDeviceControlPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceControlPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeviceControlPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteDeviceControlPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDeviceControlPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteDeviceControlPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDeviceControlPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDeviceControlPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDeviceControlPoliciesOK creates a DeleteDeviceControlPoliciesOK with default headers values
func NewDeleteDeviceControlPoliciesOK() *DeleteDeviceControlPoliciesOK {
	return &DeleteDeviceControlPoliciesOK{}
}

/*
DeleteDeviceControlPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDeviceControlPoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this delete device control policies o k response has a 2xx status code
func (o *DeleteDeviceControlPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete device control policies o k response has a 3xx status code
func (o *DeleteDeviceControlPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device control policies o k response has a 4xx status code
func (o *DeleteDeviceControlPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete device control policies o k response has a 5xx status code
func (o *DeleteDeviceControlPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete device control policies o k response a status code equal to that given
func (o *DeleteDeviceControlPoliciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteDeviceControlPoliciesOK) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceControlPoliciesOK) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceControlPoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteDeviceControlPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceControlPoliciesForbidden creates a DeleteDeviceControlPoliciesForbidden with default headers values
func NewDeleteDeviceControlPoliciesForbidden() *DeleteDeviceControlPoliciesForbidden {
	return &DeleteDeviceControlPoliciesForbidden{}
}

/*
DeleteDeviceControlPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDeviceControlPoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this delete device control policies forbidden response has a 2xx status code
func (o *DeleteDeviceControlPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device control policies forbidden response has a 3xx status code
func (o *DeleteDeviceControlPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device control policies forbidden response has a 4xx status code
func (o *DeleteDeviceControlPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete device control policies forbidden response has a 5xx status code
func (o *DeleteDeviceControlPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete device control policies forbidden response a status code equal to that given
func (o *DeleteDeviceControlPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteDeviceControlPoliciesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDeviceControlPoliciesForbidden) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDeviceControlPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *DeleteDeviceControlPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceControlPoliciesNotFound creates a DeleteDeviceControlPoliciesNotFound with default headers values
func NewDeleteDeviceControlPoliciesNotFound() *DeleteDeviceControlPoliciesNotFound {
	return &DeleteDeviceControlPoliciesNotFound{}
}

/*
DeleteDeviceControlPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDeviceControlPoliciesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this delete device control policies not found response has a 2xx status code
func (o *DeleteDeviceControlPoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device control policies not found response has a 3xx status code
func (o *DeleteDeviceControlPoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device control policies not found response has a 4xx status code
func (o *DeleteDeviceControlPoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete device control policies not found response has a 5xx status code
func (o *DeleteDeviceControlPoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete device control policies not found response a status code equal to that given
func (o *DeleteDeviceControlPoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteDeviceControlPoliciesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeviceControlPoliciesNotFound) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeviceControlPoliciesNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteDeviceControlPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceControlPoliciesTooManyRequests creates a DeleteDeviceControlPoliciesTooManyRequests with default headers values
func NewDeleteDeviceControlPoliciesTooManyRequests() *DeleteDeviceControlPoliciesTooManyRequests {
	return &DeleteDeviceControlPoliciesTooManyRequests{}
}

/*
DeleteDeviceControlPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteDeviceControlPoliciesTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this delete device control policies too many requests response has a 2xx status code
func (o *DeleteDeviceControlPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device control policies too many requests response has a 3xx status code
func (o *DeleteDeviceControlPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device control policies too many requests response has a 4xx status code
func (o *DeleteDeviceControlPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete device control policies too many requests response has a 5xx status code
func (o *DeleteDeviceControlPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete device control policies too many requests response a status code equal to that given
func (o *DeleteDeviceControlPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteDeviceControlPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteDeviceControlPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteDeviceControlPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteDeviceControlPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceControlPoliciesInternalServerError creates a DeleteDeviceControlPoliciesInternalServerError with default headers values
func NewDeleteDeviceControlPoliciesInternalServerError() *DeleteDeviceControlPoliciesInternalServerError {
	return &DeleteDeviceControlPoliciesInternalServerError{}
}

/*
DeleteDeviceControlPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDeviceControlPoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this delete device control policies internal server error response has a 2xx status code
func (o *DeleteDeviceControlPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete device control policies internal server error response has a 3xx status code
func (o *DeleteDeviceControlPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete device control policies internal server error response has a 4xx status code
func (o *DeleteDeviceControlPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete device control policies internal server error response has a 5xx status code
func (o *DeleteDeviceControlPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete device control policies internal server error response a status code equal to that given
func (o *DeleteDeviceControlPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteDeviceControlPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDeviceControlPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDeviceControlPoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteDeviceControlPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceControlPoliciesDefault creates a DeleteDeviceControlPoliciesDefault with default headers values
func NewDeleteDeviceControlPoliciesDefault(code int) *DeleteDeviceControlPoliciesDefault {
	return &DeleteDeviceControlPoliciesDefault{
		_statusCode: code,
	}
}

/*
DeleteDeviceControlPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type DeleteDeviceControlPoliciesDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the delete device control policies default response
func (o *DeleteDeviceControlPoliciesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete device control policies default response has a 2xx status code
func (o *DeleteDeviceControlPoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete device control policies default response has a 3xx status code
func (o *DeleteDeviceControlPoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete device control policies default response has a 4xx status code
func (o *DeleteDeviceControlPoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete device control policies default response has a 5xx status code
func (o *DeleteDeviceControlPoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete device control policies default response a status code equal to that given
func (o *DeleteDeviceControlPoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteDeviceControlPoliciesDefault) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeviceControlPoliciesDefault) String() string {
	return fmt.Sprintf("[DELETE /policy/entities/device-control/v1][%d] deleteDeviceControlPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeviceControlPoliciesDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteDeviceControlPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
