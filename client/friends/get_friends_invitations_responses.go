// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetFriendsInvitationsReader is a Reader for the GetFriendsInvitations structure.
type GetFriendsInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFriendsInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFriendsInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewGetFriendsInvitationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetFriendsInvitationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFriendsInvitationsOK creates a GetFriendsInvitationsOK with default headers values
func NewGetFriendsInvitationsOK() *GetFriendsInvitationsOK {
	return &GetFriendsInvitationsOK{}
}

/*GetFriendsInvitationsOK handles this case with default header values.

Successful request.
*/
type GetFriendsInvitationsOK struct {
}

func (o *GetFriendsInvitationsOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/invitations.json][%d] getFriendsInvitationsOK ", 200)
}

func (o *GetFriendsInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFriendsInvitationsCreated creates a GetFriendsInvitationsCreated with default headers values
func NewGetFriendsInvitationsCreated() *GetFriendsInvitationsCreated {
	return &GetFriendsInvitationsCreated{}
}

/*GetFriendsInvitationsCreated handles this case with default header values.

The request has been fulfilled and resulted in a new resource being created.
*/
type GetFriendsInvitationsCreated struct {
}

func (o *GetFriendsInvitationsCreated) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/invitations.json][%d] getFriendsInvitationsCreated ", 201)
}

func (o *GetFriendsInvitationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFriendsInvitationsUnauthorized creates a GetFriendsInvitationsUnauthorized with default headers values
func NewGetFriendsInvitationsUnauthorized() *GetFriendsInvitationsUnauthorized {
	return &GetFriendsInvitationsUnauthorized{}
}

/*GetFriendsInvitationsUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetFriendsInvitationsUnauthorized struct {
}

func (o *GetFriendsInvitationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/invitations.json][%d] getFriendsInvitationsUnauthorized ", 401)
}

func (o *GetFriendsInvitationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}