// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*UpdateUserOK handles this case with default header values.

userResponse
*/
type UpdateUserOK struct {
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /users][%d] updateUserOK ", 200)
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserUnauthorized creates a UpdateUserUnauthorized with default headers values
func NewUpdateUserUnauthorized() *UpdateUserUnauthorized {
	return &UpdateUserUnauthorized{}
}

/*UpdateUserUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateUserUnauthorized struct {
}

func (o *UpdateUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users][%d] updateUserUnauthorized ", 401)
}

func (o *UpdateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserUnprocessableEntity creates a UpdateUserUnprocessableEntity with default headers values
func NewUpdateUserUnprocessableEntity() *UpdateUserUnprocessableEntity {
	return &UpdateUserUnprocessableEntity{}
}

/*UpdateUserUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type UpdateUserUnprocessableEntity struct {
}

func (o *UpdateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /users][%d] updateUserUnprocessableEntity ", 422)
}

func (o *UpdateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserInternalServerError creates a UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {
	return &UpdateUserInternalServerError{}
}

/*UpdateUserInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateUserInternalServerError struct {
}

func (o *UpdateUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users][%d] updateUserInternalServerError ", 500)
}

func (o *UpdateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}