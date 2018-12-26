// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteFoodsLogReader is a Reader for the DeleteFoodsLog structure.
type DeleteFoodsLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFoodsLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteFoodsLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFoodsLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteFoodsLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteFoodsLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFoodsLogOK creates a DeleteFoodsLogOK with default headers values
func NewDeleteFoodsLogOK() *DeleteFoodsLogOK {
	return &DeleteFoodsLogOK{}
}

/*DeleteFoodsLogOK handles this case with default header values.

Successful request.
*/
type DeleteFoodsLogOK struct {
}

func (o *DeleteFoodsLogOK) Error() string {
	return fmt.Sprintf("[DELETE /1/user/-/foods/log/{food-log-id}.json][%d] deleteFoodsLogOK ", 200)
}

func (o *DeleteFoodsLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFoodsLogBadRequest creates a DeleteFoodsLogBadRequest with default headers values
func NewDeleteFoodsLogBadRequest() *DeleteFoodsLogBadRequest {
	return &DeleteFoodsLogBadRequest{}
}

/*DeleteFoodsLogBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type DeleteFoodsLogBadRequest struct {
}

func (o *DeleteFoodsLogBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1/user/-/foods/log/{food-log-id}.json][%d] deleteFoodsLogBadRequest ", 400)
}

func (o *DeleteFoodsLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFoodsLogUnauthorized creates a DeleteFoodsLogUnauthorized with default headers values
func NewDeleteFoodsLogUnauthorized() *DeleteFoodsLogUnauthorized {
	return &DeleteFoodsLogUnauthorized{}
}

/*DeleteFoodsLogUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type DeleteFoodsLogUnauthorized struct {
}

func (o *DeleteFoodsLogUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /1/user/-/foods/log/{food-log-id}.json][%d] deleteFoodsLogUnauthorized ", 401)
}

func (o *DeleteFoodsLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFoodsLogForbidden creates a DeleteFoodsLogForbidden with default headers values
func NewDeleteFoodsLogForbidden() *DeleteFoodsLogForbidden {
	return &DeleteFoodsLogForbidden{}
}

/*DeleteFoodsLogForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type DeleteFoodsLogForbidden struct {
}

func (o *DeleteFoodsLogForbidden) Error() string {
	return fmt.Sprintf("[DELETE /1/user/-/foods/log/{food-log-id}.json][%d] deleteFoodsLogForbidden ", 403)
}

func (o *DeleteFoodsLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}