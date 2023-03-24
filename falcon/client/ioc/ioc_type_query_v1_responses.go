// Code generated by go-swagger; DO NOT EDIT.

package ioc

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

// IocTypeQueryV1Reader is a Reader for the IocTypeQueryV1 structure.
type IocTypeQueryV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IocTypeQueryV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIocTypeQueryV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIocTypeQueryV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewIocTypeQueryV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIocTypeQueryV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIocTypeQueryV1OK creates a IocTypeQueryV1OK with default headers values
func NewIocTypeQueryV1OK() *IocTypeQueryV1OK {
	return &IocTypeQueryV1OK{}
}

/*
IocTypeQueryV1OK describes a response with status code 200, with default header values.

OK
*/
type IocTypeQueryV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIIndicatorQueryRespV1
}

// IsSuccess returns true when this ioc type query v1 o k response has a 2xx status code
func (o *IocTypeQueryV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ioc type query v1 o k response has a 3xx status code
func (o *IocTypeQueryV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ioc type query v1 o k response has a 4xx status code
func (o *IocTypeQueryV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ioc type query v1 o k response has a 5xx status code
func (o *IocTypeQueryV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this ioc type query v1 o k response a status code equal to that given
func (o *IocTypeQueryV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ioc type query v1 o k response
func (o *IocTypeQueryV1OK) Code() int {
	return 200
}

func (o *IocTypeQueryV1OK) Error() string {
	return fmt.Sprintf("[GET /iocs/queries/ioc-types/v1][%d] iocTypeQueryV1OK  %+v", 200, o.Payload)
}

func (o *IocTypeQueryV1OK) String() string {
	return fmt.Sprintf("[GET /iocs/queries/ioc-types/v1][%d] iocTypeQueryV1OK  %+v", 200, o.Payload)
}

func (o *IocTypeQueryV1OK) GetPayload() *models.APIIndicatorQueryRespV1 {
	return o.Payload
}

func (o *IocTypeQueryV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.APIIndicatorQueryRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIocTypeQueryV1Forbidden creates a IocTypeQueryV1Forbidden with default headers values
func NewIocTypeQueryV1Forbidden() *IocTypeQueryV1Forbidden {
	return &IocTypeQueryV1Forbidden{}
}

/*
IocTypeQueryV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type IocTypeQueryV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this ioc type query v1 forbidden response has a 2xx status code
func (o *IocTypeQueryV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ioc type query v1 forbidden response has a 3xx status code
func (o *IocTypeQueryV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ioc type query v1 forbidden response has a 4xx status code
func (o *IocTypeQueryV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ioc type query v1 forbidden response has a 5xx status code
func (o *IocTypeQueryV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ioc type query v1 forbidden response a status code equal to that given
func (o *IocTypeQueryV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ioc type query v1 forbidden response
func (o *IocTypeQueryV1Forbidden) Code() int {
	return 403
}

func (o *IocTypeQueryV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /iocs/queries/ioc-types/v1][%d] iocTypeQueryV1Forbidden  %+v", 403, o.Payload)
}

func (o *IocTypeQueryV1Forbidden) String() string {
	return fmt.Sprintf("[GET /iocs/queries/ioc-types/v1][%d] iocTypeQueryV1Forbidden  %+v", 403, o.Payload)
}

func (o *IocTypeQueryV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *IocTypeQueryV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewIocTypeQueryV1TooManyRequests creates a IocTypeQueryV1TooManyRequests with default headers values
func NewIocTypeQueryV1TooManyRequests() *IocTypeQueryV1TooManyRequests {
	return &IocTypeQueryV1TooManyRequests{}
}

/*
IocTypeQueryV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type IocTypeQueryV1TooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

// IsSuccess returns true when this ioc type query v1 too many requests response has a 2xx status code
func (o *IocTypeQueryV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ioc type query v1 too many requests response has a 3xx status code
func (o *IocTypeQueryV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ioc type query v1 too many requests response has a 4xx status code
func (o *IocTypeQueryV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this ioc type query v1 too many requests response has a 5xx status code
func (o *IocTypeQueryV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this ioc type query v1 too many requests response a status code equal to that given
func (o *IocTypeQueryV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the ioc type query v1 too many requests response
func (o *IocTypeQueryV1TooManyRequests) Code() int {
	return 429
}

func (o *IocTypeQueryV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /iocs/queries/ioc-types/v1][%d] iocTypeQueryV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *IocTypeQueryV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /iocs/queries/ioc-types/v1][%d] iocTypeQueryV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *IocTypeQueryV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *IocTypeQueryV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewIocTypeQueryV1Default creates a IocTypeQueryV1Default with default headers values
func NewIocTypeQueryV1Default(code int) *IocTypeQueryV1Default {
	return &IocTypeQueryV1Default{
		_statusCode: code,
	}
}

/*
IocTypeQueryV1Default describes a response with status code -1, with default header values.

OK
*/
type IocTypeQueryV1Default struct {
	_statusCode int

	Payload *models.APIIndicatorQueryRespV1
}

// IsSuccess returns true when this ioc type query v1 default response has a 2xx status code
func (o *IocTypeQueryV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ioc type query v1 default response has a 3xx status code
func (o *IocTypeQueryV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ioc type query v1 default response has a 4xx status code
func (o *IocTypeQueryV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ioc type query v1 default response has a 5xx status code
func (o *IocTypeQueryV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ioc type query v1 default response a status code equal to that given
func (o *IocTypeQueryV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ioc type query v1 default response
func (o *IocTypeQueryV1Default) Code() int {
	return o._statusCode
}

func (o *IocTypeQueryV1Default) Error() string {
	return fmt.Sprintf("[GET /iocs/queries/ioc-types/v1][%d] ioc_type.query.v1 default  %+v", o._statusCode, o.Payload)
}

func (o *IocTypeQueryV1Default) String() string {
	return fmt.Sprintf("[GET /iocs/queries/ioc-types/v1][%d] ioc_type.query.v1 default  %+v", o._statusCode, o.Payload)
}

func (o *IocTypeQueryV1Default) GetPayload() *models.APIIndicatorQueryRespV1 {
	return o.Payload
}

func (o *IocTypeQueryV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIIndicatorQueryRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}