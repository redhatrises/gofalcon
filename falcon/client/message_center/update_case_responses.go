// Code generated by go-swagger; DO NOT EDIT.

package message_center

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

// UpdateCaseReader is a Reader for the UpdateCase structure.
type UpdateCaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCaseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCaseTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateCaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCaseOK creates a UpdateCaseOK with default headers values
func NewUpdateCaseOK() *UpdateCaseOK {
	return &UpdateCaseOK{}
}

/*
UpdateCaseOK describes a response with status code 200, with default header values.

OK
*/
type UpdateCaseOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyAffectedEntities
}

// IsSuccess returns true when this update case o k response has a 2xx status code
func (o *UpdateCaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update case o k response has a 3xx status code
func (o *UpdateCaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update case o k response has a 4xx status code
func (o *UpdateCaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update case o k response has a 5xx status code
func (o *UpdateCaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update case o k response a status code equal to that given
func (o *UpdateCaseOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateCaseOK) Error() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseOK  %+v", 200, o.Payload)
}

func (o *UpdateCaseOK) String() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseOK  %+v", 200, o.Payload)
}

func (o *UpdateCaseOK) GetPayload() *models.MsaReplyAffectedEntities {
	return o.Payload
}

func (o *UpdateCaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyAffectedEntities)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCaseBadRequest creates a UpdateCaseBadRequest with default headers values
func NewUpdateCaseBadRequest() *UpdateCaseBadRequest {
	return &UpdateCaseBadRequest{}
}

/*
UpdateCaseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateCaseBadRequest struct {

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

// IsSuccess returns true when this update case bad request response has a 2xx status code
func (o *UpdateCaseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update case bad request response has a 3xx status code
func (o *UpdateCaseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update case bad request response has a 4xx status code
func (o *UpdateCaseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update case bad request response has a 5xx status code
func (o *UpdateCaseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update case bad request response a status code equal to that given
func (o *UpdateCaseBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateCaseBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCaseBadRequest) String() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCaseBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCaseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCaseForbidden creates a UpdateCaseForbidden with default headers values
func NewUpdateCaseForbidden() *UpdateCaseForbidden {
	return &UpdateCaseForbidden{}
}

/*
UpdateCaseForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateCaseForbidden struct {

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

// IsSuccess returns true when this update case forbidden response has a 2xx status code
func (o *UpdateCaseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update case forbidden response has a 3xx status code
func (o *UpdateCaseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update case forbidden response has a 4xx status code
func (o *UpdateCaseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update case forbidden response has a 5xx status code
func (o *UpdateCaseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update case forbidden response a status code equal to that given
func (o *UpdateCaseForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateCaseForbidden) Error() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCaseForbidden) String() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCaseForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCaseTooManyRequests creates a UpdateCaseTooManyRequests with default headers values
func NewUpdateCaseTooManyRequests() *UpdateCaseTooManyRequests {
	return &UpdateCaseTooManyRequests{}
}

/*
UpdateCaseTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateCaseTooManyRequests struct {

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

// IsSuccess returns true when this update case too many requests response has a 2xx status code
func (o *UpdateCaseTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update case too many requests response has a 3xx status code
func (o *UpdateCaseTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update case too many requests response has a 4xx status code
func (o *UpdateCaseTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update case too many requests response has a 5xx status code
func (o *UpdateCaseTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update case too many requests response a status code equal to that given
func (o *UpdateCaseTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateCaseTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateCaseTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateCaseTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCaseTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCaseInternalServerError creates a UpdateCaseInternalServerError with default headers values
func NewUpdateCaseInternalServerError() *UpdateCaseInternalServerError {
	return &UpdateCaseInternalServerError{}
}

/*
UpdateCaseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateCaseInternalServerError struct {

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

// IsSuccess returns true when this update case internal server error response has a 2xx status code
func (o *UpdateCaseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update case internal server error response has a 3xx status code
func (o *UpdateCaseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update case internal server error response has a 4xx status code
func (o *UpdateCaseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update case internal server error response has a 5xx status code
func (o *UpdateCaseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update case internal server error response a status code equal to that given
func (o *UpdateCaseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateCaseInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCaseInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] updateCaseInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCaseInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCaseDefault creates a UpdateCaseDefault with default headers values
func NewUpdateCaseDefault(code int) *UpdateCaseDefault {
	return &UpdateCaseDefault{
		_statusCode: code,
	}
}

/*
UpdateCaseDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateCaseDefault struct {
	_statusCode int

	Payload *models.MsaReplyAffectedEntities
}

// Code gets the status code for the update case default response
func (o *UpdateCaseDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update case default response has a 2xx status code
func (o *UpdateCaseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update case default response has a 3xx status code
func (o *UpdateCaseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update case default response has a 4xx status code
func (o *UpdateCaseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update case default response has a 5xx status code
func (o *UpdateCaseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update case default response a status code equal to that given
func (o *UpdateCaseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateCaseDefault) Error() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] UpdateCase default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCaseDefault) String() string {
	return fmt.Sprintf("[PATCH /message-center/entities/case/v1][%d] UpdateCase default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCaseDefault) GetPayload() *models.MsaReplyAffectedEntities {
	return o.Payload
}

func (o *UpdateCaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyAffectedEntities)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
