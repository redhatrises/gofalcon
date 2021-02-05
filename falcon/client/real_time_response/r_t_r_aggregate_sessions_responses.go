// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// RTRAggregateSessionsReader is a Reader for the RTRAggregateSessions structure.
type RTRAggregateSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRAggregateSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRAggregateSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRAggregateSessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRAggregateSessionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRAggregateSessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRAggregateSessionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRAggregateSessionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRAggregateSessionsOK creates a RTRAggregateSessionsOK with default headers values
func NewRTRAggregateSessionsOK() *RTRAggregateSessionsOK {
	return &RTRAggregateSessionsOK{}
}

/* RTRAggregateSessionsOK describes a response with status code 200, with default header values.

OK
*/
type RTRAggregateSessionsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

func (o *RTRAggregateSessionsOK) Error() string {
	return fmt.Sprintf("[POST /real-time-response/aggregates/sessions/GET/v1][%d] rTRAggregateSessionsOK  %+v", 200, o.Payload)
}
func (o *RTRAggregateSessionsOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *RTRAggregateSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRAggregateSessionsBadRequest creates a RTRAggregateSessionsBadRequest with default headers values
func NewRTRAggregateSessionsBadRequest() *RTRAggregateSessionsBadRequest {
	return &RTRAggregateSessionsBadRequest{}
}

/* RTRAggregateSessionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRAggregateSessionsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRAggregateSessionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/aggregates/sessions/GET/v1][%d] rTRAggregateSessionsBadRequest  %+v", 400, o.Payload)
}
func (o *RTRAggregateSessionsBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRAggregateSessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRAggregateSessionsForbidden creates a RTRAggregateSessionsForbidden with default headers values
func NewRTRAggregateSessionsForbidden() *RTRAggregateSessionsForbidden {
	return &RTRAggregateSessionsForbidden{}
}

/* RTRAggregateSessionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRAggregateSessionsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRAggregateSessionsForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/aggregates/sessions/GET/v1][%d] rTRAggregateSessionsForbidden  %+v", 403, o.Payload)
}
func (o *RTRAggregateSessionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRAggregateSessionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRAggregateSessionsNotFound creates a RTRAggregateSessionsNotFound with default headers values
func NewRTRAggregateSessionsNotFound() *RTRAggregateSessionsNotFound {
	return &RTRAggregateSessionsNotFound{}
}

/* RTRAggregateSessionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRAggregateSessionsNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRAggregateSessionsNotFound) Error() string {
	return fmt.Sprintf("[POST /real-time-response/aggregates/sessions/GET/v1][%d] rTRAggregateSessionsNotFound  %+v", 404, o.Payload)
}
func (o *RTRAggregateSessionsNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRAggregateSessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRAggregateSessionsTooManyRequests creates a RTRAggregateSessionsTooManyRequests with default headers values
func NewRTRAggregateSessionsTooManyRequests() *RTRAggregateSessionsTooManyRequests {
	return &RTRAggregateSessionsTooManyRequests{}
}

/* RTRAggregateSessionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRAggregateSessionsTooManyRequests struct {

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

func (o *RTRAggregateSessionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/aggregates/sessions/GET/v1][%d] rTRAggregateSessionsTooManyRequests  %+v", 429, o.Payload)
}
func (o *RTRAggregateSessionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRAggregateSessionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRAggregateSessionsDefault creates a RTRAggregateSessionsDefault with default headers values
func NewRTRAggregateSessionsDefault(code int) *RTRAggregateSessionsDefault {
	return &RTRAggregateSessionsDefault{
		_statusCode: code,
	}
}

/* RTRAggregateSessionsDefault describes a response with status code -1, with default header values.

OK
*/
type RTRAggregateSessionsDefault struct {
	_statusCode int

	Payload *models.MsaAggregatesResponse
}

// Code gets the status code for the r t r aggregate sessions default response
func (o *RTRAggregateSessionsDefault) Code() int {
	return o._statusCode
}

func (o *RTRAggregateSessionsDefault) Error() string {
	return fmt.Sprintf("[POST /real-time-response/aggregates/sessions/GET/v1][%d] RTR-AggregateSessions default  %+v", o._statusCode, o.Payload)
}
func (o *RTRAggregateSessionsDefault) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *RTRAggregateSessionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}