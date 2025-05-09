// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// DeleteRuleGroupsMixin0Reader is a Reader for the DeleteRuleGroupsMixin0 structure.
type DeleteRuleGroupsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRuleGroupsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRuleGroupsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRuleGroupsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRuleGroupsMixin0NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRuleGroupsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRuleGroupsMixin0InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /ioarules/entities/rule-groups/v1] delete-rule-groupsMixin0", response, response.Code())
	}
}

// NewDeleteRuleGroupsMixin0OK creates a DeleteRuleGroupsMixin0OK with default headers values
func NewDeleteRuleGroupsMixin0OK() *DeleteRuleGroupsMixin0OK {
	return &DeleteRuleGroupsMixin0OK{}
}

/*
DeleteRuleGroupsMixin0OK describes a response with status code 200, with default header values.

OK
*/
type DeleteRuleGroupsMixin0OK struct {

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

// IsSuccess returns true when this delete rule groups mixin0 o k response has a 2xx status code
func (o *DeleteRuleGroupsMixin0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete rule groups mixin0 o k response has a 3xx status code
func (o *DeleteRuleGroupsMixin0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rule groups mixin0 o k response has a 4xx status code
func (o *DeleteRuleGroupsMixin0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rule groups mixin0 o k response has a 5xx status code
func (o *DeleteRuleGroupsMixin0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rule groups mixin0 o k response a status code equal to that given
func (o *DeleteRuleGroupsMixin0OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete rule groups mixin0 o k response
func (o *DeleteRuleGroupsMixin0OK) Code() int {
	return 200
}

func (o *DeleteRuleGroupsMixin0OK) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0OK  %+v", 200, o.Payload)
}

func (o *DeleteRuleGroupsMixin0OK) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0OK  %+v", 200, o.Payload)
}

func (o *DeleteRuleGroupsMixin0OK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuleGroupsMixin0Forbidden creates a DeleteRuleGroupsMixin0Forbidden with default headers values
func NewDeleteRuleGroupsMixin0Forbidden() *DeleteRuleGroupsMixin0Forbidden {
	return &DeleteRuleGroupsMixin0Forbidden{}
}

/*
DeleteRuleGroupsMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRuleGroupsMixin0Forbidden struct {

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

// IsSuccess returns true when this delete rule groups mixin0 forbidden response has a 2xx status code
func (o *DeleteRuleGroupsMixin0Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rule groups mixin0 forbidden response has a 3xx status code
func (o *DeleteRuleGroupsMixin0Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rule groups mixin0 forbidden response has a 4xx status code
func (o *DeleteRuleGroupsMixin0Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rule groups mixin0 forbidden response has a 5xx status code
func (o *DeleteRuleGroupsMixin0Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rule groups mixin0 forbidden response a status code equal to that given
func (o *DeleteRuleGroupsMixin0Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete rule groups mixin0 forbidden response
func (o *DeleteRuleGroupsMixin0Forbidden) Code() int {
	return 403
}

func (o *DeleteRuleGroupsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteRuleGroupsMixin0Forbidden) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteRuleGroupsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuleGroupsMixin0NotFound creates a DeleteRuleGroupsMixin0NotFound with default headers values
func NewDeleteRuleGroupsMixin0NotFound() *DeleteRuleGroupsMixin0NotFound {
	return &DeleteRuleGroupsMixin0NotFound{}
}

/*
DeleteRuleGroupsMixin0NotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteRuleGroupsMixin0NotFound struct {

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

// IsSuccess returns true when this delete rule groups mixin0 not found response has a 2xx status code
func (o *DeleteRuleGroupsMixin0NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rule groups mixin0 not found response has a 3xx status code
func (o *DeleteRuleGroupsMixin0NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rule groups mixin0 not found response has a 4xx status code
func (o *DeleteRuleGroupsMixin0NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rule groups mixin0 not found response has a 5xx status code
func (o *DeleteRuleGroupsMixin0NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rule groups mixin0 not found response a status code equal to that given
func (o *DeleteRuleGroupsMixin0NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete rule groups mixin0 not found response
func (o *DeleteRuleGroupsMixin0NotFound) Code() int {
	return 404
}

func (o *DeleteRuleGroupsMixin0NotFound) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0NotFound  %+v", 404, o.Payload)
}

func (o *DeleteRuleGroupsMixin0NotFound) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0NotFound  %+v", 404, o.Payload)
}

func (o *DeleteRuleGroupsMixin0NotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuleGroupsMixin0TooManyRequests creates a DeleteRuleGroupsMixin0TooManyRequests with default headers values
func NewDeleteRuleGroupsMixin0TooManyRequests() *DeleteRuleGroupsMixin0TooManyRequests {
	return &DeleteRuleGroupsMixin0TooManyRequests{}
}

/*
DeleteRuleGroupsMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteRuleGroupsMixin0TooManyRequests struct {

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

// IsSuccess returns true when this delete rule groups mixin0 too many requests response has a 2xx status code
func (o *DeleteRuleGroupsMixin0TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rule groups mixin0 too many requests response has a 3xx status code
func (o *DeleteRuleGroupsMixin0TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rule groups mixin0 too many requests response has a 4xx status code
func (o *DeleteRuleGroupsMixin0TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rule groups mixin0 too many requests response has a 5xx status code
func (o *DeleteRuleGroupsMixin0TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rule groups mixin0 too many requests response a status code equal to that given
func (o *DeleteRuleGroupsMixin0TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete rule groups mixin0 too many requests response
func (o *DeleteRuleGroupsMixin0TooManyRequests) Code() int {
	return 429
}

func (o *DeleteRuleGroupsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRuleGroupsMixin0TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRuleGroupsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuleGroupsMixin0InternalServerError creates a DeleteRuleGroupsMixin0InternalServerError with default headers values
func NewDeleteRuleGroupsMixin0InternalServerError() *DeleteRuleGroupsMixin0InternalServerError {
	return &DeleteRuleGroupsMixin0InternalServerError{}
}

/*
DeleteRuleGroupsMixin0InternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type DeleteRuleGroupsMixin0InternalServerError struct {

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

// IsSuccess returns true when this delete rule groups mixin0 internal server error response has a 2xx status code
func (o *DeleteRuleGroupsMixin0InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rule groups mixin0 internal server error response has a 3xx status code
func (o *DeleteRuleGroupsMixin0InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rule groups mixin0 internal server error response has a 4xx status code
func (o *DeleteRuleGroupsMixin0InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rule groups mixin0 internal server error response has a 5xx status code
func (o *DeleteRuleGroupsMixin0InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete rule groups mixin0 internal server error response a status code equal to that given
func (o *DeleteRuleGroupsMixin0InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete rule groups mixin0 internal server error response
func (o *DeleteRuleGroupsMixin0InternalServerError) Code() int {
	return 500
}

func (o *DeleteRuleGroupsMixin0InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRuleGroupsMixin0InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRuleGroupsMixin0InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
