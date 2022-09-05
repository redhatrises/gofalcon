// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// CreateSensorUpdatePoliciesReader is a Reader for the CreateSensorUpdatePolicies structure.
type CreateSensorUpdatePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSensorUpdatePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSensorUpdatePoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSensorUpdatePoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateSensorUpdatePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSensorUpdatePoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateSensorUpdatePoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateSensorUpdatePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSensorUpdatePoliciesCreated creates a CreateSensorUpdatePoliciesCreated with default headers values
func NewCreateSensorUpdatePoliciesCreated() *CreateSensorUpdatePoliciesCreated {
	return &CreateSensorUpdatePoliciesCreated{}
}

/*
CreateSensorUpdatePoliciesCreated describes a response with status code 201, with default header values.

Created
*/
type CreateSensorUpdatePoliciesCreated struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

// IsSuccess returns true when this create sensor update policies created response has a 2xx status code
func (o *CreateSensorUpdatePoliciesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create sensor update policies created response has a 3xx status code
func (o *CreateSensorUpdatePoliciesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies created response has a 4xx status code
func (o *CreateSensorUpdatePoliciesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create sensor update policies created response has a 5xx status code
func (o *CreateSensorUpdatePoliciesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies created response a status code equal to that given
func (o *CreateSensorUpdatePoliciesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateSensorUpdatePoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesCreated  %+v", 201, o.Payload)
}

func (o *CreateSensorUpdatePoliciesCreated) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesCreated  %+v", 201, o.Payload)
}

func (o *CreateSensorUpdatePoliciesCreated) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSensorUpdatePoliciesBadRequest creates a CreateSensorUpdatePoliciesBadRequest with default headers values
func NewCreateSensorUpdatePoliciesBadRequest() *CreateSensorUpdatePoliciesBadRequest {
	return &CreateSensorUpdatePoliciesBadRequest{}
}

/*
CreateSensorUpdatePoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateSensorUpdatePoliciesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

// IsSuccess returns true when this create sensor update policies bad request response has a 2xx status code
func (o *CreateSensorUpdatePoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies bad request response has a 3xx status code
func (o *CreateSensorUpdatePoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies bad request response has a 4xx status code
func (o *CreateSensorUpdatePoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sensor update policies bad request response has a 5xx status code
func (o *CreateSensorUpdatePoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies bad request response a status code equal to that given
func (o *CreateSensorUpdatePoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateSensorUpdatePoliciesBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSensorUpdatePoliciesBadRequest) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSensorUpdatePoliciesBadRequest) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSensorUpdatePoliciesForbidden creates a CreateSensorUpdatePoliciesForbidden with default headers values
func NewCreateSensorUpdatePoliciesForbidden() *CreateSensorUpdatePoliciesForbidden {
	return &CreateSensorUpdatePoliciesForbidden{}
}

/*
CreateSensorUpdatePoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateSensorUpdatePoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this create sensor update policies forbidden response has a 2xx status code
func (o *CreateSensorUpdatePoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies forbidden response has a 3xx status code
func (o *CreateSensorUpdatePoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies forbidden response has a 4xx status code
func (o *CreateSensorUpdatePoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sensor update policies forbidden response has a 5xx status code
func (o *CreateSensorUpdatePoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies forbidden response a status code equal to that given
func (o *CreateSensorUpdatePoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateSensorUpdatePoliciesForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *CreateSensorUpdatePoliciesForbidden) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *CreateSensorUpdatePoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSensorUpdatePoliciesNotFound creates a CreateSensorUpdatePoliciesNotFound with default headers values
func NewCreateSensorUpdatePoliciesNotFound() *CreateSensorUpdatePoliciesNotFound {
	return &CreateSensorUpdatePoliciesNotFound{}
}

/*
CreateSensorUpdatePoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateSensorUpdatePoliciesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

// IsSuccess returns true when this create sensor update policies not found response has a 2xx status code
func (o *CreateSensorUpdatePoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies not found response has a 3xx status code
func (o *CreateSensorUpdatePoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies not found response has a 4xx status code
func (o *CreateSensorUpdatePoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sensor update policies not found response has a 5xx status code
func (o *CreateSensorUpdatePoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies not found response a status code equal to that given
func (o *CreateSensorUpdatePoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateSensorUpdatePoliciesNotFound) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesNotFound  %+v", 404, o.Payload)
}

func (o *CreateSensorUpdatePoliciesNotFound) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesNotFound  %+v", 404, o.Payload)
}

func (o *CreateSensorUpdatePoliciesNotFound) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSensorUpdatePoliciesTooManyRequests creates a CreateSensorUpdatePoliciesTooManyRequests with default headers values
func NewCreateSensorUpdatePoliciesTooManyRequests() *CreateSensorUpdatePoliciesTooManyRequests {
	return &CreateSensorUpdatePoliciesTooManyRequests{}
}

/*
CreateSensorUpdatePoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateSensorUpdatePoliciesTooManyRequests struct {

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

// IsSuccess returns true when this create sensor update policies too many requests response has a 2xx status code
func (o *CreateSensorUpdatePoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies too many requests response has a 3xx status code
func (o *CreateSensorUpdatePoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies too many requests response has a 4xx status code
func (o *CreateSensorUpdatePoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create sensor update policies too many requests response has a 5xx status code
func (o *CreateSensorUpdatePoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create sensor update policies too many requests response a status code equal to that given
func (o *CreateSensorUpdatePoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateSensorUpdatePoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSensorUpdatePoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSensorUpdatePoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSensorUpdatePoliciesInternalServerError creates a CreateSensorUpdatePoliciesInternalServerError with default headers values
func NewCreateSensorUpdatePoliciesInternalServerError() *CreateSensorUpdatePoliciesInternalServerError {
	return &CreateSensorUpdatePoliciesInternalServerError{}
}

/*
CreateSensorUpdatePoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateSensorUpdatePoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

// IsSuccess returns true when this create sensor update policies internal server error response has a 2xx status code
func (o *CreateSensorUpdatePoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create sensor update policies internal server error response has a 3xx status code
func (o *CreateSensorUpdatePoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create sensor update policies internal server error response has a 4xx status code
func (o *CreateSensorUpdatePoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create sensor update policies internal server error response has a 5xx status code
func (o *CreateSensorUpdatePoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create sensor update policies internal server error response a status code equal to that given
func (o *CreateSensorUpdatePoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateSensorUpdatePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSensorUpdatePoliciesInternalServerError) String() string {
	return fmt.Sprintf("[POST /policy/entities/sensor-update/v1][%d] createSensorUpdatePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSensorUpdatePoliciesInternalServerError) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *CreateSensorUpdatePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
