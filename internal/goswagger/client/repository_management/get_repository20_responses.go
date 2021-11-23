// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/goswagger/models"
)

// GetRepository20Reader is a Reader for the GetRepository20 structure.
type GetRepository20Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository20Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository20OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository20OK creates a GetRepository20OK with default headers values
func NewGetRepository20OK() *GetRepository20OK {
	return &GetRepository20OK{}
}

/* GetRepository20OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository20OK struct {
	Payload *models.SimpleAPIGroupRepository
}

func (o *GetRepository20OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/yum/group/{repositoryName}][%d] getRepository20OK  %+v", 200, o.Payload)
}
func (o *GetRepository20OK) GetPayload() *models.SimpleAPIGroupRepository {
	return o.Payload
}

func (o *GetRepository20OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIGroupRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
