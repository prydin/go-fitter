// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RespondFriendsInvitationReader is a Reader for the RespondFriendsInvitation structure.
type RespondFriendsInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RespondFriendsInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRespondFriendsInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewRespondFriendsInvitationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRespondFriendsInvitationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRespondFriendsInvitationOK creates a RespondFriendsInvitationOK with default headers values
func NewRespondFriendsInvitationOK() *RespondFriendsInvitationOK {
	return &RespondFriendsInvitationOK{}
}

/*RespondFriendsInvitationOK handles this case with default header values.

Successful request.
*/
type RespondFriendsInvitationOK struct {
}

func (o *RespondFriendsInvitationOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/invitations/{from-user-id}.json][%d] respondFriendsInvitationOK ", 200)
}

func (o *RespondFriendsInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRespondFriendsInvitationCreated creates a RespondFriendsInvitationCreated with default headers values
func NewRespondFriendsInvitationCreated() *RespondFriendsInvitationCreated {
	return &RespondFriendsInvitationCreated{}
}

/*RespondFriendsInvitationCreated handles this case with default header values.

The request has been fulfilled and resulted in a new resource being created.
*/
type RespondFriendsInvitationCreated struct {
}

func (o *RespondFriendsInvitationCreated) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/invitations/{from-user-id}.json][%d] respondFriendsInvitationCreated ", 201)
}

func (o *RespondFriendsInvitationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRespondFriendsInvitationUnauthorized creates a RespondFriendsInvitationUnauthorized with default headers values
func NewRespondFriendsInvitationUnauthorized() *RespondFriendsInvitationUnauthorized {
	return &RespondFriendsInvitationUnauthorized{}
}

/*RespondFriendsInvitationUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type RespondFriendsInvitationUnauthorized struct {
}

func (o *RespondFriendsInvitationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/invitations/{from-user-id}.json][%d] respondFriendsInvitationUnauthorized ", 401)
}

func (o *RespondFriendsInvitationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}