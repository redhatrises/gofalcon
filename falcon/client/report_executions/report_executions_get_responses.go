// Code generated by go-swagger; DO NOT EDIT.

package report_executions

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

// ReportExecutionsGetReader is a Reader for the ReportExecutionsGet structure.
type ReportExecutionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportExecutionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportExecutionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReportExecutionsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReportExecutionsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReportExecutionsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReportExecutionsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReportExecutionsGetOK creates a ReportExecutionsGetOK with default headers values
func NewReportExecutionsGetOK() *ReportExecutionsGetOK {
	return &ReportExecutionsGetOK{}
}

/*
ReportExecutionsGetOK describes a response with status code 200, with default header values.

OK
*/
type ReportExecutionsGetOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIReportExecutionsResponseV1
}

// IsSuccess returns true when this report executions get o k response has a 2xx status code
func (o *ReportExecutionsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this report executions get o k response has a 3xx status code
func (o *ReportExecutionsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions get o k response has a 4xx status code
func (o *ReportExecutionsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this report executions get o k response has a 5xx status code
func (o *ReportExecutionsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions get o k response a status code equal to that given
func (o *ReportExecutionsGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ReportExecutionsGetOK) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] reportExecutionsGetOK  %+v", 200, o.Payload)
}

func (o *ReportExecutionsGetOK) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] reportExecutionsGetOK  %+v", 200, o.Payload)
}

func (o *ReportExecutionsGetOK) GetPayload() *models.APIReportExecutionsResponseV1 {
	return o.Payload
}

func (o *ReportExecutionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIReportExecutionsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportExecutionsGetBadRequest creates a ReportExecutionsGetBadRequest with default headers values
func NewReportExecutionsGetBadRequest() *ReportExecutionsGetBadRequest {
	return &ReportExecutionsGetBadRequest{}
}

/*
ReportExecutionsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ReportExecutionsGetBadRequest struct {

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

// IsSuccess returns true when this report executions get bad request response has a 2xx status code
func (o *ReportExecutionsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions get bad request response has a 3xx status code
func (o *ReportExecutionsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions get bad request response has a 4xx status code
func (o *ReportExecutionsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions get bad request response has a 5xx status code
func (o *ReportExecutionsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions get bad request response a status code equal to that given
func (o *ReportExecutionsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ReportExecutionsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] reportExecutionsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ReportExecutionsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] reportExecutionsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ReportExecutionsGetBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsGetForbidden creates a ReportExecutionsGetForbidden with default headers values
func NewReportExecutionsGetForbidden() *ReportExecutionsGetForbidden {
	return &ReportExecutionsGetForbidden{}
}

/*
ReportExecutionsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReportExecutionsGetForbidden struct {

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

// IsSuccess returns true when this report executions get forbidden response has a 2xx status code
func (o *ReportExecutionsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions get forbidden response has a 3xx status code
func (o *ReportExecutionsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions get forbidden response has a 4xx status code
func (o *ReportExecutionsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions get forbidden response has a 5xx status code
func (o *ReportExecutionsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions get forbidden response a status code equal to that given
func (o *ReportExecutionsGetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ReportExecutionsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] reportExecutionsGetForbidden  %+v", 403, o.Payload)
}

func (o *ReportExecutionsGetForbidden) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] reportExecutionsGetForbidden  %+v", 403, o.Payload)
}

func (o *ReportExecutionsGetForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsGetTooManyRequests creates a ReportExecutionsGetTooManyRequests with default headers values
func NewReportExecutionsGetTooManyRequests() *ReportExecutionsGetTooManyRequests {
	return &ReportExecutionsGetTooManyRequests{}
}

/*
ReportExecutionsGetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReportExecutionsGetTooManyRequests struct {

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

// IsSuccess returns true when this report executions get too many requests response has a 2xx status code
func (o *ReportExecutionsGetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report executions get too many requests response has a 3xx status code
func (o *ReportExecutionsGetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report executions get too many requests response has a 4xx status code
func (o *ReportExecutionsGetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this report executions get too many requests response has a 5xx status code
func (o *ReportExecutionsGetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this report executions get too many requests response a status code equal to that given
func (o *ReportExecutionsGetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ReportExecutionsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] reportExecutionsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReportExecutionsGetTooManyRequests) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] reportExecutionsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReportExecutionsGetTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReportExecutionsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReportExecutionsGetDefault creates a ReportExecutionsGetDefault with default headers values
func NewReportExecutionsGetDefault(code int) *ReportExecutionsGetDefault {
	return &ReportExecutionsGetDefault{
		_statusCode: code,
	}
}

/*
ReportExecutionsGetDefault describes a response with status code -1, with default header values.

OK
*/
type ReportExecutionsGetDefault struct {
	_statusCode int

	Payload *models.APIReportExecutionsResponseV1
}

// Code gets the status code for the report executions get default response
func (o *ReportExecutionsGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this report executions get default response has a 2xx status code
func (o *ReportExecutionsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this report executions get default response has a 3xx status code
func (o *ReportExecutionsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this report executions get default response has a 4xx status code
func (o *ReportExecutionsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this report executions get default response has a 5xx status code
func (o *ReportExecutionsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this report executions get default response a status code equal to that given
func (o *ReportExecutionsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ReportExecutionsGetDefault) Error() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] report-executions.get default  %+v", o._statusCode, o.Payload)
}

func (o *ReportExecutionsGetDefault) String() string {
	return fmt.Sprintf("[GET /reports/entities/report-executions/v1][%d] report-executions.get default  %+v", o._statusCode, o.Payload)
}

func (o *ReportExecutionsGetDefault) GetPayload() *models.APIReportExecutionsResponseV1 {
	return o.Payload
}

func (o *ReportExecutionsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIReportExecutionsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
