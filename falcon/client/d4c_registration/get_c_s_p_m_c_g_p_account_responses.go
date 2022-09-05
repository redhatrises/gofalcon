// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// GetCSPMCGPAccountReader is a Reader for the GetCSPMCGPAccount structure.
type GetCSPMCGPAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMCGPAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMCGPAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCSPMCGPAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMCGPAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMCGPAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMCGPAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMCGPAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMCGPAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMCGPAccountOK creates a GetCSPMCGPAccountOK with default headers values
func NewGetCSPMCGPAccountOK() *GetCSPMCGPAccountOK {
	return &GetCSPMCGPAccountOK{}
}

/*
GetCSPMCGPAccountOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMCGPAccountOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this get c s p m c g p account o k response has a 2xx status code
func (o *GetCSPMCGPAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s p m c g p account o k response has a 3xx status code
func (o *GetCSPMCGPAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m c g p account o k response has a 4xx status code
func (o *GetCSPMCGPAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m c g p account o k response has a 5xx status code
func (o *GetCSPMCGPAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m c g p account o k response a status code equal to that given
func (o *GetCSPMCGPAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCSPMCGPAccountOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountOK  %+v", 200, o.Payload)
}

func (o *GetCSPMCGPAccountOK) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountOK  %+v", 200, o.Payload)
}

func (o *GetCSPMCGPAccountOK) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMCGPAccountMultiStatus creates a GetCSPMCGPAccountMultiStatus with default headers values
func NewGetCSPMCGPAccountMultiStatus() *GetCSPMCGPAccountMultiStatus {
	return &GetCSPMCGPAccountMultiStatus{}
}

/*
GetCSPMCGPAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCSPMCGPAccountMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this get c s p m c g p account multi status response has a 2xx status code
func (o *GetCSPMCGPAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s p m c g p account multi status response has a 3xx status code
func (o *GetCSPMCGPAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m c g p account multi status response has a 4xx status code
func (o *GetCSPMCGPAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m c g p account multi status response has a 5xx status code
func (o *GetCSPMCGPAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m c g p account multi status response a status code equal to that given
func (o *GetCSPMCGPAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

func (o *GetCSPMCGPAccountMultiStatus) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCSPMCGPAccountMultiStatus) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCSPMCGPAccountMultiStatus) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMCGPAccountBadRequest creates a GetCSPMCGPAccountBadRequest with default headers values
func NewGetCSPMCGPAccountBadRequest() *GetCSPMCGPAccountBadRequest {
	return &GetCSPMCGPAccountBadRequest{}
}

/*
GetCSPMCGPAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMCGPAccountBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this get c s p m c g p account bad request response has a 2xx status code
func (o *GetCSPMCGPAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m c g p account bad request response has a 3xx status code
func (o *GetCSPMCGPAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m c g p account bad request response has a 4xx status code
func (o *GetCSPMCGPAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m c g p account bad request response has a 5xx status code
func (o *GetCSPMCGPAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m c g p account bad request response a status code equal to that given
func (o *GetCSPMCGPAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCSPMCGPAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMCGPAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMCGPAccountBadRequest) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMCGPAccountForbidden creates a GetCSPMCGPAccountForbidden with default headers values
func NewGetCSPMCGPAccountForbidden() *GetCSPMCGPAccountForbidden {
	return &GetCSPMCGPAccountForbidden{}
}

/*
GetCSPMCGPAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMCGPAccountForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get c s p m c g p account forbidden response has a 2xx status code
func (o *GetCSPMCGPAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m c g p account forbidden response has a 3xx status code
func (o *GetCSPMCGPAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m c g p account forbidden response has a 4xx status code
func (o *GetCSPMCGPAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m c g p account forbidden response has a 5xx status code
func (o *GetCSPMCGPAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m c g p account forbidden response a status code equal to that given
func (o *GetCSPMCGPAccountForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCSPMCGPAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMCGPAccountForbidden) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMCGPAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMCGPAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMCGPAccountTooManyRequests creates a GetCSPMCGPAccountTooManyRequests with default headers values
func NewGetCSPMCGPAccountTooManyRequests() *GetCSPMCGPAccountTooManyRequests {
	return &GetCSPMCGPAccountTooManyRequests{}
}

/*
GetCSPMCGPAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMCGPAccountTooManyRequests struct {

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

// IsSuccess returns true when this get c s p m c g p account too many requests response has a 2xx status code
func (o *GetCSPMCGPAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m c g p account too many requests response has a 3xx status code
func (o *GetCSPMCGPAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m c g p account too many requests response has a 4xx status code
func (o *GetCSPMCGPAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m c g p account too many requests response has a 5xx status code
func (o *GetCSPMCGPAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m c g p account too many requests response a status code equal to that given
func (o *GetCSPMCGPAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCSPMCGPAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMCGPAccountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMCGPAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMCGPAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMCGPAccountInternalServerError creates a GetCSPMCGPAccountInternalServerError with default headers values
func NewGetCSPMCGPAccountInternalServerError() *GetCSPMCGPAccountInternalServerError {
	return &GetCSPMCGPAccountInternalServerError{}
}

/*
GetCSPMCGPAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMCGPAccountInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this get c s p m c g p account internal server error response has a 2xx status code
func (o *GetCSPMCGPAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m c g p account internal server error response has a 3xx status code
func (o *GetCSPMCGPAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m c g p account internal server error response has a 4xx status code
func (o *GetCSPMCGPAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m c g p account internal server error response has a 5xx status code
func (o *GetCSPMCGPAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get c s p m c g p account internal server error response a status code equal to that given
func (o *GetCSPMCGPAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCSPMCGPAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMCGPAccountInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMCGPAccountInternalServerError) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMCGPAccountDefault creates a GetCSPMCGPAccountDefault with default headers values
func NewGetCSPMCGPAccountDefault(code int) *GetCSPMCGPAccountDefault {
	return &GetCSPMCGPAccountDefault{
		_statusCode: code,
	}
}

/*
GetCSPMCGPAccountDefault describes a response with status code -1, with default header values.

OK
*/
type GetCSPMCGPAccountDefault struct {
	_statusCode int

	Payload *models.RegistrationGCPAccountResponseV1
}

// Code gets the status code for the get c s p m c g p account default response
func (o *GetCSPMCGPAccountDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get c s p m c g p account default response has a 2xx status code
func (o *GetCSPMCGPAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get c s p m c g p account default response has a 3xx status code
func (o *GetCSPMCGPAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get c s p m c g p account default response has a 4xx status code
func (o *GetCSPMCGPAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get c s p m c g p account default response has a 5xx status code
func (o *GetCSPMCGPAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get c s p m c g p account default response a status code equal to that given
func (o *GetCSPMCGPAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCSPMCGPAccountDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] GetCSPMCGPAccount default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMCGPAccountDefault) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] GetCSPMCGPAccount default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMCGPAccountDefault) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
