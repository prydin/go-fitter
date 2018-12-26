// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetFrequentActivitiesReader is a Reader for the GetFrequentActivities structure.
type GetFrequentActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFrequentActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFrequentActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFrequentActivitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFrequentActivitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetFrequentActivitiesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFrequentActivitiesOK creates a GetFrequentActivitiesOK with default headers values
func NewGetFrequentActivitiesOK() *GetFrequentActivitiesOK {
	return &GetFrequentActivitiesOK{}
}

/*GetFrequentActivitiesOK handles this case with default header values.

A successful request.
*/
type GetFrequentActivitiesOK struct {
}

func (o *GetFrequentActivitiesOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/frequent.json][%d] getFrequentActivitiesOK ", 200)
}

func (o *GetFrequentActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFrequentActivitiesBadRequest creates a GetFrequentActivitiesBadRequest with default headers values
func NewGetFrequentActivitiesBadRequest() *GetFrequentActivitiesBadRequest {
	return &GetFrequentActivitiesBadRequest{}
}

/*GetFrequentActivitiesBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type GetFrequentActivitiesBadRequest struct {
}

func (o *GetFrequentActivitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/frequent.json][%d] getFrequentActivitiesBadRequest ", 400)
}

func (o *GetFrequentActivitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFrequentActivitiesUnauthorized creates a GetFrequentActivitiesUnauthorized with default headers values
func NewGetFrequentActivitiesUnauthorized() *GetFrequentActivitiesUnauthorized {
	return &GetFrequentActivitiesUnauthorized{}
}

/*GetFrequentActivitiesUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetFrequentActivitiesUnauthorized struct {
}

func (o *GetFrequentActivitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/frequent.json][%d] getFrequentActivitiesUnauthorized ", 401)
}

func (o *GetFrequentActivitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFrequentActivitiesConflict creates a GetFrequentActivitiesConflict with default headers values
func NewGetFrequentActivitiesConflict() *GetFrequentActivitiesConflict {
	return &GetFrequentActivitiesConflict{}
}

/*GetFrequentActivitiesConflict handles this case with default header values.

Returned if the given user is already subscribed to this stream using a different subscription ID, OR if the given subscription ID is already used to identify a subscription to a different stream.
*/
type GetFrequentActivitiesConflict struct {
}

func (o *GetFrequentActivitiesConflict) Error() string {
	return fmt.Sprintf("[GET /1/user/-/activities/frequent.json][%d] getFrequentActivitiesConflict ", 409)
}

func (o *GetFrequentActivitiesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}