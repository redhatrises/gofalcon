// Code generated by go-swagger; DO NOT EDIT.

package jira

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new jira API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for jira API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AggregateJiraAllowList(params *AggregateJiraAllowListParams) (*AggregateJiraAllowListOK, error)

	AggregateJiraBlockList(params *AggregateJiraBlockListParams) (*AggregateJiraBlockListOK, error)

	AggregateJiraEscalations(params *AggregateJiraEscalationsParams) (*AggregateJiraEscalationsOK, error)

	AggregateJiraRemediations(params *AggregateJiraRemediationsParams) (*AggregateJiraRemediationsOK, error)

	QueryJiraAllowListFilter(params *QueryJiraAllowListFilterParams) (*QueryJiraAllowListFilterOK, error)

	QueryJiraBlockListFilter(params *QueryJiraBlockListFilterParams) (*QueryJiraBlockListFilterOK, error)

	QueryJiraEscalationsFilter(params *QueryJiraEscalationsFilterParams) (*QueryJiraEscalationsFilterOK, error)

	QueryJiraRemediationsFilter(params *QueryJiraRemediationsFilterParams) (*QueryJiraRemediationsFilterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AggregateJiraAllowList retrieves aggregate jira allow list ticket values based on the matched filter
*/
func (a *Client) AggregateJiraAllowList(params *AggregateJiraAllowListParams) (*AggregateJiraAllowListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateJiraAllowListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AggregateJiraAllowList",
		Method:             "POST",
		PathPattern:        "/falcon-complete-dashboards/aggregates/jira-allowlist/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateJiraAllowListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AggregateJiraAllowListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AggregateJiraAllowListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AggregateJiraBlockList retrieves aggregate jira block list ticket values based on the matched filter
*/
func (a *Client) AggregateJiraBlockList(params *AggregateJiraBlockListParams) (*AggregateJiraBlockListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateJiraBlockListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AggregateJiraBlockList",
		Method:             "POST",
		PathPattern:        "/falcon-complete-dashboards/aggregates/jira-blocklist/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateJiraBlockListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AggregateJiraBlockListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AggregateJiraBlockListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AggregateJiraEscalations retrieves aggregate jira escalation ticket values based on the matched filter
*/
func (a *Client) AggregateJiraEscalations(params *AggregateJiraEscalationsParams) (*AggregateJiraEscalationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateJiraEscalationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AggregateJiraEscalations",
		Method:             "POST",
		PathPattern:        "/falcon-complete-dashboards/aggregates/jira-escalations/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateJiraEscalationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AggregateJiraEscalationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AggregateJiraEscalationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AggregateJiraRemediations retrieves aggregate jira remediation ticket values based on the matched filter
*/
func (a *Client) AggregateJiraRemediations(params *AggregateJiraRemediationsParams) (*AggregateJiraRemediationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateJiraRemediationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AggregateJiraRemediations",
		Method:             "POST",
		PathPattern:        "/falcon-complete-dashboards/aggregates/jira-remediations/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateJiraRemediationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AggregateJiraRemediationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AggregateJiraRemediationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryJiraAllowListFilter retrieves jira allow list tickets that match the provided filter criteria with scrolling enabled
*/
func (a *Client) QueryJiraAllowListFilter(params *QueryJiraAllowListFilterParams) (*QueryJiraAllowListFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryJiraAllowListFilterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "QueryJiraAllowListFilter",
		Method:             "GET",
		PathPattern:        "/falcon-complete-dashboards/queries/jira-allowlist/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryJiraAllowListFilterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryJiraAllowListFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryJiraAllowListFilterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryJiraBlockListFilter retrieves jira block list tickets that match the provided filter criteria with scrolling enabled
*/
func (a *Client) QueryJiraBlockListFilter(params *QueryJiraBlockListFilterParams) (*QueryJiraBlockListFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryJiraBlockListFilterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "QueryJiraBlockListFilter",
		Method:             "GET",
		PathPattern:        "/falcon-complete-dashboards/queries/jira-blocklist/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryJiraBlockListFilterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryJiraBlockListFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryJiraBlockListFilterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryJiraEscalationsFilter retrieves jira escalation tickets that match the provided filter criteria with scrolling enabled
*/
func (a *Client) QueryJiraEscalationsFilter(params *QueryJiraEscalationsFilterParams) (*QueryJiraEscalationsFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryJiraEscalationsFilterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "QueryJiraEscalationsFilter",
		Method:             "GET",
		PathPattern:        "/falcon-complete-dashboards/queries/jira-escalations/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryJiraEscalationsFilterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryJiraEscalationsFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryJiraEscalationsFilterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryJiraRemediationsFilter retrieves jira remediation tickets that match the provided filter criteria with scrolling enabled
*/
func (a *Client) QueryJiraRemediationsFilter(params *QueryJiraRemediationsFilterParams) (*QueryJiraRemediationsFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryJiraRemediationsFilterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "QueryJiraRemediationsFilter",
		Method:             "GET",
		PathPattern:        "/falcon-complete-dashboards/queries/jira-remediations/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryJiraRemediationsFilterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryJiraRemediationsFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryJiraRemediationsFilterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}