// Code generated by go-swagger; DO NOT EDIT.

package iocs

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

// ProcessesRanOnReader is a Reader for the ProcessesRanOn structure.
type ProcessesRanOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProcessesRanOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProcessesRanOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProcessesRanOnForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewProcessesRanOnTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProcessesRanOnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProcessesRanOnOK creates a ProcessesRanOnOK with default headers values
func NewProcessesRanOnOK() *ProcessesRanOnOK {
	return &ProcessesRanOnOK{}
}

/* ProcessesRanOnOK describes a response with status code 200, with default header values.

OK
*/
type ProcessesRanOnOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIMsaReplyProcessesRanOn
}

func (o *ProcessesRanOnOK) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/processes/v1][%d] processesRanOnOK  %+v", 200, o.Payload)
}
func (o *ProcessesRanOnOK) GetPayload() *models.APIMsaReplyProcessesRanOn {
	return o.Payload
}

func (o *ProcessesRanOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIMsaReplyProcessesRanOn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProcessesRanOnForbidden creates a ProcessesRanOnForbidden with default headers values
func NewProcessesRanOnForbidden() *ProcessesRanOnForbidden {
	return &ProcessesRanOnForbidden{}
}

/* ProcessesRanOnForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProcessesRanOnForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *ProcessesRanOnForbidden) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/processes/v1][%d] processesRanOnForbidden  %+v", 403, o.Payload)
}
func (o *ProcessesRanOnForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ProcessesRanOnForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewProcessesRanOnTooManyRequests creates a ProcessesRanOnTooManyRequests with default headers values
func NewProcessesRanOnTooManyRequests() *ProcessesRanOnTooManyRequests {
	return &ProcessesRanOnTooManyRequests{}
}

/* ProcessesRanOnTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ProcessesRanOnTooManyRequests struct {

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

func (o *ProcessesRanOnTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/processes/v1][%d] processesRanOnTooManyRequests  %+v", 429, o.Payload)
}
func (o *ProcessesRanOnTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ProcessesRanOnTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewProcessesRanOnDefault creates a ProcessesRanOnDefault with default headers values
func NewProcessesRanOnDefault(code int) *ProcessesRanOnDefault {
	return &ProcessesRanOnDefault{
		_statusCode: code,
	}
}

/* ProcessesRanOnDefault describes a response with status code -1, with default header values.

OK
*/
type ProcessesRanOnDefault struct {
	_statusCode int

	Payload *models.APIMsaReplyProcessesRanOn
}

// Code gets the status code for the processes ran on default response
func (o *ProcessesRanOnDefault) Code() int {
	return o._statusCode
}

func (o *ProcessesRanOnDefault) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/processes/v1][%d] ProcessesRanOn default  %+v", o._statusCode, o.Payload)
}
func (o *ProcessesRanOnDefault) GetPayload() *models.APIMsaReplyProcessesRanOn {
	return o.Payload
}

func (o *ProcessesRanOnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMsaReplyProcessesRanOn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}