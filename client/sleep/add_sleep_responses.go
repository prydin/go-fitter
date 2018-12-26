// Code generated by go-swagger; DO NOT EDIT.

package sleep

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddSleepReader is a Reader for the AddSleep structure.
type AddSleepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSleepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddSleepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewAddSleepUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddSleepForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddSleepOK creates a AddSleepOK with default headers values
func NewAddSleepOK() *AddSleepOK {
	return &AddSleepOK{}
}

/*AddSleepOK handles this case with default header values.

Successful request.
*/
type AddSleepOK struct {
}

func (o *AddSleepOK) Error() string {
	return fmt.Sprintf("[POST /1.2/user/-/sleep.json][%d] addSleepOK ", 200)
}

func (o *AddSleepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSleepUnauthorized creates a AddSleepUnauthorized with default headers values
func NewAddSleepUnauthorized() *AddSleepUnauthorized {
	return &AddSleepUnauthorized{}
}

/*AddSleepUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type AddSleepUnauthorized struct {
}

func (o *AddSleepUnauthorized) Error() string {
	return fmt.Sprintf("[POST /1.2/user/-/sleep.json][%d] addSleepUnauthorized ", 401)
}

func (o *AddSleepUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSleepForbidden creates a AddSleepForbidden with default headers values
func NewAddSleepForbidden() *AddSleepForbidden {
	return &AddSleepForbidden{}
}

/*AddSleepForbidden handles this case with default header values.

The server understood the request, but is refusing to fulfill it.
*/
type AddSleepForbidden struct {
}

func (o *AddSleepForbidden) Error() string {
	return fmt.Sprintf("[POST /1.2/user/-/sleep.json][%d] addSleepForbidden ", 403)
}

func (o *AddSleepForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
