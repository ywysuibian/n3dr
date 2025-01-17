// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new blob store API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for blob store API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBlobStore(params *CreateBlobStoreParams, opts ...ClientOption) (*CreateBlobStoreCreated, error)

	CreateBlobStore1(params *CreateBlobStore1Params, opts ...ClientOption) (*CreateBlobStore1Created, error)

	CreateFileBlobStore(params *CreateFileBlobStoreParams, opts ...ClientOption) (*CreateFileBlobStoreNoContent, error)

	DeleteBlobStore(params *DeleteBlobStoreParams, opts ...ClientOption) error

	GetBlobStore(params *GetBlobStoreParams, opts ...ClientOption) (*GetBlobStoreOK, error)

	GetBlobStore1(params *GetBlobStore1Params, opts ...ClientOption) (*GetBlobStore1OK, error)

	GetFileBlobStoreConfiguration(params *GetFileBlobStoreConfigurationParams, opts ...ClientOption) (*GetFileBlobStoreConfigurationOK, error)

	ListBlobStores(params *ListBlobStoresParams, opts ...ClientOption) (*ListBlobStoresOK, error)

	QuotaStatus(params *QuotaStatusParams, opts ...ClientOption) (*QuotaStatusOK, error)

	UpdateBlobStore(params *UpdateBlobStoreParams, opts ...ClientOption) (*UpdateBlobStoreNoContent, error)

	UpdateBlobStore1(params *UpdateBlobStore1Params, opts ...ClientOption) (*UpdateBlobStore1NoContent, error)

	UpdateFileBlobStore(params *UpdateFileBlobStoreParams, opts ...ClientOption) (*UpdateFileBlobStoreNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBlobStore creates an s3 blob store
*/
func (a *Client) CreateBlobStore(params *CreateBlobStoreParams, opts ...ClientOption) (*CreateBlobStoreCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlobStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBlobStore",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/s3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBlobStoreReader{formats: a.formats},
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
	success, ok := result.(*CreateBlobStoreCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlobStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateBlobStore1 creates an azure blob store
*/
func (a *Client) CreateBlobStore1(params *CreateBlobStore1Params, opts ...ClientOption) (*CreateBlobStore1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlobStore1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBlobStore_1",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/azure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBlobStore1Reader{formats: a.formats},
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
	success, ok := result.(*CreateBlobStore1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlobStore_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateFileBlobStore creates a file blob store
*/
func (a *Client) CreateFileBlobStore(params *CreateFileBlobStoreParams, opts ...ClientOption) (*CreateFileBlobStoreNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFileBlobStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFileBlobStore",
		Method:             "POST",
		PathPattern:        "/v1/blobstores/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateFileBlobStoreReader{formats: a.formats},
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
	success, ok := result.(*CreateFileBlobStoreNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFileBlobStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBlobStore deletes a blob store by name
*/
func (a *Client) DeleteBlobStore(params *DeleteBlobStoreParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBlobStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteBlobStore",
		Method:             "DELETE",
		PathPattern:        "/v1/blobstores/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBlobStoreReader{formats: a.formats},
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
  GetBlobStore gets a s3 blob store configuration by name
*/
func (a *Client) GetBlobStore(params *GetBlobStoreParams, opts ...ClientOption) (*GetBlobStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBlobStore",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/s3/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobStoreReader{formats: a.formats},
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
	success, ok := result.(*GetBlobStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBlobStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBlobStore1 gets an azure blob store configuration by name
*/
func (a *Client) GetBlobStore1(params *GetBlobStore1Params, opts ...ClientOption) (*GetBlobStore1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlobStore1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBlobStore_1",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/azure/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlobStore1Reader{formats: a.formats},
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
	success, ok := result.(*GetBlobStore1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBlobStore_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFileBlobStoreConfiguration gets a file blob store configuration by name
*/
func (a *Client) GetFileBlobStoreConfiguration(params *GetFileBlobStoreConfigurationParams, opts ...ClientOption) (*GetFileBlobStoreConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileBlobStoreConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFileBlobStoreConfiguration",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/file/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFileBlobStoreConfigurationReader{formats: a.formats},
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
	success, ok := result.(*GetFileBlobStoreConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFileBlobStoreConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBlobStores lists the blob stores
*/
func (a *Client) ListBlobStores(params *ListBlobStoresParams, opts ...ClientOption) (*ListBlobStoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBlobStoresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listBlobStores",
		Method:             "GET",
		PathPattern:        "/v1/blobstores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBlobStoresReader{formats: a.formats},
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
	success, ok := result.(*ListBlobStoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBlobStores: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QuotaStatus gets quota status for a given blob store
*/
func (a *Client) QuotaStatus(params *QuotaStatusParams, opts ...ClientOption) (*QuotaStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuotaStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "quotaStatus",
		Method:             "GET",
		PathPattern:        "/v1/blobstores/{name}/quota-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuotaStatusReader{formats: a.formats},
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
	success, ok := result.(*QuotaStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for quotaStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBlobStore updates an s3 blob store configuration by name
*/
func (a *Client) UpdateBlobStore(params *UpdateBlobStoreParams, opts ...ClientOption) (*UpdateBlobStoreNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBlobStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBlobStore",
		Method:             "PUT",
		PathPattern:        "/v1/blobstores/s3/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateBlobStoreReader{formats: a.formats},
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
	success, ok := result.(*UpdateBlobStoreNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBlobStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBlobStore1 updates an azure blob store configuration by name
*/
func (a *Client) UpdateBlobStore1(params *UpdateBlobStore1Params, opts ...ClientOption) (*UpdateBlobStore1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBlobStore1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBlobStore_1",
		Method:             "PUT",
		PathPattern:        "/v1/blobstores/azure/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateBlobStore1Reader{formats: a.formats},
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
	success, ok := result.(*UpdateBlobStore1NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBlobStore_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateFileBlobStore updates a file blob store configuration by name
*/
func (a *Client) UpdateFileBlobStore(params *UpdateFileBlobStoreParams, opts ...ClientOption) (*UpdateFileBlobStoreNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFileBlobStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateFileBlobStore",
		Method:             "PUT",
		PathPattern:        "/v1/blobstores/file/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateFileBlobStoreReader{formats: a.formats},
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
	success, ok := result.(*UpdateFileBlobStoreNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFileBlobStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
