// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/garcialuis/Gossip/client/models"
)

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*GetUserOK handles this case with default header values.

User record
*/
type GetUserOK struct {
	Payload *models.User
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserBadRequest creates a GetUserBadRequest with default headers values
func NewGetUserBadRequest() *GetUserBadRequest {
	return &GetUserBadRequest{}
}

/*GetUserBadRequest handles this case with default header values.

Bad Request
*/
type GetUserBadRequest struct {
}

func (o *GetUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserBadRequest ", 400)
}

func (o *GetUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}