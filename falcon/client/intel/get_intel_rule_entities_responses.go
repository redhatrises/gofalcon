// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// GetIntelRuleEntitiesReader is a Reader for the GetIntelRuleEntities structure.
type GetIntelRuleEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntelRuleEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntelRuleEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetIntelRuleEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntelRuleEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntelRuleEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIntelRuleEntitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIntelRuleEntitiesOK creates a GetIntelRuleEntitiesOK with default headers values
func NewGetIntelRuleEntitiesOK() *GetIntelRuleEntitiesOK {
	return &GetIntelRuleEntitiesOK{}
}

/*
GetIntelRuleEntitiesOK describes a response with status code 200, with default header values.

OK
*/
type GetIntelRuleEntitiesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainRulesResponse
}

// IsSuccess returns true when this get intel rule entities o k response has a 2xx status code
func (o *GetIntelRuleEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get intel rule entities o k response has a 3xx status code
func (o *GetIntelRuleEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule entities o k response has a 4xx status code
func (o *GetIntelRuleEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get intel rule entities o k response has a 5xx status code
func (o *GetIntelRuleEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel rule entities o k response a status code equal to that given
func (o *GetIntelRuleEntitiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntelRuleEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] getIntelRuleEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetIntelRuleEntitiesOK) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] getIntelRuleEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetIntelRuleEntitiesOK) GetPayload() *models.DomainRulesResponse {
	return o.Payload
}

func (o *GetIntelRuleEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntelRuleEntitiesForbidden creates a GetIntelRuleEntitiesForbidden with default headers values
func NewGetIntelRuleEntitiesForbidden() *GetIntelRuleEntitiesForbidden {
	return &GetIntelRuleEntitiesForbidden{}
}

/*
GetIntelRuleEntitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetIntelRuleEntitiesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get intel rule entities forbidden response has a 2xx status code
func (o *GetIntelRuleEntitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel rule entities forbidden response has a 3xx status code
func (o *GetIntelRuleEntitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule entities forbidden response has a 4xx status code
func (o *GetIntelRuleEntitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel rule entities forbidden response has a 5xx status code
func (o *GetIntelRuleEntitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel rule entities forbidden response a status code equal to that given
func (o *GetIntelRuleEntitiesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntelRuleEntitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] getIntelRuleEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntelRuleEntitiesForbidden) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] getIntelRuleEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntelRuleEntitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntelRuleEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelRuleEntitiesTooManyRequests creates a GetIntelRuleEntitiesTooManyRequests with default headers values
func NewGetIntelRuleEntitiesTooManyRequests() *GetIntelRuleEntitiesTooManyRequests {
	return &GetIntelRuleEntitiesTooManyRequests{}
}

/*
GetIntelRuleEntitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetIntelRuleEntitiesTooManyRequests struct {

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

// IsSuccess returns true when this get intel rule entities too many requests response has a 2xx status code
func (o *GetIntelRuleEntitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel rule entities too many requests response has a 3xx status code
func (o *GetIntelRuleEntitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule entities too many requests response has a 4xx status code
func (o *GetIntelRuleEntitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel rule entities too many requests response has a 5xx status code
func (o *GetIntelRuleEntitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel rule entities too many requests response a status code equal to that given
func (o *GetIntelRuleEntitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntelRuleEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] getIntelRuleEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntelRuleEntitiesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] getIntelRuleEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntelRuleEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntelRuleEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelRuleEntitiesInternalServerError creates a GetIntelRuleEntitiesInternalServerError with default headers values
func NewGetIntelRuleEntitiesInternalServerError() *GetIntelRuleEntitiesInternalServerError {
	return &GetIntelRuleEntitiesInternalServerError{}
}

/*
GetIntelRuleEntitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetIntelRuleEntitiesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get intel rule entities internal server error response has a 2xx status code
func (o *GetIntelRuleEntitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel rule entities internal server error response has a 3xx status code
func (o *GetIntelRuleEntitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule entities internal server error response has a 4xx status code
func (o *GetIntelRuleEntitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get intel rule entities internal server error response has a 5xx status code
func (o *GetIntelRuleEntitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get intel rule entities internal server error response a status code equal to that given
func (o *GetIntelRuleEntitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntelRuleEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] getIntelRuleEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntelRuleEntitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] getIntelRuleEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntelRuleEntitiesInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetIntelRuleEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelRuleEntitiesDefault creates a GetIntelRuleEntitiesDefault with default headers values
func NewGetIntelRuleEntitiesDefault(code int) *GetIntelRuleEntitiesDefault {
	return &GetIntelRuleEntitiesDefault{
		_statusCode: code,
	}
}

/*
GetIntelRuleEntitiesDefault describes a response with status code -1, with default header values.

OK
*/
type GetIntelRuleEntitiesDefault struct {
	_statusCode int

	Payload *models.DomainRulesResponse
}

// Code gets the status code for the get intel rule entities default response
func (o *GetIntelRuleEntitiesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get intel rule entities default response has a 2xx status code
func (o *GetIntelRuleEntitiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get intel rule entities default response has a 3xx status code
func (o *GetIntelRuleEntitiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get intel rule entities default response has a 4xx status code
func (o *GetIntelRuleEntitiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get intel rule entities default response has a 5xx status code
func (o *GetIntelRuleEntitiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get intel rule entities default response a status code equal to that given
func (o *GetIntelRuleEntitiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetIntelRuleEntitiesDefault) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] GetIntelRuleEntities default  %+v", o._statusCode, o.Payload)
}

func (o *GetIntelRuleEntitiesDefault) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules/v1][%d] GetIntelRuleEntities default  %+v", o._statusCode, o.Payload)
}

func (o *GetIntelRuleEntitiesDefault) GetPayload() *models.DomainRulesResponse {
	return o.Payload
}

func (o *GetIntelRuleEntitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
