// Code generated by go-swagger; DO NOT EDIT.

package workflows

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

// V1ExecutionResultsV1Reader is a Reader for the V1ExecutionResultsV1 structure.
type V1ExecutionResultsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ExecutionResultsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ExecutionResultsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1ExecutionResultsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1ExecutionResultsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1ExecutionResultsV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewV1ExecutionResultsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1ExecutionResultsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/entities/execution-results/v1] v1.execution-results.v1", response, response.Code())
	}
}

// NewV1ExecutionResultsV1OK creates a V1ExecutionResultsV1OK with default headers values
func NewV1ExecutionResultsV1OK() *V1ExecutionResultsV1OK {
	return &V1ExecutionResultsV1OK{}
}

/*
V1ExecutionResultsV1OK describes a response with status code 200, with default header values.

OK
*/
type V1ExecutionResultsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIExecutionResultsResponse
}

// IsSuccess returns true when this v1 execution results v1 o k response has a 2xx status code
func (o *V1ExecutionResultsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 execution results v1 o k response has a 3xx status code
func (o *V1ExecutionResultsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 execution results v1 o k response has a 4xx status code
func (o *V1ExecutionResultsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 execution results v1 o k response has a 5xx status code
func (o *V1ExecutionResultsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 execution results v1 o k response a status code equal to that given
func (o *V1ExecutionResultsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 execution results v1 o k response
func (o *V1ExecutionResultsV1OK) Code() int {
	return 200
}

func (o *V1ExecutionResultsV1OK) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1OK  %+v", 200, o.Payload)
}

func (o *V1ExecutionResultsV1OK) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1OK  %+v", 200, o.Payload)
}

func (o *V1ExecutionResultsV1OK) GetPayload() *models.APIExecutionResultsResponse {
	return o.Payload
}

func (o *V1ExecutionResultsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIExecutionResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1ExecutionResultsV1BadRequest creates a V1ExecutionResultsV1BadRequest with default headers values
func NewV1ExecutionResultsV1BadRequest() *V1ExecutionResultsV1BadRequest {
	return &V1ExecutionResultsV1BadRequest{}
}

/*
V1ExecutionResultsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1ExecutionResultsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIExecutionResultsResponse
}

// IsSuccess returns true when this v1 execution results v1 bad request response has a 2xx status code
func (o *V1ExecutionResultsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 execution results v1 bad request response has a 3xx status code
func (o *V1ExecutionResultsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 execution results v1 bad request response has a 4xx status code
func (o *V1ExecutionResultsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 execution results v1 bad request response has a 5xx status code
func (o *V1ExecutionResultsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 execution results v1 bad request response a status code equal to that given
func (o *V1ExecutionResultsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 execution results v1 bad request response
func (o *V1ExecutionResultsV1BadRequest) Code() int {
	return 400
}

func (o *V1ExecutionResultsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1BadRequest  %+v", 400, o.Payload)
}

func (o *V1ExecutionResultsV1BadRequest) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1BadRequest  %+v", 400, o.Payload)
}

func (o *V1ExecutionResultsV1BadRequest) GetPayload() *models.APIExecutionResultsResponse {
	return o.Payload
}

func (o *V1ExecutionResultsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIExecutionResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1ExecutionResultsV1Forbidden creates a V1ExecutionResultsV1Forbidden with default headers values
func NewV1ExecutionResultsV1Forbidden() *V1ExecutionResultsV1Forbidden {
	return &V1ExecutionResultsV1Forbidden{}
}

/*
V1ExecutionResultsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1ExecutionResultsV1Forbidden struct {

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

// IsSuccess returns true when this v1 execution results v1 forbidden response has a 2xx status code
func (o *V1ExecutionResultsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 execution results v1 forbidden response has a 3xx status code
func (o *V1ExecutionResultsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 execution results v1 forbidden response has a 4xx status code
func (o *V1ExecutionResultsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 execution results v1 forbidden response has a 5xx status code
func (o *V1ExecutionResultsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 execution results v1 forbidden response a status code equal to that given
func (o *V1ExecutionResultsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 execution results v1 forbidden response
func (o *V1ExecutionResultsV1Forbidden) Code() int {
	return 403
}

func (o *V1ExecutionResultsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1Forbidden  %+v", 403, o.Payload)
}

func (o *V1ExecutionResultsV1Forbidden) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1Forbidden  %+v", 403, o.Payload)
}

func (o *V1ExecutionResultsV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *V1ExecutionResultsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewV1ExecutionResultsV1NotFound creates a V1ExecutionResultsV1NotFound with default headers values
func NewV1ExecutionResultsV1NotFound() *V1ExecutionResultsV1NotFound {
	return &V1ExecutionResultsV1NotFound{}
}

/*
V1ExecutionResultsV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1ExecutionResultsV1NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIExecutionResultsResponse
}

// IsSuccess returns true when this v1 execution results v1 not found response has a 2xx status code
func (o *V1ExecutionResultsV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 execution results v1 not found response has a 3xx status code
func (o *V1ExecutionResultsV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 execution results v1 not found response has a 4xx status code
func (o *V1ExecutionResultsV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 execution results v1 not found response has a 5xx status code
func (o *V1ExecutionResultsV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 execution results v1 not found response a status code equal to that given
func (o *V1ExecutionResultsV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 execution results v1 not found response
func (o *V1ExecutionResultsV1NotFound) Code() int {
	return 404
}

func (o *V1ExecutionResultsV1NotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1NotFound  %+v", 404, o.Payload)
}

func (o *V1ExecutionResultsV1NotFound) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1NotFound  %+v", 404, o.Payload)
}

func (o *V1ExecutionResultsV1NotFound) GetPayload() *models.APIExecutionResultsResponse {
	return o.Payload
}

func (o *V1ExecutionResultsV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIExecutionResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1ExecutionResultsV1TooManyRequests creates a V1ExecutionResultsV1TooManyRequests with default headers values
func NewV1ExecutionResultsV1TooManyRequests() *V1ExecutionResultsV1TooManyRequests {
	return &V1ExecutionResultsV1TooManyRequests{}
}

/*
V1ExecutionResultsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type V1ExecutionResultsV1TooManyRequests struct {

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

// IsSuccess returns true when this v1 execution results v1 too many requests response has a 2xx status code
func (o *V1ExecutionResultsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 execution results v1 too many requests response has a 3xx status code
func (o *V1ExecutionResultsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 execution results v1 too many requests response has a 4xx status code
func (o *V1ExecutionResultsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 execution results v1 too many requests response has a 5xx status code
func (o *V1ExecutionResultsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 execution results v1 too many requests response a status code equal to that given
func (o *V1ExecutionResultsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the v1 execution results v1 too many requests response
func (o *V1ExecutionResultsV1TooManyRequests) Code() int {
	return 429
}

func (o *V1ExecutionResultsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *V1ExecutionResultsV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *V1ExecutionResultsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *V1ExecutionResultsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewV1ExecutionResultsV1InternalServerError creates a V1ExecutionResultsV1InternalServerError with default headers values
func NewV1ExecutionResultsV1InternalServerError() *V1ExecutionResultsV1InternalServerError {
	return &V1ExecutionResultsV1InternalServerError{}
}

/*
V1ExecutionResultsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1ExecutionResultsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIExecutionResultsResponse
}

// IsSuccess returns true when this v1 execution results v1 internal server error response has a 2xx status code
func (o *V1ExecutionResultsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 execution results v1 internal server error response has a 3xx status code
func (o *V1ExecutionResultsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 execution results v1 internal server error response has a 4xx status code
func (o *V1ExecutionResultsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 execution results v1 internal server error response has a 5xx status code
func (o *V1ExecutionResultsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 execution results v1 internal server error response a status code equal to that given
func (o *V1ExecutionResultsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 execution results v1 internal server error response
func (o *V1ExecutionResultsV1InternalServerError) Code() int {
	return 500
}

func (o *V1ExecutionResultsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *V1ExecutionResultsV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] v1ExecutionResultsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *V1ExecutionResultsV1InternalServerError) GetPayload() *models.APIExecutionResultsResponse {
	return o.Payload
}

func (o *V1ExecutionResultsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIExecutionResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
