// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// UpdatePreventionPoliciesReader is a Reader for the UpdatePreventionPolicies structure.
type UpdatePreventionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePreventionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePreventionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePreventionPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePreventionPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePreventionPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdatePreventionPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePreventionPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdatePreventionPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePreventionPoliciesOK creates a UpdatePreventionPoliciesOK with default headers values
func NewUpdatePreventionPoliciesOK() *UpdatePreventionPoliciesOK {
	return &UpdatePreventionPoliciesOK{}
}

/*
UpdatePreventionPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type UpdatePreventionPoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this update prevention policies o k response has a 2xx status code
func (o *UpdatePreventionPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update prevention policies o k response has a 3xx status code
func (o *UpdatePreventionPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update prevention policies o k response has a 4xx status code
func (o *UpdatePreventionPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update prevention policies o k response has a 5xx status code
func (o *UpdatePreventionPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update prevention policies o k response a status code equal to that given
func (o *UpdatePreventionPoliciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdatePreventionPoliciesOK) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *UpdatePreventionPoliciesOK) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *UpdatePreventionPoliciesOK) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *UpdatePreventionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePreventionPoliciesBadRequest creates a UpdatePreventionPoliciesBadRequest with default headers values
func NewUpdatePreventionPoliciesBadRequest() *UpdatePreventionPoliciesBadRequest {
	return &UpdatePreventionPoliciesBadRequest{}
}

/*
UpdatePreventionPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdatePreventionPoliciesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this update prevention policies bad request response has a 2xx status code
func (o *UpdatePreventionPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update prevention policies bad request response has a 3xx status code
func (o *UpdatePreventionPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update prevention policies bad request response has a 4xx status code
func (o *UpdatePreventionPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update prevention policies bad request response has a 5xx status code
func (o *UpdatePreventionPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update prevention policies bad request response a status code equal to that given
func (o *UpdatePreventionPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdatePreventionPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePreventionPoliciesBadRequest) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePreventionPoliciesBadRequest) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *UpdatePreventionPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePreventionPoliciesForbidden creates a UpdatePreventionPoliciesForbidden with default headers values
func NewUpdatePreventionPoliciesForbidden() *UpdatePreventionPoliciesForbidden {
	return &UpdatePreventionPoliciesForbidden{}
}

/*
UpdatePreventionPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdatePreventionPoliciesForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this update prevention policies forbidden response has a 2xx status code
func (o *UpdatePreventionPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update prevention policies forbidden response has a 3xx status code
func (o *UpdatePreventionPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update prevention policies forbidden response has a 4xx status code
func (o *UpdatePreventionPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update prevention policies forbidden response has a 5xx status code
func (o *UpdatePreventionPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update prevention policies forbidden response a status code equal to that given
func (o *UpdatePreventionPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdatePreventionPoliciesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *UpdatePreventionPoliciesForbidden) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *UpdatePreventionPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdatePreventionPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePreventionPoliciesNotFound creates a UpdatePreventionPoliciesNotFound with default headers values
func NewUpdatePreventionPoliciesNotFound() *UpdatePreventionPoliciesNotFound {
	return &UpdatePreventionPoliciesNotFound{}
}

/*
UpdatePreventionPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdatePreventionPoliciesNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this update prevention policies not found response has a 2xx status code
func (o *UpdatePreventionPoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update prevention policies not found response has a 3xx status code
func (o *UpdatePreventionPoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update prevention policies not found response has a 4xx status code
func (o *UpdatePreventionPoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update prevention policies not found response has a 5xx status code
func (o *UpdatePreventionPoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update prevention policies not found response a status code equal to that given
func (o *UpdatePreventionPoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdatePreventionPoliciesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *UpdatePreventionPoliciesNotFound) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *UpdatePreventionPoliciesNotFound) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *UpdatePreventionPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePreventionPoliciesTooManyRequests creates a UpdatePreventionPoliciesTooManyRequests with default headers values
func NewUpdatePreventionPoliciesTooManyRequests() *UpdatePreventionPoliciesTooManyRequests {
	return &UpdatePreventionPoliciesTooManyRequests{}
}

/*
UpdatePreventionPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdatePreventionPoliciesTooManyRequests struct {

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

// IsSuccess returns true when this update prevention policies too many requests response has a 2xx status code
func (o *UpdatePreventionPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update prevention policies too many requests response has a 3xx status code
func (o *UpdatePreventionPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update prevention policies too many requests response has a 4xx status code
func (o *UpdatePreventionPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update prevention policies too many requests response has a 5xx status code
func (o *UpdatePreventionPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update prevention policies too many requests response a status code equal to that given
func (o *UpdatePreventionPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdatePreventionPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdatePreventionPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdatePreventionPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdatePreventionPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdatePreventionPoliciesInternalServerError creates a UpdatePreventionPoliciesInternalServerError with default headers values
func NewUpdatePreventionPoliciesInternalServerError() *UpdatePreventionPoliciesInternalServerError {
	return &UpdatePreventionPoliciesInternalServerError{}
}

/*
UpdatePreventionPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdatePreventionPoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this update prevention policies internal server error response has a 2xx status code
func (o *UpdatePreventionPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update prevention policies internal server error response has a 3xx status code
func (o *UpdatePreventionPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update prevention policies internal server error response has a 4xx status code
func (o *UpdatePreventionPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update prevention policies internal server error response has a 5xx status code
func (o *UpdatePreventionPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update prevention policies internal server error response a status code equal to that given
func (o *UpdatePreventionPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdatePreventionPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdatePreventionPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdatePreventionPoliciesInternalServerError) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *UpdatePreventionPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePreventionPoliciesDefault creates a UpdatePreventionPoliciesDefault with default headers values
func NewUpdatePreventionPoliciesDefault(code int) *UpdatePreventionPoliciesDefault {
	return &UpdatePreventionPoliciesDefault{
		_statusCode: code,
	}
}

/*
UpdatePreventionPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type UpdatePreventionPoliciesDefault struct {
	_statusCode int

	Payload *models.ResponsesPreventionPoliciesV1
}

// Code gets the status code for the update prevention policies default response
func (o *UpdatePreventionPoliciesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update prevention policies default response has a 2xx status code
func (o *UpdatePreventionPoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update prevention policies default response has a 3xx status code
func (o *UpdatePreventionPoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update prevention policies default response has a 4xx status code
func (o *UpdatePreventionPoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update prevention policies default response has a 5xx status code
func (o *UpdatePreventionPoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update prevention policies default response a status code equal to that given
func (o *UpdatePreventionPoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdatePreventionPoliciesDefault) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *UpdatePreventionPoliciesDefault) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/prevention/v1][%d] updatePreventionPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *UpdatePreventionPoliciesDefault) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *UpdatePreventionPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
