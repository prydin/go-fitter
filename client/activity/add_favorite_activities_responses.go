// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddFavoriteActivitiesReader is a Reader for the AddFavoriteActivities structure.
type AddFavoriteActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddFavoriteActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddFavoriteActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewAddFavoriteActivitiesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddFavoriteActivitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddFavoriteActivitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAddFavoriteActivitiesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddFavoriteActivitiesOK creates a AddFavoriteActivitiesOK with default headers values
func NewAddFavoriteActivitiesOK() *AddFavoriteActivitiesOK {
	return &AddFavoriteActivitiesOK{}
}

/*AddFavoriteActivitiesOK handles this case with default header values.

A successful request.
*/
type AddFavoriteActivitiesOK struct {
}

func (o *AddFavoriteActivitiesOK) Error() string {
	return fmt.Sprintf("[POST /1/user/-/activities/favorite/{activity-id}.json][%d] addFavoriteActivitiesOK ", 200)
}

func (o *AddFavoriteActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddFavoriteActivitiesCreated creates a AddFavoriteActivitiesCreated with default headers values
func NewAddFavoriteActivitiesCreated() *AddFavoriteActivitiesCreated {
	return &AddFavoriteActivitiesCreated{}
}

/*AddFavoriteActivitiesCreated handles this case with default header values.

The request fulfilled and new resource being created.
*/
type AddFavoriteActivitiesCreated struct {
}

func (o *AddFavoriteActivitiesCreated) Error() string {
	return fmt.Sprintf("[POST /1/user/-/activities/favorite/{activity-id}.json][%d] addFavoriteActivitiesCreated ", 201)
}

func (o *AddFavoriteActivitiesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddFavoriteActivitiesBadRequest creates a AddFavoriteActivitiesBadRequest with default headers values
func NewAddFavoriteActivitiesBadRequest() *AddFavoriteActivitiesBadRequest {
	return &AddFavoriteActivitiesBadRequest{}
}

/*AddFavoriteActivitiesBadRequest handles this case with default header values.

The request had bad syntax or was inherently impossible to be satified.
*/
type AddFavoriteActivitiesBadRequest struct {
}

func (o *AddFavoriteActivitiesBadRequest) Error() string {
	return fmt.Sprintf("[POST /1/user/-/activities/favorite/{activity-id}.json][%d] addFavoriteActivitiesBadRequest ", 400)
}

func (o *AddFavoriteActivitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddFavoriteActivitiesUnauthorized creates a AddFavoriteActivitiesUnauthorized with default headers values
func NewAddFavoriteActivitiesUnauthorized() *AddFavoriteActivitiesUnauthorized {
	return &AddFavoriteActivitiesUnauthorized{}
}

/*AddFavoriteActivitiesUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type AddFavoriteActivitiesUnauthorized struct {
}

func (o *AddFavoriteActivitiesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /1/user/-/activities/favorite/{activity-id}.json][%d] addFavoriteActivitiesUnauthorized ", 401)
}

func (o *AddFavoriteActivitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddFavoriteActivitiesConflict creates a AddFavoriteActivitiesConflict with default headers values
func NewAddFavoriteActivitiesConflict() *AddFavoriteActivitiesConflict {
	return &AddFavoriteActivitiesConflict{}
}

/*AddFavoriteActivitiesConflict handles this case with default header values.

Returned if the given user is already subscribed to this stream using a different subscription ID, OR if the given subscription ID is already used to identify a subscription to a different stream.
*/
type AddFavoriteActivitiesConflict struct {
}

func (o *AddFavoriteActivitiesConflict) Error() string {
	return fmt.Sprintf("[POST /1/user/-/activities/favorite/{activity-id}.json][%d] addFavoriteActivitiesConflict ", 409)
}

func (o *AddFavoriteActivitiesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
