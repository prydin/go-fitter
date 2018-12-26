// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteWaterLogReader is a Reader for the DeleteWaterLog structure.
type DeleteWaterLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWaterLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteWaterLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteWaterLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteWaterLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteWaterLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteWaterLogOK creates a DeleteWaterLogOK with default headers values
func NewDeleteWaterLogOK() *DeleteWaterLogOK {
	return &DeleteWaterLogOK{}
}

/*DeleteWaterLogOK handles this case with default header values.

Successful request.
*/
type DeleteWaterLogOK struct {
}

func (o *DeleteWaterLogOK) Error() string {
	return fmt.Sprintf("[DELETE /1/user/-/foods/log/water/{water-log-id}.json][%d] deleteWaterLogOK ", 200)
}

func (o *DeleteWaterLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWaterLogBadRequest creates a DeleteWaterLogBadRequest with default headers values
func NewDeleteWaterLogBadRequest() *DeleteWaterLogBadRequest {
	return &DeleteWaterLogBadRequest{}
}

/*DeleteWaterLogBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type DeleteWaterLogBadRequest struct {
}

func (o *DeleteWaterLogBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1/user/-/foods/log/water/{water-log-id}.json][%d] deleteWaterLogBadRequest ", 400)
}

func (o *DeleteWaterLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWaterLogUnauthorized creates a DeleteWaterLogUnauthorized with default headers values
func NewDeleteWaterLogUnauthorized() *DeleteWaterLogUnauthorized {
	return &DeleteWaterLogUnauthorized{}
}

/*DeleteWaterLogUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type DeleteWaterLogUnauthorized struct {
}

func (o *DeleteWaterLogUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /1/user/-/foods/log/water/{water-log-id}.json][%d] deleteWaterLogUnauthorized ", 401)
}

func (o *DeleteWaterLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWaterLogForbidden creates a DeleteWaterLogForbidden with default headers values
func NewDeleteWaterLogForbidden() *DeleteWaterLogForbidden {
	return &DeleteWaterLogForbidden{}
}

/*DeleteWaterLogForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type DeleteWaterLogForbidden struct {
}

func (o *DeleteWaterLogForbidden) Error() string {
	return fmt.Sprintf("[DELETE /1/user/-/foods/log/water/{water-log-id}.json][%d] deleteWaterLogForbidden ", 403)
}

func (o *DeleteWaterLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}