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

// GetCSPMAwsConsoleSetupURLsReader is a Reader for the GetCSPMAwsConsoleSetupURLs structure.
type GetCSPMAwsConsoleSetupURLsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMAwsConsoleSetupURLsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMAwsConsoleSetupURLsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCSPMAwsConsoleSetupURLsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMAwsConsoleSetupURLsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMAwsConsoleSetupURLsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMAwsConsoleSetupURLsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMAwsConsoleSetupURLsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMAwsConsoleSetupURLsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMAwsConsoleSetupURLsOK creates a GetCSPMAwsConsoleSetupURLsOK with default headers values
func NewGetCSPMAwsConsoleSetupURLsOK() *GetCSPMAwsConsoleSetupURLsOK {
	return &GetCSPMAwsConsoleSetupURLsOK{}
}

/* GetCSPMAwsConsoleSetupURLsOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMAwsConsoleSetupURLsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountConsoleURL
}

func (o *GetCSPMAwsConsoleSetupURLsOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/console-setup-urls/v1][%d] getCSPMAwsConsoleSetupURLsOK  %+v", 200, o.Payload)
}
func (o *GetCSPMAwsConsoleSetupURLsOK) GetPayload() *models.RegistrationAWSAccountConsoleURL {
	return o.Payload
}

func (o *GetCSPMAwsConsoleSetupURLsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountConsoleURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAwsConsoleSetupURLsMultiStatus creates a GetCSPMAwsConsoleSetupURLsMultiStatus with default headers values
func NewGetCSPMAwsConsoleSetupURLsMultiStatus() *GetCSPMAwsConsoleSetupURLsMultiStatus {
	return &GetCSPMAwsConsoleSetupURLsMultiStatus{}
}

/* GetCSPMAwsConsoleSetupURLsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCSPMAwsConsoleSetupURLsMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountConsoleURL
}

func (o *GetCSPMAwsConsoleSetupURLsMultiStatus) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/console-setup-urls/v1][%d] getCSPMAwsConsoleSetupURLsMultiStatus  %+v", 207, o.Payload)
}
func (o *GetCSPMAwsConsoleSetupURLsMultiStatus) GetPayload() *models.RegistrationAWSAccountConsoleURL {
	return o.Payload
}

func (o *GetCSPMAwsConsoleSetupURLsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountConsoleURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAwsConsoleSetupURLsBadRequest creates a GetCSPMAwsConsoleSetupURLsBadRequest with default headers values
func NewGetCSPMAwsConsoleSetupURLsBadRequest() *GetCSPMAwsConsoleSetupURLsBadRequest {
	return &GetCSPMAwsConsoleSetupURLsBadRequest{}
}

/* GetCSPMAwsConsoleSetupURLsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMAwsConsoleSetupURLsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountConsoleURL
}

func (o *GetCSPMAwsConsoleSetupURLsBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/console-setup-urls/v1][%d] getCSPMAwsConsoleSetupURLsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCSPMAwsConsoleSetupURLsBadRequest) GetPayload() *models.RegistrationAWSAccountConsoleURL {
	return o.Payload
}

func (o *GetCSPMAwsConsoleSetupURLsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountConsoleURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAwsConsoleSetupURLsForbidden creates a GetCSPMAwsConsoleSetupURLsForbidden with default headers values
func NewGetCSPMAwsConsoleSetupURLsForbidden() *GetCSPMAwsConsoleSetupURLsForbidden {
	return &GetCSPMAwsConsoleSetupURLsForbidden{}
}

/* GetCSPMAwsConsoleSetupURLsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMAwsConsoleSetupURLsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetCSPMAwsConsoleSetupURLsForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/console-setup-urls/v1][%d] getCSPMAwsConsoleSetupURLsForbidden  %+v", 403, o.Payload)
}
func (o *GetCSPMAwsConsoleSetupURLsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMAwsConsoleSetupURLsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAwsConsoleSetupURLsTooManyRequests creates a GetCSPMAwsConsoleSetupURLsTooManyRequests with default headers values
func NewGetCSPMAwsConsoleSetupURLsTooManyRequests() *GetCSPMAwsConsoleSetupURLsTooManyRequests {
	return &GetCSPMAwsConsoleSetupURLsTooManyRequests{}
}

/* GetCSPMAwsConsoleSetupURLsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMAwsConsoleSetupURLsTooManyRequests struct {

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

func (o *GetCSPMAwsConsoleSetupURLsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/console-setup-urls/v1][%d] getCSPMAwsConsoleSetupURLsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCSPMAwsConsoleSetupURLsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMAwsConsoleSetupURLsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAwsConsoleSetupURLsInternalServerError creates a GetCSPMAwsConsoleSetupURLsInternalServerError with default headers values
func NewGetCSPMAwsConsoleSetupURLsInternalServerError() *GetCSPMAwsConsoleSetupURLsInternalServerError {
	return &GetCSPMAwsConsoleSetupURLsInternalServerError{}
}

/* GetCSPMAwsConsoleSetupURLsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMAwsConsoleSetupURLsInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountConsoleURL
}

func (o *GetCSPMAwsConsoleSetupURLsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/console-setup-urls/v1][%d] getCSPMAwsConsoleSetupURLsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCSPMAwsConsoleSetupURLsInternalServerError) GetPayload() *models.RegistrationAWSAccountConsoleURL {
	return o.Payload
}

func (o *GetCSPMAwsConsoleSetupURLsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountConsoleURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAwsConsoleSetupURLsDefault creates a GetCSPMAwsConsoleSetupURLsDefault with default headers values
func NewGetCSPMAwsConsoleSetupURLsDefault(code int) *GetCSPMAwsConsoleSetupURLsDefault {
	return &GetCSPMAwsConsoleSetupURLsDefault{
		_statusCode: code,
	}
}

/* GetCSPMAwsConsoleSetupURLsDefault describes a response with status code -1, with default header values.

OK
*/
type GetCSPMAwsConsoleSetupURLsDefault struct {
	_statusCode int

	Payload *models.RegistrationAWSAccountConsoleURL
}

// Code gets the status code for the get c s p m aws console setup u r ls default response
func (o *GetCSPMAwsConsoleSetupURLsDefault) Code() int {
	return o._statusCode
}

func (o *GetCSPMAwsConsoleSetupURLsDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/console-setup-urls/v1][%d] GetCSPMAwsConsoleSetupURLs default  %+v", o._statusCode, o.Payload)
}
func (o *GetCSPMAwsConsoleSetupURLsDefault) GetPayload() *models.RegistrationAWSAccountConsoleURL {
	return o.Payload
}

func (o *GetCSPMAwsConsoleSetupURLsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationAWSAccountConsoleURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}