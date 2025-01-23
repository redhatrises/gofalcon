// Code generated by go-swagger; DO NOT EDIT.

package cloud_aws_registration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cloud aws registration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cloud aws registration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CloudRegistrationAwsCreateAccount(params *CloudRegistrationAwsCreateAccountParams, opts ...ClientOption) (*CloudRegistrationAwsCreateAccountCreated, *CloudRegistrationAwsCreateAccountMultiStatus, error)

	CloudRegistrationAwsDeleteAccount(params *CloudRegistrationAwsDeleteAccountParams, opts ...ClientOption) (*CloudRegistrationAwsDeleteAccountOK, *CloudRegistrationAwsDeleteAccountMultiStatus, error)

	CloudRegistrationAwsGetAccounts(params *CloudRegistrationAwsGetAccountsParams, opts ...ClientOption) (*CloudRegistrationAwsGetAccountsOK, *CloudRegistrationAwsGetAccountsMultiStatus, error)

	CloudRegistrationAwsQueryAccounts(params *CloudRegistrationAwsQueryAccountsParams, opts ...ClientOption) (*CloudRegistrationAwsQueryAccountsOK, *CloudRegistrationAwsQueryAccountsMultiStatus, error)

	CloudRegistrationAwsUpdateAccount(params *CloudRegistrationAwsUpdateAccountParams, opts ...ClientOption) (*CloudRegistrationAwsUpdateAccountOK, *CloudRegistrationAwsUpdateAccountMultiStatus, error)

	CloudRegistrationAwsValidateAccounts(params *CloudRegistrationAwsValidateAccountsParams, opts ...ClientOption) (*CloudRegistrationAwsValidateAccountsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CloudRegistrationAwsCreateAccount creates a new account in our system for a customer
*/
func (a *Client) CloudRegistrationAwsCreateAccount(params *CloudRegistrationAwsCreateAccountParams, opts ...ClientOption) (*CloudRegistrationAwsCreateAccountCreated, *CloudRegistrationAwsCreateAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudRegistrationAwsCreateAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud-registration-aws-create-account",
		Method:             "POST",
		PathPattern:        "/cloud-security-registration-aws/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudRegistrationAwsCreateAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CloudRegistrationAwsCreateAccountCreated:
		return value, nil, nil
	case *CloudRegistrationAwsCreateAccountMultiStatus:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cloud_aws_registration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CloudRegistrationAwsDeleteAccount deletes an existing a w s account or organization in our system
*/
func (a *Client) CloudRegistrationAwsDeleteAccount(params *CloudRegistrationAwsDeleteAccountParams, opts ...ClientOption) (*CloudRegistrationAwsDeleteAccountOK, *CloudRegistrationAwsDeleteAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudRegistrationAwsDeleteAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud-registration-aws-delete-account",
		Method:             "DELETE",
		PathPattern:        "/cloud-security-registration-aws/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudRegistrationAwsDeleteAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CloudRegistrationAwsDeleteAccountOK:
		return value, nil, nil
	case *CloudRegistrationAwsDeleteAccountMultiStatus:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cloud_aws_registration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CloudRegistrationAwsGetAccounts retrieves existing a w s accounts by account i ds
*/
func (a *Client) CloudRegistrationAwsGetAccounts(params *CloudRegistrationAwsGetAccountsParams, opts ...ClientOption) (*CloudRegistrationAwsGetAccountsOK, *CloudRegistrationAwsGetAccountsMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudRegistrationAwsGetAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud-registration-aws-get-accounts",
		Method:             "GET",
		PathPattern:        "/cloud-security-registration-aws/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudRegistrationAwsGetAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CloudRegistrationAwsGetAccountsOK:
		return value, nil, nil
	case *CloudRegistrationAwsGetAccountsMultiStatus:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cloud_aws_registration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CloudRegistrationAwsQueryAccounts retrieves existing a w s accounts by account i ds
*/
func (a *Client) CloudRegistrationAwsQueryAccounts(params *CloudRegistrationAwsQueryAccountsParams, opts ...ClientOption) (*CloudRegistrationAwsQueryAccountsOK, *CloudRegistrationAwsQueryAccountsMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudRegistrationAwsQueryAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud-registration-aws-query-accounts",
		Method:             "GET",
		PathPattern:        "/cloud-security-registration-aws/queries/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudRegistrationAwsQueryAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CloudRegistrationAwsQueryAccountsOK:
		return value, nil, nil
	case *CloudRegistrationAwsQueryAccountsMultiStatus:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cloud_aws_registration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CloudRegistrationAwsUpdateAccount patches a existing account in our system for a customer
*/
func (a *Client) CloudRegistrationAwsUpdateAccount(params *CloudRegistrationAwsUpdateAccountParams, opts ...ClientOption) (*CloudRegistrationAwsUpdateAccountOK, *CloudRegistrationAwsUpdateAccountMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudRegistrationAwsUpdateAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud-registration-aws-update-account",
		Method:             "PATCH",
		PathPattern:        "/cloud-security-registration-aws/entities/account/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudRegistrationAwsUpdateAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CloudRegistrationAwsUpdateAccountOK:
		return value, nil, nil
	case *CloudRegistrationAwsUpdateAccountMultiStatus:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cloud_aws_registration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CloudRegistrationAwsValidateAccounts validates the a w s account in our system for a provided c ID for internal clients only
*/
func (a *Client) CloudRegistrationAwsValidateAccounts(params *CloudRegistrationAwsValidateAccountsParams, opts ...ClientOption) (*CloudRegistrationAwsValidateAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudRegistrationAwsValidateAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud-registration-aws-validate-accounts",
		Method:             "POST",
		PathPattern:        "/cloud-security-registration-aws/entities/account/validate/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudRegistrationAwsValidateAccountsReader{formats: a.formats},
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
	success, ok := result.(*CloudRegistrationAwsValidateAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cloud-registration-aws-validate-accounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}