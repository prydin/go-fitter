// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetActivitiesTCXReader is a Reader for the GetActivitiesTCX structure.
type GetActivitiesTCXReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActivitiesTCXReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetActivitiesTCXOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetActivitiesTCXBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetActivitiesTCXUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetActivitiesTCXForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetActivitiesTCXConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetActivitiesTCXOK creates a GetActivitiesTCXOK with default headers values
func NewGetActivitiesTCXOK() *GetActivitiesTCXOK {
	return &GetActivitiesTCXOK{}
}

/*GetActivitiesTCXOK handles this case with default header values.

A successful request.
*/
type GetActivitiesTCXOK struct {
}

func (o *GetActivitiesTCXOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{log-id}.tcx][%d] getActivitiesTCXOK ", 200)
}

func (o *GetActivitiesTCXOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesTCXBadRequest creates a GetActivitiesTCXBadRequest with default headers values
func NewGetActivitiesTCXBadRequest() *GetActivitiesTCXBadRequest {
	return &GetActivitiesTCXBadRequest{}
}

/*GetActivitiesTCXBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetActivitiesTCXBadRequest struct {
}

func (o *GetActivitiesTCXBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{log-id}.tcx][%d] getActivitiesTCXBadRequest ", 400)
}

func (o *GetActivitiesTCXBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesTCXUnauthorized creates a GetActivitiesTCXUnauthorized with default headers values
func NewGetActivitiesTCXUnauthorized() *GetActivitiesTCXUnauthorized {
	return &GetActivitiesTCXUnauthorized{}
}

/*GetActivitiesTCXUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetActivitiesTCXUnauthorized struct {
}

func (o *GetActivitiesTCXUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{log-id}.tcx][%d] getActivitiesTCXUnauthorized ", 401)
}

func (o *GetActivitiesTCXUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesTCXForbidden creates a GetActivitiesTCXForbidden with default headers values
func NewGetActivitiesTCXForbidden() *GetActivitiesTCXForbidden {
	return &GetActivitiesTCXForbidden{}
}

/*GetActivitiesTCXForbidden handles this case with default header values.

The request was a valid request, but the server is refusing to respond to it.
*/
type GetActivitiesTCXForbidden struct {
}

func (o *GetActivitiesTCXForbidden) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{log-id}.tcx][%d] getActivitiesTCXForbidden ", 403)
}

func (o *GetActivitiesTCXForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesTCXConflict creates a GetActivitiesTCXConflict with default headers values
func NewGetActivitiesTCXConflict() *GetActivitiesTCXConflict {
	return &GetActivitiesTCXConflict{}
}

/*GetActivitiesTCXConflict handles this case with default header values.

Returned if the given user is already subscribed to this stream using a different subscription ID, OR if the given subscription ID is already used to identify a subscription to a different stream.
*/
type GetActivitiesTCXConflict struct {
}

func (o *GetActivitiesTCXConflict) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/{log-id}.tcx][%d] getActivitiesTCXConflict ", 409)
}

func (o *GetActivitiesTCXConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}