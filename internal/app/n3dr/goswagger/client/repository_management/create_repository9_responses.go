// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository9Reader is a Reader for the CreateRepository9 structure.
type CreateRepository9Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository9Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository9Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository9Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository9Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository9Created creates a CreateRepository9Created with default headers values
func NewCreateRepository9Created() *CreateRepository9Created {
	return &CreateRepository9Created{}
}

/* CreateRepository9Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository9Created struct {
}

func (o *CreateRepository9Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] createRepository9Created ", 201)
}

func (o *CreateRepository9Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository9Unauthorized creates a CreateRepository9Unauthorized with default headers values
func NewCreateRepository9Unauthorized() *CreateRepository9Unauthorized {
	return &CreateRepository9Unauthorized{}
}

/* CreateRepository9Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository9Unauthorized struct {
}

func (o *CreateRepository9Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] createRepository9Unauthorized ", 401)
}

func (o *CreateRepository9Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository9Forbidden creates a CreateRepository9Forbidden with default headers values
func NewCreateRepository9Forbidden() *CreateRepository9Forbidden {
	return &CreateRepository9Forbidden{}
}

/* CreateRepository9Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository9Forbidden struct {
}

func (o *CreateRepository9Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/hosted][%d] createRepository9Forbidden ", 403)
}

func (o *CreateRepository9Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
