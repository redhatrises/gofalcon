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

// RTRDeleteFileV2Reader is a Reader for the RTRDeleteFileV2 structure.
type RTRDeleteFileV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRDeleteFileV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRTRDeleteFileV2NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRDeleteFileV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRDeleteFileV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRDeleteFileV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRDeleteFileV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRTRDeleteFileV2NoContent creates a RTRDeleteFileV2NoContent with default headers values
func NewRTRDeleteFileV2NoContent() *RTRDeleteFileV2NoContent {
	return &RTRDeleteFileV2NoContent{}
}

/*
RTRDeleteFileV2NoContent describes a response with status code 204, with default header values.

No Content
*/
type RTRDeleteFileV2NoContent struct {

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

// IsSuccess returns true when this r t r delete file v2 no content response has a 2xx status code
func (o *RTRDeleteFileV2NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r delete file v2 no content response has a 3xx status code
func (o *RTRDeleteFileV2NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file v2 no content response has a 4xx status code
func (o *RTRDeleteFileV2NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r delete file v2 no content response has a 5xx status code
func (o *RTRDeleteFileV2NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file v2 no content response a status code equal to that given
func (o *RTRDeleteFileV2NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the r t r delete file v2 no content response
func (o *RTRDeleteFileV2NoContent) Code() int {
	return 204
}

func (o *RTRDeleteFileV2NoContent) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2NoContent  %+v", 204, o.Payload)
}

func (o *RTRDeleteFileV2NoContent) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2NoContent  %+v", 204, o.Payload)
}

func (o *RTRDeleteFileV2NoContent) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteFileV2NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteFileV2BadRequest creates a RTRDeleteFileV2BadRequest with default headers values
func NewRTRDeleteFileV2BadRequest() *RTRDeleteFileV2BadRequest {
	return &RTRDeleteFileV2BadRequest{}
}

/*
RTRDeleteFileV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRDeleteFileV2BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r delete file v2 bad request response has a 2xx status code
func (o *RTRDeleteFileV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file v2 bad request response has a 3xx status code
func (o *RTRDeleteFileV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file v2 bad request response has a 4xx status code
func (o *RTRDeleteFileV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete file v2 bad request response has a 5xx status code
func (o *RTRDeleteFileV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file v2 bad request response a status code equal to that given
func (o *RTRDeleteFileV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r delete file v2 bad request response
func (o *RTRDeleteFileV2BadRequest) Code() int {
	return 400
}

func (o *RTRDeleteFileV2BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2BadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeleteFileV2BadRequest) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2BadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeleteFileV2BadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRDeleteFileV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRDeleteFileV2Forbidden creates a RTRDeleteFileV2Forbidden with default headers values
func NewRTRDeleteFileV2Forbidden() *RTRDeleteFileV2Forbidden {
	return &RTRDeleteFileV2Forbidden{}
}

/*
RTRDeleteFileV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRDeleteFileV2Forbidden struct {

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

// IsSuccess returns true when this r t r delete file v2 forbidden response has a 2xx status code
func (o *RTRDeleteFileV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file v2 forbidden response has a 3xx status code
func (o *RTRDeleteFileV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file v2 forbidden response has a 4xx status code
func (o *RTRDeleteFileV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete file v2 forbidden response has a 5xx status code
func (o *RTRDeleteFileV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file v2 forbidden response a status code equal to that given
func (o *RTRDeleteFileV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r delete file v2 forbidden response
func (o *RTRDeleteFileV2Forbidden) Code() int {
	return 403
}

func (o *RTRDeleteFileV2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2Forbidden  %+v", 403, o.Payload)
}

func (o *RTRDeleteFileV2Forbidden) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2Forbidden  %+v", 403, o.Payload)
}

func (o *RTRDeleteFileV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteFileV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteFileV2NotFound creates a RTRDeleteFileV2NotFound with default headers values
func NewRTRDeleteFileV2NotFound() *RTRDeleteFileV2NotFound {
	return &RTRDeleteFileV2NotFound{}
}

/*
RTRDeleteFileV2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRDeleteFileV2NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r delete file v2 not found response has a 2xx status code
func (o *RTRDeleteFileV2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file v2 not found response has a 3xx status code
func (o *RTRDeleteFileV2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file v2 not found response has a 4xx status code
func (o *RTRDeleteFileV2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete file v2 not found response has a 5xx status code
func (o *RTRDeleteFileV2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file v2 not found response a status code equal to that given
func (o *RTRDeleteFileV2NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r delete file v2 not found response
func (o *RTRDeleteFileV2NotFound) Code() int {
	return 404
}

func (o *RTRDeleteFileV2NotFound) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2NotFound  %+v", 404, o.Payload)
}

func (o *RTRDeleteFileV2NotFound) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2NotFound  %+v", 404, o.Payload)
}

func (o *RTRDeleteFileV2NotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRDeleteFileV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRDeleteFileV2TooManyRequests creates a RTRDeleteFileV2TooManyRequests with default headers values
func NewRTRDeleteFileV2TooManyRequests() *RTRDeleteFileV2TooManyRequests {
	return &RTRDeleteFileV2TooManyRequests{}
}

/*
RTRDeleteFileV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRDeleteFileV2TooManyRequests struct {

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

// IsSuccess returns true when this r t r delete file v2 too many requests response has a 2xx status code
func (o *RTRDeleteFileV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file v2 too many requests response has a 3xx status code
func (o *RTRDeleteFileV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file v2 too many requests response has a 4xx status code
func (o *RTRDeleteFileV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete file v2 too many requests response has a 5xx status code
func (o *RTRDeleteFileV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file v2 too many requests response a status code equal to that given
func (o *RTRDeleteFileV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r delete file v2 too many requests response
func (o *RTRDeleteFileV2TooManyRequests) Code() int {
	return 429
}

func (o *RTRDeleteFileV2TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeleteFileV2TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v2][%d] rTRDeleteFileV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeleteFileV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteFileV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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