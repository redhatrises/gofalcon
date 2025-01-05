// Code generated by go-swagger; DO NOT EDIT.

package falcon_container_image

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

// ReadExportJobsReader is a Reader for the ReadExportJobs structure.
type ReadExportJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadExportJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadExportJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadExportJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadExportJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReadExportJobsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadExportJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/entities/exports/v1] ReadExportJobs", response, response.Code())
	}
}

// NewReadExportJobsOK creates a ReadExportJobsOK with default headers values
func NewReadExportJobsOK() *ReadExportJobsOK {
	return &ReadExportJobsOK{}
}

/*
ReadExportJobsOK describes a response with status code 200, with default header values.

OK
*/
type ReadExportJobsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ExportsExportsResponse
}

// IsSuccess returns true when this read export jobs o k response has a 2xx status code
func (o *ReadExportJobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read export jobs o k response has a 3xx status code
func (o *ReadExportJobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read export jobs o k response has a 4xx status code
func (o *ReadExportJobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read export jobs o k response has a 5xx status code
func (o *ReadExportJobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read export jobs o k response a status code equal to that given
func (o *ReadExportJobsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read export jobs o k response
func (o *ReadExportJobsOK) Code() int {
	return 200
}

func (o *ReadExportJobsOK) Error() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsOK  %+v", 200, o.Payload)
}

func (o *ReadExportJobsOK) String() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsOK  %+v", 200, o.Payload)
}

func (o *ReadExportJobsOK) GetPayload() *models.ExportsExportsResponse {
	return o.Payload
}

func (o *ReadExportJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ExportsExportsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadExportJobsBadRequest creates a ReadExportJobsBadRequest with default headers values
func NewReadExportJobsBadRequest() *ReadExportJobsBadRequest {
	return &ReadExportJobsBadRequest{}
}

/*
ReadExportJobsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ReadExportJobsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this read export jobs bad request response has a 2xx status code
func (o *ReadExportJobsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read export jobs bad request response has a 3xx status code
func (o *ReadExportJobsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read export jobs bad request response has a 4xx status code
func (o *ReadExportJobsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this read export jobs bad request response has a 5xx status code
func (o *ReadExportJobsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this read export jobs bad request response a status code equal to that given
func (o *ReadExportJobsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the read export jobs bad request response
func (o *ReadExportJobsBadRequest) Code() int {
	return 400
}

func (o *ReadExportJobsBadRequest) Error() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsBadRequest  %+v", 400, o.Payload)
}

func (o *ReadExportJobsBadRequest) String() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsBadRequest  %+v", 400, o.Payload)
}

func (o *ReadExportJobsBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ReadExportJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadExportJobsForbidden creates a ReadExportJobsForbidden with default headers values
func NewReadExportJobsForbidden() *ReadExportJobsForbidden {
	return &ReadExportJobsForbidden{}
}

/*
ReadExportJobsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReadExportJobsForbidden struct {

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

// IsSuccess returns true when this read export jobs forbidden response has a 2xx status code
func (o *ReadExportJobsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read export jobs forbidden response has a 3xx status code
func (o *ReadExportJobsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read export jobs forbidden response has a 4xx status code
func (o *ReadExportJobsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this read export jobs forbidden response has a 5xx status code
func (o *ReadExportJobsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this read export jobs forbidden response a status code equal to that given
func (o *ReadExportJobsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the read export jobs forbidden response
func (o *ReadExportJobsForbidden) Code() int {
	return 403
}

func (o *ReadExportJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsForbidden  %+v", 403, o.Payload)
}

func (o *ReadExportJobsForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsForbidden  %+v", 403, o.Payload)
}

func (o *ReadExportJobsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadExportJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadExportJobsTooManyRequests creates a ReadExportJobsTooManyRequests with default headers values
func NewReadExportJobsTooManyRequests() *ReadExportJobsTooManyRequests {
	return &ReadExportJobsTooManyRequests{}
}

/*
ReadExportJobsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReadExportJobsTooManyRequests struct {

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

// IsSuccess returns true when this read export jobs too many requests response has a 2xx status code
func (o *ReadExportJobsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read export jobs too many requests response has a 3xx status code
func (o *ReadExportJobsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read export jobs too many requests response has a 4xx status code
func (o *ReadExportJobsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this read export jobs too many requests response has a 5xx status code
func (o *ReadExportJobsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this read export jobs too many requests response a status code equal to that given
func (o *ReadExportJobsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the read export jobs too many requests response
func (o *ReadExportJobsTooManyRequests) Code() int {
	return 429
}

func (o *ReadExportJobsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadExportJobsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadExportJobsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadExportJobsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadExportJobsInternalServerError creates a ReadExportJobsInternalServerError with default headers values
func NewReadExportJobsInternalServerError() *ReadExportJobsInternalServerError {
	return &ReadExportJobsInternalServerError{}
}

/*
ReadExportJobsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReadExportJobsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this read export jobs internal server error response has a 2xx status code
func (o *ReadExportJobsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read export jobs internal server error response has a 3xx status code
func (o *ReadExportJobsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read export jobs internal server error response has a 4xx status code
func (o *ReadExportJobsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read export jobs internal server error response has a 5xx status code
func (o *ReadExportJobsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read export jobs internal server error response a status code equal to that given
func (o *ReadExportJobsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read export jobs internal server error response
func (o *ReadExportJobsInternalServerError) Code() int {
	return 500
}

func (o *ReadExportJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadExportJobsInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/entities/exports/v1][%d] readExportJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadExportJobsInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ReadExportJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}