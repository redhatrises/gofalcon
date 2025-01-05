// Code generated by go-swagger; DO NOT EDIT.

package downloads_api

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

// DownloadFileReader is a Reader for the DownloadFile structure.
type DownloadFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDownloadFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDownloadFileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /csdownloads/entities/files/download/v1] DownloadFile", response, response.Code())
	}
}

// NewDownloadFileOK creates a DownloadFileOK with default headers values
func NewDownloadFileOK() *DownloadFileOK {
	return &DownloadFileOK{}
}

/*
DownloadFileOK describes a response with status code 200, with default header values.

OK
*/
type DownloadFileOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CommonEntitiesResponse
}

// IsSuccess returns true when this download file o k response has a 2xx status code
func (o *DownloadFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download file o k response has a 3xx status code
func (o *DownloadFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download file o k response has a 4xx status code
func (o *DownloadFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this download file o k response has a 5xx status code
func (o *DownloadFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this download file o k response a status code equal to that given
func (o *DownloadFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the download file o k response
func (o *DownloadFileOK) Code() int {
	return 200
}

func (o *DownloadFileOK) Error() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileOK  %+v", 200, o.Payload)
}

func (o *DownloadFileOK) String() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileOK  %+v", 200, o.Payload)
}

func (o *DownloadFileOK) GetPayload() *models.CommonEntitiesResponse {
	return o.Payload
}

func (o *DownloadFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CommonEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadFileBadRequest creates a DownloadFileBadRequest with default headers values
func NewDownloadFileBadRequest() *DownloadFileBadRequest {
	return &DownloadFileBadRequest{}
}

/*
DownloadFileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DownloadFileBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CommonEntitiesResponse
}

// IsSuccess returns true when this download file bad request response has a 2xx status code
func (o *DownloadFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download file bad request response has a 3xx status code
func (o *DownloadFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download file bad request response has a 4xx status code
func (o *DownloadFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this download file bad request response has a 5xx status code
func (o *DownloadFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this download file bad request response a status code equal to that given
func (o *DownloadFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the download file bad request response
func (o *DownloadFileBadRequest) Code() int {
	return 400
}

func (o *DownloadFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileBadRequest  %+v", 400, o.Payload)
}

func (o *DownloadFileBadRequest) String() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileBadRequest  %+v", 400, o.Payload)
}

func (o *DownloadFileBadRequest) GetPayload() *models.CommonEntitiesResponse {
	return o.Payload
}

func (o *DownloadFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CommonEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadFileUnauthorized creates a DownloadFileUnauthorized with default headers values
func NewDownloadFileUnauthorized() *DownloadFileUnauthorized {
	return &DownloadFileUnauthorized{}
}

/*
DownloadFileUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DownloadFileUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CommonEntitiesResponse
}

// IsSuccess returns true when this download file unauthorized response has a 2xx status code
func (o *DownloadFileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download file unauthorized response has a 3xx status code
func (o *DownloadFileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download file unauthorized response has a 4xx status code
func (o *DownloadFileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this download file unauthorized response has a 5xx status code
func (o *DownloadFileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this download file unauthorized response a status code equal to that given
func (o *DownloadFileUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the download file unauthorized response
func (o *DownloadFileUnauthorized) Code() int {
	return 401
}

func (o *DownloadFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadFileUnauthorized) String() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadFileUnauthorized) GetPayload() *models.CommonEntitiesResponse {
	return o.Payload
}

func (o *DownloadFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CommonEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadFileForbidden creates a DownloadFileForbidden with default headers values
func NewDownloadFileForbidden() *DownloadFileForbidden {
	return &DownloadFileForbidden{}
}

/*
DownloadFileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DownloadFileForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CommonEntitiesResponse
}

// IsSuccess returns true when this download file forbidden response has a 2xx status code
func (o *DownloadFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download file forbidden response has a 3xx status code
func (o *DownloadFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download file forbidden response has a 4xx status code
func (o *DownloadFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this download file forbidden response has a 5xx status code
func (o *DownloadFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this download file forbidden response a status code equal to that given
func (o *DownloadFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the download file forbidden response
func (o *DownloadFileForbidden) Code() int {
	return 403
}

func (o *DownloadFileForbidden) Error() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileForbidden  %+v", 403, o.Payload)
}

func (o *DownloadFileForbidden) String() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileForbidden  %+v", 403, o.Payload)
}

func (o *DownloadFileForbidden) GetPayload() *models.CommonEntitiesResponse {
	return o.Payload
}

func (o *DownloadFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CommonEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadFileTooManyRequests creates a DownloadFileTooManyRequests with default headers values
func NewDownloadFileTooManyRequests() *DownloadFileTooManyRequests {
	return &DownloadFileTooManyRequests{}
}

/*
DownloadFileTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DownloadFileTooManyRequests struct {

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

// IsSuccess returns true when this download file too many requests response has a 2xx status code
func (o *DownloadFileTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download file too many requests response has a 3xx status code
func (o *DownloadFileTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download file too many requests response has a 4xx status code
func (o *DownloadFileTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this download file too many requests response has a 5xx status code
func (o *DownloadFileTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this download file too many requests response a status code equal to that given
func (o *DownloadFileTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the download file too many requests response
func (o *DownloadFileTooManyRequests) Code() int {
	return 429
}

func (o *DownloadFileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileTooManyRequests  %+v", 429, o.Payload)
}

func (o *DownloadFileTooManyRequests) String() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileTooManyRequests  %+v", 429, o.Payload)
}

func (o *DownloadFileTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DownloadFileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDownloadFileInternalServerError creates a DownloadFileInternalServerError with default headers values
func NewDownloadFileInternalServerError() *DownloadFileInternalServerError {
	return &DownloadFileInternalServerError{}
}

/*
DownloadFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DownloadFileInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CommonEntitiesResponse
}

// IsSuccess returns true when this download file internal server error response has a 2xx status code
func (o *DownloadFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download file internal server error response has a 3xx status code
func (o *DownloadFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download file internal server error response has a 4xx status code
func (o *DownloadFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this download file internal server error response has a 5xx status code
func (o *DownloadFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this download file internal server error response a status code equal to that given
func (o *DownloadFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the download file internal server error response
func (o *DownloadFileInternalServerError) Code() int {
	return 500
}

func (o *DownloadFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadFileInternalServerError) String() string {
	return fmt.Sprintf("[GET /csdownloads/entities/files/download/v1][%d] downloadFileInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadFileInternalServerError) GetPayload() *models.CommonEntitiesResponse {
	return o.Payload
}

func (o *DownloadFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CommonEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}