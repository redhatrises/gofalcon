// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// PostAggregatesAlertsV1Reader is a Reader for the PostAggregatesAlertsV1 structure.
type PostAggregatesAlertsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAggregatesAlertsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAggregatesAlertsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAggregatesAlertsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAggregatesAlertsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAggregatesAlertsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAggregatesAlertsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostAggregatesAlertsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAggregatesAlertsV1OK creates a PostAggregatesAlertsV1OK with default headers values
func NewPostAggregatesAlertsV1OK() *PostAggregatesAlertsV1OK {
	return &PostAggregatesAlertsV1OK{}
}

/*
PostAggregatesAlertsV1OK describes a response with status code 200, with default header values.

OK
*/
type PostAggregatesAlertsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this post aggregates alerts v1 o k response has a 2xx status code
func (o *PostAggregatesAlertsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post aggregates alerts v1 o k response has a 3xx status code
func (o *PostAggregatesAlertsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post aggregates alerts v1 o k response has a 4xx status code
func (o *PostAggregatesAlertsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post aggregates alerts v1 o k response has a 5xx status code
func (o *PostAggregatesAlertsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this post aggregates alerts v1 o k response a status code equal to that given
func (o *PostAggregatesAlertsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post aggregates alerts v1 o k response
func (o *PostAggregatesAlertsV1OK) Code() int {
	return 200
}

func (o *PostAggregatesAlertsV1OK) Error() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1OK  %+v", 200, o.Payload)
}

func (o *PostAggregatesAlertsV1OK) String() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1OK  %+v", 200, o.Payload)
}

func (o *PostAggregatesAlertsV1OK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *PostAggregatesAlertsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAggregatesAlertsV1BadRequest creates a PostAggregatesAlertsV1BadRequest with default headers values
func NewPostAggregatesAlertsV1BadRequest() *PostAggregatesAlertsV1BadRequest {
	return &PostAggregatesAlertsV1BadRequest{}
}

/*
PostAggregatesAlertsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAggregatesAlertsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this post aggregates alerts v1 bad request response has a 2xx status code
func (o *PostAggregatesAlertsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post aggregates alerts v1 bad request response has a 3xx status code
func (o *PostAggregatesAlertsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post aggregates alerts v1 bad request response has a 4xx status code
func (o *PostAggregatesAlertsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post aggregates alerts v1 bad request response has a 5xx status code
func (o *PostAggregatesAlertsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post aggregates alerts v1 bad request response a status code equal to that given
func (o *PostAggregatesAlertsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post aggregates alerts v1 bad request response
func (o *PostAggregatesAlertsV1BadRequest) Code() int {
	return 400
}

func (o *PostAggregatesAlertsV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1BadRequest  %+v", 400, o.Payload)
}

func (o *PostAggregatesAlertsV1BadRequest) String() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1BadRequest  %+v", 400, o.Payload)
}

func (o *PostAggregatesAlertsV1BadRequest) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *PostAggregatesAlertsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAggregatesAlertsV1Forbidden creates a PostAggregatesAlertsV1Forbidden with default headers values
func NewPostAggregatesAlertsV1Forbidden() *PostAggregatesAlertsV1Forbidden {
	return &PostAggregatesAlertsV1Forbidden{}
}

/*
PostAggregatesAlertsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAggregatesAlertsV1Forbidden struct {

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

// IsSuccess returns true when this post aggregates alerts v1 forbidden response has a 2xx status code
func (o *PostAggregatesAlertsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post aggregates alerts v1 forbidden response has a 3xx status code
func (o *PostAggregatesAlertsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post aggregates alerts v1 forbidden response has a 4xx status code
func (o *PostAggregatesAlertsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post aggregates alerts v1 forbidden response has a 5xx status code
func (o *PostAggregatesAlertsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post aggregates alerts v1 forbidden response a status code equal to that given
func (o *PostAggregatesAlertsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post aggregates alerts v1 forbidden response
func (o *PostAggregatesAlertsV1Forbidden) Code() int {
	return 403
}

func (o *PostAggregatesAlertsV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1Forbidden  %+v", 403, o.Payload)
}

func (o *PostAggregatesAlertsV1Forbidden) String() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1Forbidden  %+v", 403, o.Payload)
}

func (o *PostAggregatesAlertsV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PostAggregatesAlertsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostAggregatesAlertsV1TooManyRequests creates a PostAggregatesAlertsV1TooManyRequests with default headers values
func NewPostAggregatesAlertsV1TooManyRequests() *PostAggregatesAlertsV1TooManyRequests {
	return &PostAggregatesAlertsV1TooManyRequests{}
}

/*
PostAggregatesAlertsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PostAggregatesAlertsV1TooManyRequests struct {

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

// IsSuccess returns true when this post aggregates alerts v1 too many requests response has a 2xx status code
func (o *PostAggregatesAlertsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post aggregates alerts v1 too many requests response has a 3xx status code
func (o *PostAggregatesAlertsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post aggregates alerts v1 too many requests response has a 4xx status code
func (o *PostAggregatesAlertsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post aggregates alerts v1 too many requests response has a 5xx status code
func (o *PostAggregatesAlertsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post aggregates alerts v1 too many requests response a status code equal to that given
func (o *PostAggregatesAlertsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the post aggregates alerts v1 too many requests response
func (o *PostAggregatesAlertsV1TooManyRequests) Code() int {
	return 429
}

func (o *PostAggregatesAlertsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAggregatesAlertsV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAggregatesAlertsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PostAggregatesAlertsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostAggregatesAlertsV1InternalServerError creates a PostAggregatesAlertsV1InternalServerError with default headers values
func NewPostAggregatesAlertsV1InternalServerError() *PostAggregatesAlertsV1InternalServerError {
	return &PostAggregatesAlertsV1InternalServerError{}
}

/*
PostAggregatesAlertsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAggregatesAlertsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this post aggregates alerts v1 internal server error response has a 2xx status code
func (o *PostAggregatesAlertsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post aggregates alerts v1 internal server error response has a 3xx status code
func (o *PostAggregatesAlertsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post aggregates alerts v1 internal server error response has a 4xx status code
func (o *PostAggregatesAlertsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post aggregates alerts v1 internal server error response has a 5xx status code
func (o *PostAggregatesAlertsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post aggregates alerts v1 internal server error response a status code equal to that given
func (o *PostAggregatesAlertsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post aggregates alerts v1 internal server error response
func (o *PostAggregatesAlertsV1InternalServerError) Code() int {
	return 500
}

func (o *PostAggregatesAlertsV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PostAggregatesAlertsV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] postAggregatesAlertsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PostAggregatesAlertsV1InternalServerError) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *PostAggregatesAlertsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAggregatesAlertsV1Default creates a PostAggregatesAlertsV1Default with default headers values
func NewPostAggregatesAlertsV1Default(code int) *PostAggregatesAlertsV1Default {
	return &PostAggregatesAlertsV1Default{
		_statusCode: code,
	}
}

/*
PostAggregatesAlertsV1Default describes a response with status code -1, with default header values.

OK
*/
type PostAggregatesAlertsV1Default struct {
	_statusCode int

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this post aggregates alerts v1 default response has a 2xx status code
func (o *PostAggregatesAlertsV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post aggregates alerts v1 default response has a 3xx status code
func (o *PostAggregatesAlertsV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post aggregates alerts v1 default response has a 4xx status code
func (o *PostAggregatesAlertsV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post aggregates alerts v1 default response has a 5xx status code
func (o *PostAggregatesAlertsV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post aggregates alerts v1 default response a status code equal to that given
func (o *PostAggregatesAlertsV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post aggregates alerts v1 default response
func (o *PostAggregatesAlertsV1Default) Code() int {
	return o._statusCode
}

func (o *PostAggregatesAlertsV1Default) Error() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] PostAggregatesAlertsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PostAggregatesAlertsV1Default) String() string {
	return fmt.Sprintf("[POST /alerts/aggregates/alerts/v1][%d] PostAggregatesAlertsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PostAggregatesAlertsV1Default) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *PostAggregatesAlertsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}