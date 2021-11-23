// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository5Reader is a Reader for the CreateRepository5 structure.
type CreateRepository5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository5Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository5Created creates a CreateRepository5Created with default headers values
func NewCreateRepository5Created() *CreateRepository5Created {
	return &CreateRepository5Created{}
}

/* CreateRepository5Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository5Created struct {
}

func (o *CreateRepository5Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/group][%d] createRepository5Created ", 201)
}

func (o *CreateRepository5Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository5Unauthorized creates a CreateRepository5Unauthorized with default headers values
func NewCreateRepository5Unauthorized() *CreateRepository5Unauthorized {
	return &CreateRepository5Unauthorized{}
}

/* CreateRepository5Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository5Unauthorized struct {
}

func (o *CreateRepository5Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/group][%d] createRepository5Unauthorized ", 401)
}

func (o *CreateRepository5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository5Forbidden creates a CreateRepository5Forbidden with default headers values
func NewCreateRepository5Forbidden() *CreateRepository5Forbidden {
	return &CreateRepository5Forbidden{}
}

/* CreateRepository5Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository5Forbidden struct {
}

func (o *CreateRepository5Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/group][%d] createRepository5Forbidden ", 403)
}

func (o *CreateRepository5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
