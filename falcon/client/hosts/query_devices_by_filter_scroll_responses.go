// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

// QueryDevicesByFilterScrollReader is a Reader for the QueryDevicesByFilterScroll structure.
type QueryDevicesByFilterScrollReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDevicesByFilterScrollReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDevicesByFilterScrollOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryDevicesByFilterScrollForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryDevicesByFilterScrollTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryDevicesByFilterScrollDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryDevicesByFilterScrollOK creates a QueryDevicesByFilterScrollOK with default headers values
func NewQueryDevicesByFilterScrollOK() *QueryDevicesByFilterScrollOK {
	return &QueryDevicesByFilterScrollOK{}
}

/*
QueryDevicesByFilterScrollOK describes a response with status code 200, with default header values.

OK
*/
type QueryDevicesByFilterScrollOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainDeviceResponse
}

// IsSuccess returns true when this query devices by filter scroll o k response has a 2xx status code
func (o *QueryDevicesByFilterScrollOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query devices by filter scroll o k response has a 3xx status code
func (o *QueryDevicesByFilterScrollOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query devices by filter scroll o k response has a 4xx status code
func (o *QueryDevicesByFilterScrollOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query devices by filter scroll o k response has a 5xx status code
func (o *QueryDevicesByFilterScrollOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query devices by filter scroll o k response a status code equal to that given
func (o *QueryDevicesByFilterScrollOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryDevicesByFilterScrollOK) Error() string {
	return fmt.Sprintf("[GET /devices/queries/devices-scroll/v1][%d] queryDevicesByFilterScrollOK  %+v", 200, o.Payload)
}

func (o *QueryDevicesByFilterScrollOK) String() string {
	return fmt.Sprintf("[GET /devices/queries/devices-scroll/v1][%d] queryDevicesByFilterScrollOK  %+v", 200, o.Payload)
}

func (o *QueryDevicesByFilterScrollOK) GetPayload() *models.DomainDeviceResponse {
	return o.Payload
}

func (o *QueryDevicesByFilterScrollOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainDeviceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDevicesByFilterScrollForbidden creates a QueryDevicesByFilterScrollForbidden with default headers values
func NewQueryDevicesByFilterScrollForbidden() *QueryDevicesByFilterScrollForbidden {
	return &QueryDevicesByFilterScrollForbidden{}
}

/*
QueryDevicesByFilterScrollForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryDevicesByFilterScrollForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query devices by filter scroll forbidden response has a 2xx status code
func (o *QueryDevicesByFilterScrollForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query devices by filter scroll forbidden response has a 3xx status code
func (o *QueryDevicesByFilterScrollForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query devices by filter scroll forbidden response has a 4xx status code
func (o *QueryDevicesByFilterScrollForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query devices by filter scroll forbidden response has a 5xx status code
func (o *QueryDevicesByFilterScrollForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query devices by filter scroll forbidden response a status code equal to that given
func (o *QueryDevicesByFilterScrollForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryDevicesByFilterScrollForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/queries/devices-scroll/v1][%d] queryDevicesByFilterScrollForbidden  %+v", 403, o.Payload)
}

func (o *QueryDevicesByFilterScrollForbidden) String() string {
	return fmt.Sprintf("[GET /devices/queries/devices-scroll/v1][%d] queryDevicesByFilterScrollForbidden  %+v", 403, o.Payload)
}

func (o *QueryDevicesByFilterScrollForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDevicesByFilterScrollForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDevicesByFilterScrollTooManyRequests creates a QueryDevicesByFilterScrollTooManyRequests with default headers values
func NewQueryDevicesByFilterScrollTooManyRequests() *QueryDevicesByFilterScrollTooManyRequests {
	return &QueryDevicesByFilterScrollTooManyRequests{}
}

/*
QueryDevicesByFilterScrollTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryDevicesByFilterScrollTooManyRequests struct {

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

// IsSuccess returns true when this query devices by filter scroll too many requests response has a 2xx status code
func (o *QueryDevicesByFilterScrollTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query devices by filter scroll too many requests response has a 3xx status code
func (o *QueryDevicesByFilterScrollTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query devices by filter scroll too many requests response has a 4xx status code
func (o *QueryDevicesByFilterScrollTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query devices by filter scroll too many requests response has a 5xx status code
func (o *QueryDevicesByFilterScrollTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query devices by filter scroll too many requests response a status code equal to that given
func (o *QueryDevicesByFilterScrollTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueryDevicesByFilterScrollTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/queries/devices-scroll/v1][%d] queryDevicesByFilterScrollTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDevicesByFilterScrollTooManyRequests) String() string {
	return fmt.Sprintf("[GET /devices/queries/devices-scroll/v1][%d] queryDevicesByFilterScrollTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDevicesByFilterScrollTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDevicesByFilterScrollTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDevicesByFilterScrollDefault creates a QueryDevicesByFilterScrollDefault with default headers values
func NewQueryDevicesByFilterScrollDefault(code int) *QueryDevicesByFilterScrollDefault {
	return &QueryDevicesByFilterScrollDefault{
		_statusCode: code,
	}
}

/*
QueryDevicesByFilterScrollDefault describes a response with status code -1, with default header values.

OK
*/
type QueryDevicesByFilterScrollDefault struct {
	_statusCode int

	Payload *models.DomainDeviceResponse
}

// Code gets the status code for the query devices by filter scroll default response
func (o *QueryDevicesByFilterScrollDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this query devices by filter scroll default response has a 2xx status code
func (o *QueryDevicesByFilterScrollDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query devices by filter scroll default response has a 3xx status code
func (o *QueryDevicesByFilterScrollDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query devices by filter scroll default response has a 4xx status code
func (o *QueryDevicesByFilterScrollDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query devices by filter scroll default response has a 5xx status code
func (o *QueryDevicesByFilterScrollDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query devices by filter scroll default response a status code equal to that given
func (o *QueryDevicesByFilterScrollDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueryDevicesByFilterScrollDefault) Error() string {
	return fmt.Sprintf("[GET /devices/queries/devices-scroll/v1][%d] QueryDevicesByFilterScroll default  %+v", o._statusCode, o.Payload)
}

func (o *QueryDevicesByFilterScrollDefault) String() string {
	return fmt.Sprintf("[GET /devices/queries/devices-scroll/v1][%d] QueryDevicesByFilterScroll default  %+v", o._statusCode, o.Payload)
}

func (o *QueryDevicesByFilterScrollDefault) GetPayload() *models.DomainDeviceResponse {
	return o.Payload
}

func (o *QueryDevicesByFilterScrollDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainDeviceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
