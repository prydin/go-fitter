// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateWaterLogReader is a Reader for the UpdateWaterLog structure.
type UpdateWaterLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWaterLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateWaterLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateWaterLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateWaterLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateWaterLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateWaterLogOK creates a UpdateWaterLogOK with default headers values
func NewUpdateWaterLogOK() *UpdateWaterLogOK {
	return &UpdateWaterLogOK{}
}

/*UpdateWaterLogOK handles this case with default header values.

Successful request.
*/
type UpdateWaterLogOK struct {
}

func (o *UpdateWaterLogOK) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/water/{water-log-id}.json][%d] updateWaterLogOK ", 200)
}

func (o *UpdateWaterLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateWaterLogBadRequest creates a UpdateWaterLogBadRequest with default headers values
func NewUpdateWaterLogBadRequest() *UpdateWaterLogBadRequest {
	return &UpdateWaterLogBadRequest{}
}

/*UpdateWaterLogBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type UpdateWaterLogBadRequest struct {
}

func (o *UpdateWaterLogBadRequest) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/water/{water-log-id}.json][%d] updateWaterLogBadRequest ", 400)
}

func (o *UpdateWaterLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateWaterLogUnauthorized creates a UpdateWaterLogUnauthorized with default headers values
func NewUpdateWaterLogUnauthorized() *UpdateWaterLogUnauthorized {
	return &UpdateWaterLogUnauthorized{}
}

/*UpdateWaterLogUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type UpdateWaterLogUnauthorized struct {
}

func (o *UpdateWaterLogUnauthorized) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/water/{water-log-id}.json][%d] updateWaterLogUnauthorized ", 401)
}

func (o *UpdateWaterLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateWaterLogForbidden creates a UpdateWaterLogForbidden with default headers values
func NewUpdateWaterLogForbidden() *UpdateWaterLogForbidden {
	return &UpdateWaterLogForbidden{}
}

/*UpdateWaterLogForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type UpdateWaterLogForbidden struct {
}

func (o *UpdateWaterLogForbidden) Error() string {
	return fmt.Sprintf("[POST /1/user/-/foods/log/water/{water-log-id}.json][%d] updateWaterLogForbidden ", 403)
}

func (o *UpdateWaterLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
