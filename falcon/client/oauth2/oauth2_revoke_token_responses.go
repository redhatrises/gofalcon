// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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

// Oauth2RevokeTokenReader is a Reader for the Oauth2RevokeToken structure.
type Oauth2RevokeTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Oauth2RevokeTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOauth2RevokeTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOauth2RevokeTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOauth2RevokeTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOauth2RevokeTokenTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOauth2RevokeTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOauth2RevokeTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOauth2RevokeTokenOK creates a Oauth2RevokeTokenOK with default headers values
func NewOauth2RevokeTokenOK() *Oauth2RevokeTokenOK {
	return &Oauth2RevokeTokenOK{}
}

/* Oauth2RevokeTokenOK describes a response with status code 200, with default header values.

Successfully revoked token
*/
type Oauth2RevokeTokenOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *Oauth2RevokeTokenOK) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenOK  %+v", 200, o.Payload)
}
func (o *Oauth2RevokeTokenOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenBadRequest creates a Oauth2RevokeTokenBadRequest with default headers values
func NewOauth2RevokeTokenBadRequest() *Oauth2RevokeTokenBadRequest {
	return &Oauth2RevokeTokenBadRequest{}
}

/* Oauth2RevokeTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type Oauth2RevokeTokenBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *Oauth2RevokeTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenBadRequest  %+v", 400, o.Payload)
}
func (o *Oauth2RevokeTokenBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenForbidden creates a Oauth2RevokeTokenForbidden with default headers values
func NewOauth2RevokeTokenForbidden() *Oauth2RevokeTokenForbidden {
	return &Oauth2RevokeTokenForbidden{}
}

/* Oauth2RevokeTokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type Oauth2RevokeTokenForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *Oauth2RevokeTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenForbidden  %+v", 403, o.Payload)
}
func (o *Oauth2RevokeTokenForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenTooManyRequests creates a Oauth2RevokeTokenTooManyRequests with default headers values
func NewOauth2RevokeTokenTooManyRequests() *Oauth2RevokeTokenTooManyRequests {
	return &Oauth2RevokeTokenTooManyRequests{}
}

/* Oauth2RevokeTokenTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type Oauth2RevokeTokenTooManyRequests struct {

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

func (o *Oauth2RevokeTokenTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenTooManyRequests  %+v", 429, o.Payload)
}
func (o *Oauth2RevokeTokenTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenInternalServerError creates a Oauth2RevokeTokenInternalServerError with default headers values
func NewOauth2RevokeTokenInternalServerError() *Oauth2RevokeTokenInternalServerError {
	return &Oauth2RevokeTokenInternalServerError{}
}

/* Oauth2RevokeTokenInternalServerError describes a response with status code 500, with default header values.

Failed to revoke token
*/
type Oauth2RevokeTokenInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *Oauth2RevokeTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *Oauth2RevokeTokenInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenDefault creates a Oauth2RevokeTokenDefault with default headers values
func NewOauth2RevokeTokenDefault(code int) *Oauth2RevokeTokenDefault {
	return &Oauth2RevokeTokenDefault{
		_statusCode: code,
	}
}

/* Oauth2RevokeTokenDefault describes a response with status code -1, with default header values.

Successfully revoked token
*/
type Oauth2RevokeTokenDefault struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// Code gets the status code for the oauth2 revoke token default response
func (o *Oauth2RevokeTokenDefault) Code() int {
	return o._statusCode
}

func (o *Oauth2RevokeTokenDefault) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeToken default  %+v", o._statusCode, o.Payload)
}
func (o *Oauth2RevokeTokenDefault) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}