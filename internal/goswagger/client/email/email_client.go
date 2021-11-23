// Code generated by go-swagger; DO NOT EDIT.

package email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new email API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for email API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteEmailConfiguration(params *DeleteEmailConfigurationParams, opts ...ClientOption) (*DeleteEmailConfigurationNoContent, error)

	GetEmailConfiguration(params *GetEmailConfigurationParams, opts ...ClientOption) (*GetEmailConfigurationOK, error)

	SetEmailConfiguration(params *SetEmailConfigurationParams, opts ...ClientOption) (*SetEmailConfigurationNoContent, error)

	TestEmailConfiguration(params *TestEmailConfigurationParams, opts ...ClientOption) (*TestEmailConfigurationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteEmailConfiguration disables and clear the email configuration
*/
func (a *Client) DeleteEmailConfiguration(params *DeleteEmailConfigurationParams, opts ...ClientOption) (*DeleteEmailConfigurationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEmailConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteEmailConfiguration",
		Method:             "DELETE",
		PathPattern:        "/v1/email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEmailConfigurationReader{formats: a.formats},
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
	success, ok := result.(*DeleteEmailConfigurationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteEmailConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEmailConfiguration retrieves the current email configuration
*/
func (a *Client) GetEmailConfiguration(params *GetEmailConfigurationParams, opts ...ClientOption) (*GetEmailConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmailConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEmailConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmailConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetEmailConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEmailConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetEmailConfiguration sets the current email configuration
*/
func (a *Client) SetEmailConfiguration(params *SetEmailConfigurationParams, opts ...ClientOption) (*SetEmailConfigurationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetEmailConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setEmailConfiguration",
		Method:             "PUT",
		PathPattern:        "/v1/email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetEmailConfigurationReader{formats: a.formats},
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
	success, ok := result.(*SetEmailConfigurationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setEmailConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TestEmailConfiguration sends a test email to the email address provided in the request body
*/
func (a *Client) TestEmailConfiguration(params *TestEmailConfigurationParams, opts ...ClientOption) (*TestEmailConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestEmailConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "testEmailConfiguration",
		Method:             "POST",
		PathPattern:        "/v1/email/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TestEmailConfigurationReader{formats: a.formats},
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
	success, ok := result.(*TestEmailConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for testEmailConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
