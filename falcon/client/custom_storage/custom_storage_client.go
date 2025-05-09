// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new custom storage API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom storage API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteObject(params *DeleteObjectParams, opts ...ClientOption) (*DeleteObjectOK, error)

	DeleteVersionedObject(params *DeleteVersionedObjectParams, opts ...ClientOption) (*DeleteVersionedObjectOK, error)

	DescribeCollection(params *DescribeCollectionParams, opts ...ClientOption) (*DescribeCollectionOK, error)

	DescribeCollections(params *DescribeCollectionsParams, opts ...ClientOption) (*DescribeCollectionsOK, error)

	GetObject(params *GetObjectParams, writer io.Writer, opts ...ClientOption) (*GetObjectOK, error)

	GetObjectMetadata(params *GetObjectMetadataParams, opts ...ClientOption) (*GetObjectMetadataOK, error)

	GetSchema(params *GetSchemaParams, writer io.Writer, opts ...ClientOption) (*GetSchemaOK, error)

	GetSchemaMetadata(params *GetSchemaMetadataParams, opts ...ClientOption) (*GetSchemaMetadataOK, error)

	GetVersionedObject(params *GetVersionedObjectParams, writer io.Writer, opts ...ClientOption) (*GetVersionedObjectOK, error)

	GetVersionedObjectMetadata(params *GetVersionedObjectMetadataParams, opts ...ClientOption) (*GetVersionedObjectMetadataOK, error)

	ListCollections(params *ListCollectionsParams, opts ...ClientOption) (*ListCollectionsOK, error)

	ListObjects(params *ListObjectsParams, opts ...ClientOption) (*ListObjectsOK, error)

	ListObjectsByVersion(params *ListObjectsByVersionParams, opts ...ClientOption) (*ListObjectsByVersionOK, error)

	ListSchemas(params *ListSchemasParams, opts ...ClientOption) (*ListSchemasOK, error)

	PutObject(params *PutObjectParams, opts ...ClientOption) (*PutObjectOK, error)

	PutObjectByVersion(params *PutObjectByVersionParams, opts ...ClientOption) (*PutObjectByVersionOK, error)

	SearchObjects(params *SearchObjectsParams, opts ...ClientOption) (*SearchObjectsOK, error)

	SearchObjectsByVersion(params *SearchObjectsByVersionParams, opts ...ClientOption) (*SearchObjectsByVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteObject deletes the specified object
*/
func (a *Client) DeleteObject(params *DeleteObjectParams, opts ...ClientOption) (*DeleteObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteObject",
		Method:             "DELETE",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteObjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteVersionedObject deletes the specified versioned object
*/
func (a *Client) DeleteVersionedObject(params *DeleteVersionedObjectParams, opts ...ClientOption) (*DeleteVersionedObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionedObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteVersionedObject",
		Method:             "DELETE",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVersionedObjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVersionedObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteVersionedObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DescribeCollection fetches metadata about an existing collection
*/
func (a *Client) DescribeCollection(params *DescribeCollectionParams, opts ...ClientOption) (*DescribeCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DescribeCollection",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DescribeCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DescribeCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DescribeCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DescribeCollections fetches metadata about one or more existing collections
*/
func (a *Client) DescribeCollections(params *DescribeCollectionsParams, opts ...ClientOption) (*DescribeCollectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeCollectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DescribeCollections",
		Method:             "PUT",
		PathPattern:        "/customobjects/v1/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DescribeCollectionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DescribeCollectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DescribeCollections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetObject gets the bytes for the specified object
*/
func (a *Client) GetObject(params *GetObjectParams, writer io.Writer, opts ...ClientOption) (*GetObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetObject",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetObjectReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetObjectMetadata gets the metadata for the specified object
*/
func (a *Client) GetObjectMetadata(params *GetObjectMetadataParams, opts ...ClientOption) (*GetObjectMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetObjectMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetObjectMetadata",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects/{object_key}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetObjectMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetObjectMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetObjectMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSchema gets the bytes of the specified schema of the requested collection
*/
func (a *Client) GetSchema(params *GetSchemaParams, writer io.Writer, opts ...ClientOption) (*GetSchemaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSchema",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/schemas/{schema_version}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream", "application/schema+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSchemaReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSchema: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSchemaMetadata gets the metadata for the specified schema of the requested collection
*/
func (a *Client) GetSchemaMetadata(params *GetSchemaMetadataParams, opts ...ClientOption) (*GetSchemaMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchemaMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSchemaMetadata",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/schemas/{schema_version}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSchemaMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSchemaMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSchemaMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVersionedObject gets the bytes for the specified object
*/
func (a *Client) GetVersionedObject(params *GetVersionedObjectParams, writer io.Writer, opts ...ClientOption) (*GetVersionedObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionedObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVersionedObject",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionedObjectReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVersionedObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVersionedObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVersionedObjectMetadata gets the metadata for the specified object
*/
func (a *Client) GetVersionedObjectMetadata(params *GetVersionedObjectMetadataParams, opts ...ClientOption) (*GetVersionedObjectMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionedObjectMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVersionedObjectMetadata",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects/{object_key}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionedObjectMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVersionedObjectMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVersionedObjectMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListCollections lists available collection names in alphabetical order
*/
func (a *Client) ListCollections(params *ListCollectionsParams, opts ...ClientOption) (*ListCollectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCollectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListCollections",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCollectionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCollectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCollections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListObjects lists the object keys in the specified collection in alphabetical order
*/
func (a *Client) ListObjects(params *ListObjectsParams, opts ...ClientOption) (*ListObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListObjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListObjects",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListObjectsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListObjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListObjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListObjectsByVersion lists the object keys in the specified collection in alphabetical order
*/
func (a *Client) ListObjectsByVersion(params *ListObjectsByVersionParams, opts ...ClientOption) (*ListObjectsByVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListObjectsByVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListObjectsByVersion",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListObjectsByVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListObjectsByVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListObjectsByVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListSchemas gets the list of schemas for the requested collection in reverse version order latest first
*/
func (a *Client) ListSchemas(params *ListSchemasParams, opts ...ClientOption) (*ListSchemasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSchemasParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListSchemas",
		Method:             "GET",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/schemas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSchemasReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSchemasOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListSchemas: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutObject puts the specified new object at the given key or overwrite an existing object at the given key
*/
func (a *Client) PutObject(params *PutObjectParams, opts ...ClientOption) (*PutObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutObject",
		Method:             "PUT",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutObjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutObjectByVersion puts the specified new object at the given key or overwrite an existing object at the given key
*/
func (a *Client) PutObjectByVersion(params *PutObjectByVersionParams, opts ...ClientOption) (*PutObjectByVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutObjectByVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutObjectByVersion",
		Method:             "PUT",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects/{object_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutObjectByVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutObjectByVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutObjectByVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchObjects searches for objects that match the specified filter criteria returns metadata not actual objects
*/
func (a *Client) SearchObjects(params *SearchObjectsParams, opts ...ClientOption) (*SearchObjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchObjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchObjects",
		Method:             "POST",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchObjectsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchObjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchObjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchObjectsByVersion searches for objects that match the specified filter criteria returns metadata not actual objects
*/
func (a *Client) SearchObjectsByVersion(params *SearchObjectsByVersionParams, opts ...ClientOption) (*SearchObjectsByVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchObjectsByVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchObjectsByVersion",
		Method:             "POST",
		PathPattern:        "/customobjects/v1/collections/{collection_name}/{collection_version}/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchObjectsByVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchObjectsByVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchObjectsByVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
