// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetActivitiesTypesReader is a Reader for the GetActivitiesTypes structure.
type GetActivitiesTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActivitiesTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetActivitiesTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetActivitiesTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetActivitiesTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetActivitiesTypesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetActivitiesTypesOK creates a GetActivitiesTypesOK with default headers values
func NewGetActivitiesTypesOK() *GetActivitiesTypesOK {
	return &GetActivitiesTypesOK{}
}

/*GetActivitiesTypesOK handles this case with default header values.

A successful request.
*/
type GetActivitiesTypesOK struct {
}

func (o *GetActivitiesTypesOK) Error() string {
	return fmt.Sprintf("[GET /1/activities.json][%d] getActivitiesTypesOK ", 200)
}

func (o *GetActivitiesTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesTypesBadRequest creates a GetActivitiesTypesBadRequest with default headers values
func NewGetActivitiesTypesBadRequest() *GetActivitiesTypesBadRequest {
	return &GetActivitiesTypesBadRequest{}
}

/*GetActivitiesTypesBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetActivitiesTypesBadRequest struct {
}

func (o *GetActivitiesTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/activities.json][%d] getActivitiesTypesBadRequest ", 400)
}

func (o *GetActivitiesTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesTypesUnauthorized creates a GetActivitiesTypesUnauthorized with default headers values
func NewGetActivitiesTypesUnauthorized() *GetActivitiesTypesUnauthorized {
	return &GetActivitiesTypesUnauthorized{}
}

/*GetActivitiesTypesUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetActivitiesTypesUnauthorized struct {
}

func (o *GetActivitiesTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/activities.json][%d] getActivitiesTypesUnauthorized ", 401)
}

func (o *GetActivitiesTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActivitiesTypesConflict creates a GetActivitiesTypesConflict with default headers values
func NewGetActivitiesTypesConflict() *GetActivitiesTypesConflict {
	return &GetActivitiesTypesConflict{}
}

/*GetActivitiesTypesConflict handles this case with default header values.

Returned if the given user is already subscribed to this stream using a different subscription ID, OR if the given subscription ID is already used to identify a subscription to a different stream.
*/
type GetActivitiesTypesConflict struct {
}

func (o *GetActivitiesTypesConflict) Error() string {
	return fmt.Sprintf("[GET /1/activities.json][%d] getActivitiesTypesConflict ", 409)
}

func (o *GetActivitiesTypesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
