// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

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

// DescribeCollectionsReader is a Reader for the DescribeCollections structure.
type DescribeCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDescribeCollectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDescribeCollectionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDescribeCollectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /customobjects/v1/collections] DescribeCollections", response, response.Code())
	}
}

// NewDescribeCollectionsOK creates a DescribeCollectionsOK with default headers values
func NewDescribeCollectionsOK() *DescribeCollectionsOK {
	return &DescribeCollectionsOK{}
}

/*
DescribeCollectionsOK describes a response with status code 200, with default header values.

OK
*/
type DescribeCollectionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CustomType1942251022
}

// IsSuccess returns true when this describe collections o k response has a 2xx status code
func (o *DescribeCollectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe collections o k response has a 3xx status code
func (o *DescribeCollectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe collections o k response has a 4xx status code
func (o *DescribeCollectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe collections o k response has a 5xx status code
func (o *DescribeCollectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe collections o k response a status code equal to that given
func (o *DescribeCollectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe collections o k response
func (o *DescribeCollectionsOK) Code() int {
	return 200
}

func (o *DescribeCollectionsOK) Error() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections][%d] describeCollectionsOK  %+v", 200, o.Payload)
}

func (o *DescribeCollectionsOK) String() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections][%d] describeCollectionsOK  %+v", 200, o.Payload)
}

func (o *DescribeCollectionsOK) GetPayload() *models.CustomType1942251022 {
	return o.Payload
}

func (o *DescribeCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CustomType1942251022)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeCollectionsForbidden creates a DescribeCollectionsForbidden with default headers values
func NewDescribeCollectionsForbidden() *DescribeCollectionsForbidden {
	return &DescribeCollectionsForbidden{}
}

/*
DescribeCollectionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DescribeCollectionsForbidden struct {

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

// IsSuccess returns true when this describe collections forbidden response has a 2xx status code
func (o *DescribeCollectionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe collections forbidden response has a 3xx status code
func (o *DescribeCollectionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe collections forbidden response has a 4xx status code
func (o *DescribeCollectionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe collections forbidden response has a 5xx status code
func (o *DescribeCollectionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this describe collections forbidden response a status code equal to that given
func (o *DescribeCollectionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the describe collections forbidden response
func (o *DescribeCollectionsForbidden) Code() int {
	return 403
}

func (o *DescribeCollectionsForbidden) Error() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections][%d] describeCollectionsForbidden  %+v", 403, o.Payload)
}

func (o *DescribeCollectionsForbidden) String() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections][%d] describeCollectionsForbidden  %+v", 403, o.Payload)
}

func (o *DescribeCollectionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DescribeCollectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDescribeCollectionsTooManyRequests creates a DescribeCollectionsTooManyRequests with default headers values
func NewDescribeCollectionsTooManyRequests() *DescribeCollectionsTooManyRequests {
	return &DescribeCollectionsTooManyRequests{}
}

/*
DescribeCollectionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DescribeCollectionsTooManyRequests struct {

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

// IsSuccess returns true when this describe collections too many requests response has a 2xx status code
func (o *DescribeCollectionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe collections too many requests response has a 3xx status code
func (o *DescribeCollectionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe collections too many requests response has a 4xx status code
func (o *DescribeCollectionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this describe collections too many requests response has a 5xx status code
func (o *DescribeCollectionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this describe collections too many requests response a status code equal to that given
func (o *DescribeCollectionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the describe collections too many requests response
func (o *DescribeCollectionsTooManyRequests) Code() int {
	return 429
}

func (o *DescribeCollectionsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections][%d] describeCollectionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DescribeCollectionsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections][%d] describeCollectionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DescribeCollectionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DescribeCollectionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDescribeCollectionsInternalServerError creates a DescribeCollectionsInternalServerError with default headers values
func NewDescribeCollectionsInternalServerError() *DescribeCollectionsInternalServerError {
	return &DescribeCollectionsInternalServerError{}
}

/*
DescribeCollectionsInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type DescribeCollectionsInternalServerError struct {

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

// IsSuccess returns true when this describe collections internal server error response has a 2xx status code
func (o *DescribeCollectionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this describe collections internal server error response has a 3xx status code
func (o *DescribeCollectionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe collections internal server error response has a 4xx status code
func (o *DescribeCollectionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe collections internal server error response has a 5xx status code
func (o *DescribeCollectionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this describe collections internal server error response a status code equal to that given
func (o *DescribeCollectionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the describe collections internal server error response
func (o *DescribeCollectionsInternalServerError) Code() int {
	return 500
}

func (o *DescribeCollectionsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections][%d] describeCollectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeCollectionsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections][%d] describeCollectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeCollectionsInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DescribeCollectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
