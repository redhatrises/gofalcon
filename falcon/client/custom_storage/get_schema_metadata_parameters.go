// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetSchemaMetadataParams creates a new GetSchemaMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSchemaMetadataParams() *GetSchemaMetadataParams {
	return &GetSchemaMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchemaMetadataParamsWithTimeout creates a new GetSchemaMetadataParams object
// with the ability to set a timeout on a request.
func NewGetSchemaMetadataParamsWithTimeout(timeout time.Duration) *GetSchemaMetadataParams {
	return &GetSchemaMetadataParams{
		timeout: timeout,
	}
}

// NewGetSchemaMetadataParamsWithContext creates a new GetSchemaMetadataParams object
// with the ability to set a context for a request.
func NewGetSchemaMetadataParamsWithContext(ctx context.Context) *GetSchemaMetadataParams {
	return &GetSchemaMetadataParams{
		Context: ctx,
	}
}

// NewGetSchemaMetadataParamsWithHTTPClient creates a new GetSchemaMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSchemaMetadataParamsWithHTTPClient(client *http.Client) *GetSchemaMetadataParams {
	return &GetSchemaMetadataParams{
		HTTPClient: client,
	}
}

/*
GetSchemaMetadataParams contains all the parameters to send to the API endpoint

	for the get schema metadata operation.

	Typically these are written to a http.Request.
*/
type GetSchemaMetadataParams struct {

	/* CollectionName.

	   The name of the collection
	*/
	CollectionName string

	/* SchemaVersion.

	   The version of the collection schema or 'latest' for the latest version
	*/
	SchemaVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schema metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchemaMetadataParams) WithDefaults() *GetSchemaMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schema metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchemaMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schema metadata params
func (o *GetSchemaMetadataParams) WithTimeout(timeout time.Duration) *GetSchemaMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schema metadata params
func (o *GetSchemaMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schema metadata params
func (o *GetSchemaMetadataParams) WithContext(ctx context.Context) *GetSchemaMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schema metadata params
func (o *GetSchemaMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schema metadata params
func (o *GetSchemaMetadataParams) WithHTTPClient(client *http.Client) *GetSchemaMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schema metadata params
func (o *GetSchemaMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionName adds the collectionName to the get schema metadata params
func (o *GetSchemaMetadataParams) WithCollectionName(collectionName string) *GetSchemaMetadataParams {
	o.SetCollectionName(collectionName)
	return o
}

// SetCollectionName adds the collectionName to the get schema metadata params
func (o *GetSchemaMetadataParams) SetCollectionName(collectionName string) {
	o.CollectionName = collectionName
}

// WithSchemaVersion adds the schemaVersion to the get schema metadata params
func (o *GetSchemaMetadataParams) WithSchemaVersion(schemaVersion string) *GetSchemaMetadataParams {
	o.SetSchemaVersion(schemaVersion)
	return o
}

// SetSchemaVersion adds the schemaVersion to the get schema metadata params
func (o *GetSchemaMetadataParams) SetSchemaVersion(schemaVersion string) {
	o.SchemaVersion = schemaVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemaMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection_name
	if err := r.SetPathParam("collection_name", o.CollectionName); err != nil {
		return err
	}

	// path param schema_version
	if err := r.SetPathParam("schema_version", o.SchemaVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
