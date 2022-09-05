// Code generated by go-swagger; DO NOT EDIT.

package sensor_visibility_exclusions

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

// GetSensorVisibilityExclusionsV1Reader is a Reader for the GetSensorVisibilityExclusionsV1 structure.
type GetSensorVisibilityExclusionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSensorVisibilityExclusionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSensorVisibilityExclusionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSensorVisibilityExclusionsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSensorVisibilityExclusionsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSensorVisibilityExclusionsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSensorVisibilityExclusionsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSensorVisibilityExclusionsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSensorVisibilityExclusionsV1OK creates a GetSensorVisibilityExclusionsV1OK with default headers values
func NewGetSensorVisibilityExclusionsV1OK() *GetSensorVisibilityExclusionsV1OK {
	return &GetSensorVisibilityExclusionsV1OK{}
}

/*
GetSensorVisibilityExclusionsV1OK describes a response with status code 200, with default header values.

OK
*/
type GetSensorVisibilityExclusionsV1OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSvExclusionRespV1
}

// IsSuccess returns true when this get sensor visibility exclusions v1 o k response has a 2xx status code
func (o *GetSensorVisibilityExclusionsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sensor visibility exclusions v1 o k response has a 3xx status code
func (o *GetSensorVisibilityExclusionsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor visibility exclusions v1 o k response has a 4xx status code
func (o *GetSensorVisibilityExclusionsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sensor visibility exclusions v1 o k response has a 5xx status code
func (o *GetSensorVisibilityExclusionsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor visibility exclusions v1 o k response a status code equal to that given
func (o *GetSensorVisibilityExclusionsV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSensorVisibilityExclusionsV1OK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1OK) String() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1OK) GetPayload() *models.ResponsesSvExclusionRespV1 {
	return o.Payload
}

func (o *GetSensorVisibilityExclusionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSvExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorVisibilityExclusionsV1BadRequest creates a GetSensorVisibilityExclusionsV1BadRequest with default headers values
func NewGetSensorVisibilityExclusionsV1BadRequest() *GetSensorVisibilityExclusionsV1BadRequest {
	return &GetSensorVisibilityExclusionsV1BadRequest{}
}

/*
GetSensorVisibilityExclusionsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSensorVisibilityExclusionsV1BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSvExclusionRespV1
}

// IsSuccess returns true when this get sensor visibility exclusions v1 bad request response has a 2xx status code
func (o *GetSensorVisibilityExclusionsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor visibility exclusions v1 bad request response has a 3xx status code
func (o *GetSensorVisibilityExclusionsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor visibility exclusions v1 bad request response has a 4xx status code
func (o *GetSensorVisibilityExclusionsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor visibility exclusions v1 bad request response has a 5xx status code
func (o *GetSensorVisibilityExclusionsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor visibility exclusions v1 bad request response a status code equal to that given
func (o *GetSensorVisibilityExclusionsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSensorVisibilityExclusionsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1BadRequest) String() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1BadRequest) GetPayload() *models.ResponsesSvExclusionRespV1 {
	return o.Payload
}

func (o *GetSensorVisibilityExclusionsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSvExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorVisibilityExclusionsV1Forbidden creates a GetSensorVisibilityExclusionsV1Forbidden with default headers values
func NewGetSensorVisibilityExclusionsV1Forbidden() *GetSensorVisibilityExclusionsV1Forbidden {
	return &GetSensorVisibilityExclusionsV1Forbidden{}
}

/*
GetSensorVisibilityExclusionsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSensorVisibilityExclusionsV1Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get sensor visibility exclusions v1 forbidden response has a 2xx status code
func (o *GetSensorVisibilityExclusionsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor visibility exclusions v1 forbidden response has a 3xx status code
func (o *GetSensorVisibilityExclusionsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor visibility exclusions v1 forbidden response has a 4xx status code
func (o *GetSensorVisibilityExclusionsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor visibility exclusions v1 forbidden response has a 5xx status code
func (o *GetSensorVisibilityExclusionsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor visibility exclusions v1 forbidden response a status code equal to that given
func (o *GetSensorVisibilityExclusionsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSensorVisibilityExclusionsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1Forbidden) String() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetSensorVisibilityExclusionsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorVisibilityExclusionsV1TooManyRequests creates a GetSensorVisibilityExclusionsV1TooManyRequests with default headers values
func NewGetSensorVisibilityExclusionsV1TooManyRequests() *GetSensorVisibilityExclusionsV1TooManyRequests {
	return &GetSensorVisibilityExclusionsV1TooManyRequests{}
}

/*
GetSensorVisibilityExclusionsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSensorVisibilityExclusionsV1TooManyRequests struct {

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

// IsSuccess returns true when this get sensor visibility exclusions v1 too many requests response has a 2xx status code
func (o *GetSensorVisibilityExclusionsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor visibility exclusions v1 too many requests response has a 3xx status code
func (o *GetSensorVisibilityExclusionsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor visibility exclusions v1 too many requests response has a 4xx status code
func (o *GetSensorVisibilityExclusionsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor visibility exclusions v1 too many requests response has a 5xx status code
func (o *GetSensorVisibilityExclusionsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor visibility exclusions v1 too many requests response a status code equal to that given
func (o *GetSensorVisibilityExclusionsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetSensorVisibilityExclusionsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorVisibilityExclusionsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorVisibilityExclusionsV1InternalServerError creates a GetSensorVisibilityExclusionsV1InternalServerError with default headers values
func NewGetSensorVisibilityExclusionsV1InternalServerError() *GetSensorVisibilityExclusionsV1InternalServerError {
	return &GetSensorVisibilityExclusionsV1InternalServerError{}
}

/*
GetSensorVisibilityExclusionsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSensorVisibilityExclusionsV1InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSvExclusionRespV1
}

// IsSuccess returns true when this get sensor visibility exclusions v1 internal server error response has a 2xx status code
func (o *GetSensorVisibilityExclusionsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor visibility exclusions v1 internal server error response has a 3xx status code
func (o *GetSensorVisibilityExclusionsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor visibility exclusions v1 internal server error response has a 4xx status code
func (o *GetSensorVisibilityExclusionsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sensor visibility exclusions v1 internal server error response has a 5xx status code
func (o *GetSensorVisibilityExclusionsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sensor visibility exclusions v1 internal server error response a status code equal to that given
func (o *GetSensorVisibilityExclusionsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSensorVisibilityExclusionsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1InternalServerError) GetPayload() *models.ResponsesSvExclusionRespV1 {
	return o.Payload
}

func (o *GetSensorVisibilityExclusionsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSvExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorVisibilityExclusionsV1Default creates a GetSensorVisibilityExclusionsV1Default with default headers values
func NewGetSensorVisibilityExclusionsV1Default(code int) *GetSensorVisibilityExclusionsV1Default {
	return &GetSensorVisibilityExclusionsV1Default{
		_statusCode: code,
	}
}

/*
GetSensorVisibilityExclusionsV1Default describes a response with status code -1, with default header values.

OK
*/
type GetSensorVisibilityExclusionsV1Default struct {
	_statusCode int

	Payload *models.ResponsesSvExclusionRespV1
}

// Code gets the status code for the get sensor visibility exclusions v1 default response
func (o *GetSensorVisibilityExclusionsV1Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get sensor visibility exclusions v1 default response has a 2xx status code
func (o *GetSensorVisibilityExclusionsV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get sensor visibility exclusions v1 default response has a 3xx status code
func (o *GetSensorVisibilityExclusionsV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get sensor visibility exclusions v1 default response has a 4xx status code
func (o *GetSensorVisibilityExclusionsV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get sensor visibility exclusions v1 default response has a 5xx status code
func (o *GetSensorVisibilityExclusionsV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get sensor visibility exclusions v1 default response a status code equal to that given
func (o *GetSensorVisibilityExclusionsV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetSensorVisibilityExclusionsV1Default) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1Default) String() string {
	return fmt.Sprintf("[GET /policy/entities/sv-exclusions/v1][%d] getSensorVisibilityExclusionsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetSensorVisibilityExclusionsV1Default) GetPayload() *models.ResponsesSvExclusionRespV1 {
	return o.Payload
}

func (o *GetSensorVisibilityExclusionsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesSvExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
