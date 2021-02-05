// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

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

// QueryCombinedFirewallPolicyMembersReader is a Reader for the QueryCombinedFirewallPolicyMembers structure.
type QueryCombinedFirewallPolicyMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedFirewallPolicyMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedFirewallPolicyMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedFirewallPolicyMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedFirewallPolicyMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryCombinedFirewallPolicyMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedFirewallPolicyMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedFirewallPolicyMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedFirewallPolicyMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedFirewallPolicyMembersOK creates a QueryCombinedFirewallPolicyMembersOK with default headers values
func NewQueryCombinedFirewallPolicyMembersOK() *QueryCombinedFirewallPolicyMembersOK {
	return &QueryCombinedFirewallPolicyMembersOK{}
}

/* QueryCombinedFirewallPolicyMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedFirewallPolicyMembersOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedFirewallPolicyMembersOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall-members/v1][%d] queryCombinedFirewallPolicyMembersOK  %+v", 200, o.Payload)
}
func (o *QueryCombinedFirewallPolicyMembersOK) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPolicyMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPolicyMembersBadRequest creates a QueryCombinedFirewallPolicyMembersBadRequest with default headers values
func NewQueryCombinedFirewallPolicyMembersBadRequest() *QueryCombinedFirewallPolicyMembersBadRequest {
	return &QueryCombinedFirewallPolicyMembersBadRequest{}
}

/* QueryCombinedFirewallPolicyMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedFirewallPolicyMembersBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedFirewallPolicyMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall-members/v1][%d] queryCombinedFirewallPolicyMembersBadRequest  %+v", 400, o.Payload)
}
func (o *QueryCombinedFirewallPolicyMembersBadRequest) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPolicyMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPolicyMembersForbidden creates a QueryCombinedFirewallPolicyMembersForbidden with default headers values
func NewQueryCombinedFirewallPolicyMembersForbidden() *QueryCombinedFirewallPolicyMembersForbidden {
	return &QueryCombinedFirewallPolicyMembersForbidden{}
}

/* QueryCombinedFirewallPolicyMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedFirewallPolicyMembersForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryCombinedFirewallPolicyMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall-members/v1][%d] queryCombinedFirewallPolicyMembersForbidden  %+v", 403, o.Payload)
}
func (o *QueryCombinedFirewallPolicyMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedFirewallPolicyMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedFirewallPolicyMembersNotFound creates a QueryCombinedFirewallPolicyMembersNotFound with default headers values
func NewQueryCombinedFirewallPolicyMembersNotFound() *QueryCombinedFirewallPolicyMembersNotFound {
	return &QueryCombinedFirewallPolicyMembersNotFound{}
}

/* QueryCombinedFirewallPolicyMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryCombinedFirewallPolicyMembersNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedFirewallPolicyMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall-members/v1][%d] queryCombinedFirewallPolicyMembersNotFound  %+v", 404, o.Payload)
}
func (o *QueryCombinedFirewallPolicyMembersNotFound) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPolicyMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPolicyMembersTooManyRequests creates a QueryCombinedFirewallPolicyMembersTooManyRequests with default headers values
func NewQueryCombinedFirewallPolicyMembersTooManyRequests() *QueryCombinedFirewallPolicyMembersTooManyRequests {
	return &QueryCombinedFirewallPolicyMembersTooManyRequests{}
}

/* QueryCombinedFirewallPolicyMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedFirewallPolicyMembersTooManyRequests struct {

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

func (o *QueryCombinedFirewallPolicyMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall-members/v1][%d] queryCombinedFirewallPolicyMembersTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryCombinedFirewallPolicyMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedFirewallPolicyMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedFirewallPolicyMembersInternalServerError creates a QueryCombinedFirewallPolicyMembersInternalServerError with default headers values
func NewQueryCombinedFirewallPolicyMembersInternalServerError() *QueryCombinedFirewallPolicyMembersInternalServerError {
	return &QueryCombinedFirewallPolicyMembersInternalServerError{}
}

/* QueryCombinedFirewallPolicyMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedFirewallPolicyMembersInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

func (o *QueryCombinedFirewallPolicyMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall-members/v1][%d] queryCombinedFirewallPolicyMembersInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryCombinedFirewallPolicyMembersInternalServerError) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPolicyMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPolicyMembersDefault creates a QueryCombinedFirewallPolicyMembersDefault with default headers values
func NewQueryCombinedFirewallPolicyMembersDefault(code int) *QueryCombinedFirewallPolicyMembersDefault {
	return &QueryCombinedFirewallPolicyMembersDefault{
		_statusCode: code,
	}
}

/* QueryCombinedFirewallPolicyMembersDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedFirewallPolicyMembersDefault struct {
	_statusCode int

	Payload *models.ResponsesPolicyMembersRespV1
}

// Code gets the status code for the query combined firewall policy members default response
func (o *QueryCombinedFirewallPolicyMembersDefault) Code() int {
	return o._statusCode
}

func (o *QueryCombinedFirewallPolicyMembersDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall-members/v1][%d] queryCombinedFirewallPolicyMembers default  %+v", o._statusCode, o.Payload)
}
func (o *QueryCombinedFirewallPolicyMembersDefault) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPolicyMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}