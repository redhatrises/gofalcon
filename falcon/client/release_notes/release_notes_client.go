// Code generated by go-swagger; DO NOT EDIT.

package release_notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new release notes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for release notes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CombinedReleaseNotesV1(params *CombinedReleaseNotesV1Params, opts ...ClientOption) (*CombinedReleaseNotesV1OK, error)

	GetEntityIDsByQueryPOST(params *GetEntityIDsByQueryPOSTParams, opts ...ClientOption) (*GetEntityIDsByQueryPOSTOK, error)

	QueryReleaseNotesV1(params *QueryReleaseNotesV1Params, opts ...ClientOption) (*QueryReleaseNotesV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CombinedReleaseNotesV1 queries for release notes resources and returns details
*/
func (a *Client) CombinedReleaseNotesV1(params *CombinedReleaseNotesV1Params, opts ...ClientOption) (*CombinedReleaseNotesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCombinedReleaseNotesV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "CombinedReleaseNotesV1",
		Method:             "GET",
		PathPattern:        "/deployment-coordinator/combined/release-notes/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CombinedReleaseNotesV1Reader{formats: a.formats},
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
	success, ok := result.(*CombinedReleaseNotesV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CombinedReleaseNotesV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEntityIDsByQueryPOST returns the i ds of all entities in the database for the given page
*/
func (a *Client) GetEntityIDsByQueryPOST(params *GetEntityIDsByQueryPOSTParams, opts ...ClientOption) (*GetEntityIDsByQueryPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntityIDsByQueryPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEntityIDsByQueryPOST",
		Method:             "POST",
		PathPattern:        "/deployment-coordinator/entities/release-notes/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntityIDsByQueryPOSTReader{formats: a.formats},
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
	success, ok := result.(*GetEntityIDsByQueryPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEntityIDsByQueryPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryReleaseNotesV1 queries for release notes resources and returns ids
*/
func (a *Client) QueryReleaseNotesV1(params *QueryReleaseNotesV1Params, opts ...ClientOption) (*QueryReleaseNotesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryReleaseNotesV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryReleaseNotesV1",
		Method:             "GET",
		PathPattern:        "/deployment-coordinator/queries/release-notes/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryReleaseNotesV1Reader{formats: a.formats},
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
	success, ok := result.(*QueryReleaseNotesV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryReleaseNotesV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}