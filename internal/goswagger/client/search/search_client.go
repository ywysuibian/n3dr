// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Search(params *SearchParams, opts ...ClientOption) (*SearchOK, error)

	SearchAndDownloadAssets(params *SearchAndDownloadAssetsParams, opts ...ClientOption) error

	SearchAssets(params *SearchAssetsParams, opts ...ClientOption) (*SearchAssetsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Search searches components
*/
func (a *Client) Search(params *SearchParams, opts ...ClientOption) (*SearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search",
		Method:             "GET",
		PathPattern:        "/v1/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchReader{formats: a.formats},
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
	success, ok := result.(*SearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchAndDownloadAssets searches and download asset

  Returns a 302 Found with location header field set to download URL. Unless a sort parameter is supplied, the search must return a single asset to receive download URL.
*/
func (a *Client) SearchAndDownloadAssets(params *SearchAndDownloadAssetsParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchAndDownloadAssetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchAndDownloadAssets",
		Method:             "GET",
		PathPattern:        "/v1/search/assets/download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchAndDownloadAssetsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  SearchAssets searches assets
*/
func (a *Client) SearchAssets(params *SearchAssetsParams, opts ...ClientOption) (*SearchAssetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchAssetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchAssets",
		Method:             "GET",
		PathPattern:        "/v1/search/assets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchAssetsReader{formats: a.formats},
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
	success, ok := result.(*SearchAssetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchAssets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}