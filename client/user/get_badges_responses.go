// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBadgesReader is a Reader for the GetBadges structure.
type GetBadgesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBadgesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBadgesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewGetBadgesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewGetBadgesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBadgesOK creates a GetBadgesOK with default headers values
func NewGetBadgesOK() *GetBadgesOK {
	return &GetBadgesOK{}
}

/*GetBadgesOK handles this case with default header values.

A successful request.
*/
type GetBadgesOK struct {
}

func (o *GetBadgesOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/badges.json][%d] getBadgesOK ", 200)
}

func (o *GetBadgesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBadgesCreated creates a GetBadgesCreated with default headers values
func NewGetBadgesCreated() *GetBadgesCreated {
	return &GetBadgesCreated{}
}

/*GetBadgesCreated handles this case with default header values.

Returned if a new subscription was created in response to your request.
*/
type GetBadgesCreated struct {
}

func (o *GetBadgesCreated) Error() string {
	return fmt.Sprintf("[GET /1/user/-/badges.json][%d] getBadgesCreated ", 201)
}

func (o *GetBadgesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBadgesConflict creates a GetBadgesConflict with default headers values
func NewGetBadgesConflict() *GetBadgesConflict {
	return &GetBadgesConflict{}
}

/*GetBadgesConflict handles this case with default header values.

Returned if the given user is already subscribed to this stream using a different subscription ID, OR if the given subscription ID is already used to identify a subscription to a different stream.
*/
type GetBadgesConflict struct {
}

func (o *GetBadgesConflict) Error() string {
	return fmt.Sprintf("[GET /1/user/-/badges.json][%d] getBadgesConflict ", 409)
}

func (o *GetBadgesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}