// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/goswagger/models"
)

// GetComponentByIDReader is a Reader for the GetComponentByID structure.
type GetComponentByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComponentByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComponentByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetComponentByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetComponentByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetComponentByIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComponentByIDOK creates a GetComponentByIDOK with default headers values
func NewGetComponentByIDOK() *GetComponentByIDOK {
	return &GetComponentByIDOK{}
}

/* GetComponentByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetComponentByIDOK struct {
	Payload *models.ComponentXO
}

func (o *GetComponentByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/components/{id}][%d] getComponentByIdOK  %+v", 200, o.Payload)
}
func (o *GetComponentByIDOK) GetPayload() *models.ComponentXO {
	return o.Payload
}

func (o *GetComponentByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentByIDForbidden creates a GetComponentByIDForbidden with default headers values
func NewGetComponentByIDForbidden() *GetComponentByIDForbidden {
	return &GetComponentByIDForbidden{}
}

/* GetComponentByIDForbidden describes a response with status code 403, with default header values.

Insufficient permissions to get component
*/
type GetComponentByIDForbidden struct {
}

func (o *GetComponentByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/components/{id}][%d] getComponentByIdForbidden ", 403)
}

func (o *GetComponentByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComponentByIDNotFound creates a GetComponentByIDNotFound with default headers values
func NewGetComponentByIDNotFound() *GetComponentByIDNotFound {
	return &GetComponentByIDNotFound{}
}

/* GetComponentByIDNotFound describes a response with status code 404, with default header values.

Component not found
*/
type GetComponentByIDNotFound struct {
}

func (o *GetComponentByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/components/{id}][%d] getComponentByIdNotFound ", 404)
}

func (o *GetComponentByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComponentByIDUnprocessableEntity creates a GetComponentByIDUnprocessableEntity with default headers values
func NewGetComponentByIDUnprocessableEntity() *GetComponentByIDUnprocessableEntity {
	return &GetComponentByIDUnprocessableEntity{}
}

/* GetComponentByIDUnprocessableEntity describes a response with status code 422, with default header values.

Malformed ID
*/
type GetComponentByIDUnprocessableEntity struct {
}

func (o *GetComponentByIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /v1/components/{id}][%d] getComponentByIdUnprocessableEntity ", 422)
}

func (o *GetComponentByIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
