// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

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

// AggregateFCIncidentsReader is a Reader for the AggregateFCIncidents structure.
type AggregateFCIncidentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateFCIncidentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateFCIncidentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregateFCIncidentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateFCIncidentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregateFCIncidentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1] AggregateFCIncidents", response, response.Code())
	}
}

// NewAggregateFCIncidentsOK creates a AggregateFCIncidentsOK with default headers values
func NewAggregateFCIncidentsOK() *AggregateFCIncidentsOK {
	return &AggregateFCIncidentsOK{}
}

/*
AggregateFCIncidentsOK describes a response with status code 200, with default header values.

OK
*/
type AggregateFCIncidentsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this aggregate f c incidents o k response has a 2xx status code
func (o *AggregateFCIncidentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate f c incidents o k response has a 3xx status code
func (o *AggregateFCIncidentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate f c incidents o k response has a 4xx status code
func (o *AggregateFCIncidentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate f c incidents o k response has a 5xx status code
func (o *AggregateFCIncidentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate f c incidents o k response a status code equal to that given
func (o *AggregateFCIncidentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate f c incidents o k response
func (o *AggregateFCIncidentsOK) Code() int {
	return 200
}

func (o *AggregateFCIncidentsOK) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1][%d] aggregateFCIncidentsOK  %+v", 200, o.Payload)
}

func (o *AggregateFCIncidentsOK) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1][%d] aggregateFCIncidentsOK  %+v", 200, o.Payload)
}

func (o *AggregateFCIncidentsOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateFCIncidentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateFCIncidentsForbidden creates a AggregateFCIncidentsForbidden with default headers values
func NewAggregateFCIncidentsForbidden() *AggregateFCIncidentsForbidden {
	return &AggregateFCIncidentsForbidden{}
}

/*
AggregateFCIncidentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateFCIncidentsForbidden struct {

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

// IsSuccess returns true when this aggregate f c incidents forbidden response has a 2xx status code
func (o *AggregateFCIncidentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate f c incidents forbidden response has a 3xx status code
func (o *AggregateFCIncidentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate f c incidents forbidden response has a 4xx status code
func (o *AggregateFCIncidentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate f c incidents forbidden response has a 5xx status code
func (o *AggregateFCIncidentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate f c incidents forbidden response a status code equal to that given
func (o *AggregateFCIncidentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate f c incidents forbidden response
func (o *AggregateFCIncidentsForbidden) Code() int {
	return 403
}

func (o *AggregateFCIncidentsForbidden) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1][%d] aggregateFCIncidentsForbidden  %+v", 403, o.Payload)
}

func (o *AggregateFCIncidentsForbidden) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1][%d] aggregateFCIncidentsForbidden  %+v", 403, o.Payload)
}

func (o *AggregateFCIncidentsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateFCIncidentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateFCIncidentsTooManyRequests creates a AggregateFCIncidentsTooManyRequests with default headers values
func NewAggregateFCIncidentsTooManyRequests() *AggregateFCIncidentsTooManyRequests {
	return &AggregateFCIncidentsTooManyRequests{}
}

/*
AggregateFCIncidentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateFCIncidentsTooManyRequests struct {

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

// IsSuccess returns true when this aggregate f c incidents too many requests response has a 2xx status code
func (o *AggregateFCIncidentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate f c incidents too many requests response has a 3xx status code
func (o *AggregateFCIncidentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate f c incidents too many requests response has a 4xx status code
func (o *AggregateFCIncidentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate f c incidents too many requests response has a 5xx status code
func (o *AggregateFCIncidentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate f c incidents too many requests response a status code equal to that given
func (o *AggregateFCIncidentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate f c incidents too many requests response
func (o *AggregateFCIncidentsTooManyRequests) Code() int {
	return 429
}

func (o *AggregateFCIncidentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1][%d] aggregateFCIncidentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateFCIncidentsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1][%d] aggregateFCIncidentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateFCIncidentsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateFCIncidentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateFCIncidentsInternalServerError creates a AggregateFCIncidentsInternalServerError with default headers values
func NewAggregateFCIncidentsInternalServerError() *AggregateFCIncidentsInternalServerError {
	return &AggregateFCIncidentsInternalServerError{}
}

/*
AggregateFCIncidentsInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type AggregateFCIncidentsInternalServerError struct {

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

// IsSuccess returns true when this aggregate f c incidents internal server error response has a 2xx status code
func (o *AggregateFCIncidentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate f c incidents internal server error response has a 3xx status code
func (o *AggregateFCIncidentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate f c incidents internal server error response has a 4xx status code
func (o *AggregateFCIncidentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate f c incidents internal server error response has a 5xx status code
func (o *AggregateFCIncidentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate f c incidents internal server error response a status code equal to that given
func (o *AggregateFCIncidentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate f c incidents internal server error response
func (o *AggregateFCIncidentsInternalServerError) Code() int {
	return 500
}

func (o *AggregateFCIncidentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1][%d] aggregateFCIncidentsInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateFCIncidentsInternalServerError) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/incidents/GET/v1][%d] aggregateFCIncidentsInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateFCIncidentsInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateFCIncidentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
