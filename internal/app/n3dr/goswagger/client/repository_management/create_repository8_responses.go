// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository8Reader is a Reader for the CreateRepository8 structure.
type CreateRepository8Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository8Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository8Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository8Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository8Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository8Created creates a CreateRepository8Created with default headers values
func NewCreateRepository8Created() *CreateRepository8Created {
	return &CreateRepository8Created{}
}

/* CreateRepository8Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository8Created struct {
}

func (o *CreateRepository8Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/group][%d] createRepository8Created ", 201)
}

func (o *CreateRepository8Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository8Unauthorized creates a CreateRepository8Unauthorized with default headers values
func NewCreateRepository8Unauthorized() *CreateRepository8Unauthorized {
	return &CreateRepository8Unauthorized{}
}

/* CreateRepository8Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository8Unauthorized struct {
}

func (o *CreateRepository8Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/group][%d] createRepository8Unauthorized ", 401)
}

func (o *CreateRepository8Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository8Forbidden creates a CreateRepository8Forbidden with default headers values
func NewCreateRepository8Forbidden() *CreateRepository8Forbidden {
	return &CreateRepository8Forbidden{}
}

/* CreateRepository8Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository8Forbidden struct {
}

func (o *CreateRepository8Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/group][%d] createRepository8Forbidden ", 403)
}

func (o *CreateRepository8Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
