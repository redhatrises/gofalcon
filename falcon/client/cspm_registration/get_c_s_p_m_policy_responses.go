// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// GetCSPMPolicyReader is a Reader for the GetCSPMPolicy structure.
type GetCSPMPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCSPMPolicyMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMPolicyOK creates a GetCSPMPolicyOK with default headers values
func NewGetCSPMPolicyOK() *GetCSPMPolicyOK {
	return &GetCSPMPolicyOK{}
}

/*
GetCSPMPolicyOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMPolicyOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationPolicyResponseV1
}

// IsSuccess returns true when this get c s p m policy o k response has a 2xx status code
func (o *GetCSPMPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s p m policy o k response has a 3xx status code
func (o *GetCSPMPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy o k response has a 4xx status code
func (o *GetCSPMPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m policy o k response has a 5xx status code
func (o *GetCSPMPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy o k response a status code equal to that given
func (o *GetCSPMPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCSPMPolicyOK) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyOK  %+v", 200, o.Payload)
}

func (o *GetCSPMPolicyOK) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyOK  %+v", 200, o.Payload)
}

func (o *GetCSPMPolicyOK) GetPayload() *models.RegistrationPolicyResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationPolicyResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMPolicyMultiStatus creates a GetCSPMPolicyMultiStatus with default headers values
func NewGetCSPMPolicyMultiStatus() *GetCSPMPolicyMultiStatus {
	return &GetCSPMPolicyMultiStatus{}
}

/*
GetCSPMPolicyMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCSPMPolicyMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationPolicyResponseV1
}

// IsSuccess returns true when this get c s p m policy multi status response has a 2xx status code
func (o *GetCSPMPolicyMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s p m policy multi status response has a 3xx status code
func (o *GetCSPMPolicyMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy multi status response has a 4xx status code
func (o *GetCSPMPolicyMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m policy multi status response has a 5xx status code
func (o *GetCSPMPolicyMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy multi status response a status code equal to that given
func (o *GetCSPMPolicyMultiStatus) IsCode(code int) bool {
	return code == 207
}

func (o *GetCSPMPolicyMultiStatus) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCSPMPolicyMultiStatus) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCSPMPolicyMultiStatus) GetPayload() *models.RegistrationPolicyResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicyMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationPolicyResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMPolicyBadRequest creates a GetCSPMPolicyBadRequest with default headers values
func NewGetCSPMPolicyBadRequest() *GetCSPMPolicyBadRequest {
	return &GetCSPMPolicyBadRequest{}
}

/*
GetCSPMPolicyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMPolicyBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationPolicyResponseV1
}

// IsSuccess returns true when this get c s p m policy bad request response has a 2xx status code
func (o *GetCSPMPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m policy bad request response has a 3xx status code
func (o *GetCSPMPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy bad request response has a 4xx status code
func (o *GetCSPMPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m policy bad request response has a 5xx status code
func (o *GetCSPMPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy bad request response a status code equal to that given
func (o *GetCSPMPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCSPMPolicyBadRequest) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMPolicyBadRequest) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMPolicyBadRequest) GetPayload() *models.RegistrationPolicyResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationPolicyResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMPolicyForbidden creates a GetCSPMPolicyForbidden with default headers values
func NewGetCSPMPolicyForbidden() *GetCSPMPolicyForbidden {
	return &GetCSPMPolicyForbidden{}
}

/*
GetCSPMPolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMPolicyForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get c s p m policy forbidden response has a 2xx status code
func (o *GetCSPMPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m policy forbidden response has a 3xx status code
func (o *GetCSPMPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy forbidden response has a 4xx status code
func (o *GetCSPMPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m policy forbidden response has a 5xx status code
func (o *GetCSPMPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy forbidden response a status code equal to that given
func (o *GetCSPMPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCSPMPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMPolicyForbidden) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMPolicyForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMPolicyTooManyRequests creates a GetCSPMPolicyTooManyRequests with default headers values
func NewGetCSPMPolicyTooManyRequests() *GetCSPMPolicyTooManyRequests {
	return &GetCSPMPolicyTooManyRequests{}
}

/*
GetCSPMPolicyTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMPolicyTooManyRequests struct {

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

// IsSuccess returns true when this get c s p m policy too many requests response has a 2xx status code
func (o *GetCSPMPolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m policy too many requests response has a 3xx status code
func (o *GetCSPMPolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy too many requests response has a 4xx status code
func (o *GetCSPMPolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m policy too many requests response has a 5xx status code
func (o *GetCSPMPolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m policy too many requests response a status code equal to that given
func (o *GetCSPMPolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCSPMPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMPolicyTooManyRequests) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMPolicyTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMPolicyInternalServerError creates a GetCSPMPolicyInternalServerError with default headers values
func NewGetCSPMPolicyInternalServerError() *GetCSPMPolicyInternalServerError {
	return &GetCSPMPolicyInternalServerError{}
}

/*
GetCSPMPolicyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMPolicyInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationPolicyResponseV1
}

// IsSuccess returns true when this get c s p m policy internal server error response has a 2xx status code
func (o *GetCSPMPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m policy internal server error response has a 3xx status code
func (o *GetCSPMPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m policy internal server error response has a 4xx status code
func (o *GetCSPMPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m policy internal server error response has a 5xx status code
func (o *GetCSPMPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get c s p m policy internal server error response a status code equal to that given
func (o *GetCSPMPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCSPMPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMPolicyInternalServerError) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] getCSPMPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMPolicyInternalServerError) GetPayload() *models.RegistrationPolicyResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationPolicyResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMPolicyDefault creates a GetCSPMPolicyDefault with default headers values
func NewGetCSPMPolicyDefault(code int) *GetCSPMPolicyDefault {
	return &GetCSPMPolicyDefault{
		_statusCode: code,
	}
}

/*
GetCSPMPolicyDefault describes a response with status code -1, with default header values.

OK
*/
type GetCSPMPolicyDefault struct {
	_statusCode int

	Payload *models.RegistrationPolicyResponseV1
}

// Code gets the status code for the get c s p m policy default response
func (o *GetCSPMPolicyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get c s p m policy default response has a 2xx status code
func (o *GetCSPMPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get c s p m policy default response has a 3xx status code
func (o *GetCSPMPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get c s p m policy default response has a 4xx status code
func (o *GetCSPMPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get c s p m policy default response has a 5xx status code
func (o *GetCSPMPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get c s p m policy default response a status code equal to that given
func (o *GetCSPMPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCSPMPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] GetCSPMPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMPolicyDefault) String() string {
	return fmt.Sprintf("[GET /settings/entities/policy-details/v1][%d] GetCSPMPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMPolicyDefault) GetPayload() *models.RegistrationPolicyResponseV1 {
	return o.Payload
}

func (o *GetCSPMPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationPolicyResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
