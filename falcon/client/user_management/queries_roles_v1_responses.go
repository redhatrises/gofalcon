// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// QueriesRolesV1Reader is a Reader for the QueriesRolesV1 structure.
type QueriesRolesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueriesRolesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueriesRolesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueriesRolesV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueriesRolesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueriesRolesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueriesRolesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueriesRolesV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueriesRolesV1OK creates a QueriesRolesV1OK with default headers values
func NewQueriesRolesV1OK() *QueriesRolesV1OK {
	return &QueriesRolesV1OK{}
}

/*
QueriesRolesV1OK describes a response with status code 200, with default header values.

OK
*/
type QueriesRolesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this queries roles v1 o k response has a 2xx status code
func (o *QueriesRolesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this queries roles v1 o k response has a 3xx status code
func (o *QueriesRolesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries roles v1 o k response has a 4xx status code
func (o *QueriesRolesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this queries roles v1 o k response has a 5xx status code
func (o *QueriesRolesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this queries roles v1 o k response a status code equal to that given
func (o *QueriesRolesV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *QueriesRolesV1OK) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1OK  %+v", 200, o.Payload)
}

func (o *QueriesRolesV1OK) String() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1OK  %+v", 200, o.Payload)
}

func (o *QueriesRolesV1OK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueriesRolesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueriesRolesV1BadRequest creates a QueriesRolesV1BadRequest with default headers values
func NewQueriesRolesV1BadRequest() *QueriesRolesV1BadRequest {
	return &QueriesRolesV1BadRequest{}
}

/*
QueriesRolesV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueriesRolesV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this queries roles v1 bad request response has a 2xx status code
func (o *QueriesRolesV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this queries roles v1 bad request response has a 3xx status code
func (o *QueriesRolesV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries roles v1 bad request response has a 4xx status code
func (o *QueriesRolesV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this queries roles v1 bad request response has a 5xx status code
func (o *QueriesRolesV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this queries roles v1 bad request response a status code equal to that given
func (o *QueriesRolesV1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *QueriesRolesV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1BadRequest  %+v", 400, o.Payload)
}

func (o *QueriesRolesV1BadRequest) String() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1BadRequest  %+v", 400, o.Payload)
}

func (o *QueriesRolesV1BadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueriesRolesV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueriesRolesV1Forbidden creates a QueriesRolesV1Forbidden with default headers values
func NewQueriesRolesV1Forbidden() *QueriesRolesV1Forbidden {
	return &QueriesRolesV1Forbidden{}
}

/*
QueriesRolesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueriesRolesV1Forbidden struct {

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

// IsSuccess returns true when this queries roles v1 forbidden response has a 2xx status code
func (o *QueriesRolesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this queries roles v1 forbidden response has a 3xx status code
func (o *QueriesRolesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries roles v1 forbidden response has a 4xx status code
func (o *QueriesRolesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this queries roles v1 forbidden response has a 5xx status code
func (o *QueriesRolesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this queries roles v1 forbidden response a status code equal to that given
func (o *QueriesRolesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueriesRolesV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueriesRolesV1Forbidden) String() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueriesRolesV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueriesRolesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueriesRolesV1TooManyRequests creates a QueriesRolesV1TooManyRequests with default headers values
func NewQueriesRolesV1TooManyRequests() *QueriesRolesV1TooManyRequests {
	return &QueriesRolesV1TooManyRequests{}
}

/*
QueriesRolesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueriesRolesV1TooManyRequests struct {

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

// IsSuccess returns true when this queries roles v1 too many requests response has a 2xx status code
func (o *QueriesRolesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this queries roles v1 too many requests response has a 3xx status code
func (o *QueriesRolesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries roles v1 too many requests response has a 4xx status code
func (o *QueriesRolesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this queries roles v1 too many requests response has a 5xx status code
func (o *QueriesRolesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this queries roles v1 too many requests response a status code equal to that given
func (o *QueriesRolesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueriesRolesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueriesRolesV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueriesRolesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueriesRolesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueriesRolesV1InternalServerError creates a QueriesRolesV1InternalServerError with default headers values
func NewQueriesRolesV1InternalServerError() *QueriesRolesV1InternalServerError {
	return &QueriesRolesV1InternalServerError{}
}

/*
QueriesRolesV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueriesRolesV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this queries roles v1 internal server error response has a 2xx status code
func (o *QueriesRolesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this queries roles v1 internal server error response has a 3xx status code
func (o *QueriesRolesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries roles v1 internal server error response has a 4xx status code
func (o *QueriesRolesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this queries roles v1 internal server error response has a 5xx status code
func (o *QueriesRolesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this queries roles v1 internal server error response a status code equal to that given
func (o *QueriesRolesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *QueriesRolesV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *QueriesRolesV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *QueriesRolesV1InternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueriesRolesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueriesRolesV1Default creates a QueriesRolesV1Default with default headers values
func NewQueriesRolesV1Default(code int) *QueriesRolesV1Default {
	return &QueriesRolesV1Default{
		_statusCode: code,
	}
}

/*
QueriesRolesV1Default describes a response with status code -1, with default header values.

OK
*/
type QueriesRolesV1Default struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the queries roles v1 default response
func (o *QueriesRolesV1Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this queries roles v1 default response has a 2xx status code
func (o *QueriesRolesV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this queries roles v1 default response has a 3xx status code
func (o *QueriesRolesV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this queries roles v1 default response has a 4xx status code
func (o *QueriesRolesV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this queries roles v1 default response has a 5xx status code
func (o *QueriesRolesV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this queries roles v1 default response a status code equal to that given
func (o *QueriesRolesV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueriesRolesV1Default) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *QueriesRolesV1Default) String() string {
	return fmt.Sprintf("[GET /user-management/queries/roles/v1][%d] queriesRolesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *QueriesRolesV1Default) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueriesRolesV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
