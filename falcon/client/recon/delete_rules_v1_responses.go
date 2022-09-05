// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// DeleteRulesV1Reader is a Reader for the DeleteRulesV1 structure.
type DeleteRulesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRulesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRulesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRulesV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRulesV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRulesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRulesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRulesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRulesV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRulesV1OK creates a DeleteRulesV1OK with default headers values
func NewDeleteRulesV1OK() *DeleteRulesV1OK {
	return &DeleteRulesV1OK{}
}

/*
DeleteRulesV1OK describes a response with status code 200, with default header values.

OK
*/
type DeleteRulesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainRuleQueryResponseV1
}

// IsSuccess returns true when this delete rules v1 o k response has a 2xx status code
func (o *DeleteRulesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete rules v1 o k response has a 3xx status code
func (o *DeleteRulesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules v1 o k response has a 4xx status code
func (o *DeleteRulesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rules v1 o k response has a 5xx status code
func (o *DeleteRulesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules v1 o k response a status code equal to that given
func (o *DeleteRulesV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRulesV1OK) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1OK  %+v", 200, o.Payload)
}

func (o *DeleteRulesV1OK) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1OK  %+v", 200, o.Payload)
}

func (o *DeleteRulesV1OK) GetPayload() *models.DomainRuleQueryResponseV1 {
	return o.Payload
}

func (o *DeleteRulesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainRuleQueryResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRulesV1BadRequest creates a DeleteRulesV1BadRequest with default headers values
func NewDeleteRulesV1BadRequest() *DeleteRulesV1BadRequest {
	return &DeleteRulesV1BadRequest{}
}

/*
DeleteRulesV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteRulesV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this delete rules v1 bad request response has a 2xx status code
func (o *DeleteRulesV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rules v1 bad request response has a 3xx status code
func (o *DeleteRulesV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules v1 bad request response has a 4xx status code
func (o *DeleteRulesV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rules v1 bad request response has a 5xx status code
func (o *DeleteRulesV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules v1 bad request response a status code equal to that given
func (o *DeleteRulesV1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRulesV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRulesV1BadRequest) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRulesV1BadRequest) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *DeleteRulesV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRulesV1Unauthorized creates a DeleteRulesV1Unauthorized with default headers values
func NewDeleteRulesV1Unauthorized() *DeleteRulesV1Unauthorized {
	return &DeleteRulesV1Unauthorized{}
}

/*
DeleteRulesV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRulesV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this delete rules v1 unauthorized response has a 2xx status code
func (o *DeleteRulesV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rules v1 unauthorized response has a 3xx status code
func (o *DeleteRulesV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules v1 unauthorized response has a 4xx status code
func (o *DeleteRulesV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rules v1 unauthorized response has a 5xx status code
func (o *DeleteRulesV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules v1 unauthorized response a status code equal to that given
func (o *DeleteRulesV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRulesV1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRulesV1Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRulesV1Unauthorized) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *DeleteRulesV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRulesV1Forbidden creates a DeleteRulesV1Forbidden with default headers values
func NewDeleteRulesV1Forbidden() *DeleteRulesV1Forbidden {
	return &DeleteRulesV1Forbidden{}
}

/*
DeleteRulesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRulesV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this delete rules v1 forbidden response has a 2xx status code
func (o *DeleteRulesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rules v1 forbidden response has a 3xx status code
func (o *DeleteRulesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules v1 forbidden response has a 4xx status code
func (o *DeleteRulesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rules v1 forbidden response has a 5xx status code
func (o *DeleteRulesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules v1 forbidden response a status code equal to that given
func (o *DeleteRulesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRulesV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteRulesV1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteRulesV1Forbidden) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *DeleteRulesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRulesV1TooManyRequests creates a DeleteRulesV1TooManyRequests with default headers values
func NewDeleteRulesV1TooManyRequests() *DeleteRulesV1TooManyRequests {
	return &DeleteRulesV1TooManyRequests{}
}

/*
DeleteRulesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteRulesV1TooManyRequests struct {

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

// IsSuccess returns true when this delete rules v1 too many requests response has a 2xx status code
func (o *DeleteRulesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rules v1 too many requests response has a 3xx status code
func (o *DeleteRulesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules v1 too many requests response has a 4xx status code
func (o *DeleteRulesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete rules v1 too many requests response has a 5xx status code
func (o *DeleteRulesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete rules v1 too many requests response a status code equal to that given
func (o *DeleteRulesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRulesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRulesV1TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRulesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRulesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRulesV1InternalServerError creates a DeleteRulesV1InternalServerError with default headers values
func NewDeleteRulesV1InternalServerError() *DeleteRulesV1InternalServerError {
	return &DeleteRulesV1InternalServerError{}
}

/*
DeleteRulesV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteRulesV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this delete rules v1 internal server error response has a 2xx status code
func (o *DeleteRulesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete rules v1 internal server error response has a 3xx status code
func (o *DeleteRulesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete rules v1 internal server error response has a 4xx status code
func (o *DeleteRulesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete rules v1 internal server error response has a 5xx status code
func (o *DeleteRulesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete rules v1 internal server error response a status code equal to that given
func (o *DeleteRulesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRulesV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRulesV1InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] deleteRulesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRulesV1InternalServerError) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *DeleteRulesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRulesV1Default creates a DeleteRulesV1Default with default headers values
func NewDeleteRulesV1Default(code int) *DeleteRulesV1Default {
	return &DeleteRulesV1Default{
		_statusCode: code,
	}
}

/*
DeleteRulesV1Default describes a response with status code -1, with default header values.

OK
*/
type DeleteRulesV1Default struct {
	_statusCode int

	Payload *models.DomainRuleQueryResponseV1
}

// Code gets the status code for the delete rules v1 default response
func (o *DeleteRulesV1Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete rules v1 default response has a 2xx status code
func (o *DeleteRulesV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete rules v1 default response has a 3xx status code
func (o *DeleteRulesV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete rules v1 default response has a 4xx status code
func (o *DeleteRulesV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete rules v1 default response has a 5xx status code
func (o *DeleteRulesV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete rules v1 default response a status code equal to that given
func (o *DeleteRulesV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteRulesV1Default) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] DeleteRulesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRulesV1Default) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/rules/v1][%d] DeleteRulesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRulesV1Default) GetPayload() *models.DomainRuleQueryResponseV1 {
	return o.Payload
}

func (o *DeleteRulesV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainRuleQueryResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
