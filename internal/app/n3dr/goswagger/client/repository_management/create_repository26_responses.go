// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository26Reader is a Reader for the CreateRepository26 structure.
type CreateRepository26Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository26Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository26Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository26Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository26Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository26Created creates a CreateRepository26Created with default headers values
func NewCreateRepository26Created() *CreateRepository26Created {
	return &CreateRepository26Created{}
}

/* CreateRepository26Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository26Created struct {
}

func (o *CreateRepository26Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] createRepository26Created ", 201)
}

func (o *CreateRepository26Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository26Unauthorized creates a CreateRepository26Unauthorized with default headers values
func NewCreateRepository26Unauthorized() *CreateRepository26Unauthorized {
	return &CreateRepository26Unauthorized{}
}

/* CreateRepository26Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository26Unauthorized struct {
}

func (o *CreateRepository26Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] createRepository26Unauthorized ", 401)
}

func (o *CreateRepository26Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository26Forbidden creates a CreateRepository26Forbidden with default headers values
func NewCreateRepository26Forbidden() *CreateRepository26Forbidden {
	return &CreateRepository26Forbidden{}
}

/* CreateRepository26Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository26Forbidden struct {
}

func (o *CreateRepository26Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/pypi/group][%d] createRepository26Forbidden ", 403)
}

func (o *CreateRepository26Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
