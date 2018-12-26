// Code generated by go-swagger; DO NOT EDIT.

package body_and_weight

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddWeightLogReader is a Reader for the AddWeightLog structure.
type AddWeightLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWeightLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddWeightLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewAddWeightLogCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddWeightLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddWeightLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddWeightLogOK creates a AddWeightLogOK with default headers values
func NewAddWeightLogOK() *AddWeightLogOK {
	return &AddWeightLogOK{}
}

/*AddWeightLogOK handles this case with default header values.

A successful request.
*/
type AddWeightLogOK struct {
}

func (o *AddWeightLogOK) Error() string {
	return fmt.Sprintf("[POST /1/user/-/body/log/weight.json][%d] addWeightLogOK ", 200)
}

func (o *AddWeightLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddWeightLogCreated creates a AddWeightLogCreated with default headers values
func NewAddWeightLogCreated() *AddWeightLogCreated {
	return &AddWeightLogCreated{}
}

/*AddWeightLogCreated handles this case with default header values.

The request has been fulfilled and resulted in a new resouce being created.
*/
type AddWeightLogCreated struct {
}

func (o *AddWeightLogCreated) Error() string {
	return fmt.Sprintf("[POST /1/user/-/body/log/weight.json][%d] addWeightLogCreated ", 201)
}

func (o *AddWeightLogCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddWeightLogBadRequest creates a AddWeightLogBadRequest with default headers values
func NewAddWeightLogBadRequest() *AddWeightLogBadRequest {
	return &AddWeightLogBadRequest{}
}

/*AddWeightLogBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type AddWeightLogBadRequest struct {
}

func (o *AddWeightLogBadRequest) Error() string {
	return fmt.Sprintf("[POST /1/user/-/body/log/weight.json][%d] addWeightLogBadRequest ", 400)
}

func (o *AddWeightLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddWeightLogUnauthorized creates a AddWeightLogUnauthorized with default headers values
func NewAddWeightLogUnauthorized() *AddWeightLogUnauthorized {
	return &AddWeightLogUnauthorized{}
}

/*AddWeightLogUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type AddWeightLogUnauthorized struct {
}

func (o *AddWeightLogUnauthorized) Error() string {
	return fmt.Sprintf("[POST /1/user/-/body/log/weight.json][%d] addWeightLogUnauthorized ", 401)
}

func (o *AddWeightLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}