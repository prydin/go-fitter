// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetFriendsLeaderboardReader is a Reader for the GetFriendsLeaderboard structure.
type GetFriendsLeaderboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFriendsLeaderboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFriendsLeaderboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewGetFriendsLeaderboardCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetFriendsLeaderboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFriendsLeaderboardOK creates a GetFriendsLeaderboardOK with default headers values
func NewGetFriendsLeaderboardOK() *GetFriendsLeaderboardOK {
	return &GetFriendsLeaderboardOK{}
}

/*GetFriendsLeaderboardOK handles this case with default header values.

Successful request.
*/
type GetFriendsLeaderboardOK struct {
}

func (o *GetFriendsLeaderboardOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/leaderboard.json][%d] getFriendsLeaderboardOK ", 200)
}

func (o *GetFriendsLeaderboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFriendsLeaderboardCreated creates a GetFriendsLeaderboardCreated with default headers values
func NewGetFriendsLeaderboardCreated() *GetFriendsLeaderboardCreated {
	return &GetFriendsLeaderboardCreated{}
}

/*GetFriendsLeaderboardCreated handles this case with default header values.

The request has been fulfilled and resulted in a new resource being created.
*/
type GetFriendsLeaderboardCreated struct {
}

func (o *GetFriendsLeaderboardCreated) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/leaderboard.json][%d] getFriendsLeaderboardCreated ", 201)
}

func (o *GetFriendsLeaderboardCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFriendsLeaderboardUnauthorized creates a GetFriendsLeaderboardUnauthorized with default headers values
func NewGetFriendsLeaderboardUnauthorized() *GetFriendsLeaderboardUnauthorized {
	return &GetFriendsLeaderboardUnauthorized{}
}

/*GetFriendsLeaderboardUnauthorized handles this case with default header values.

The request requires user authentication.
*/
type GetFriendsLeaderboardUnauthorized struct {
}

func (o *GetFriendsLeaderboardUnauthorized) Error() string {
	return fmt.Sprintf("[GET /1/user/-/friends/leaderboard.json][%d] getFriendsLeaderboardUnauthorized ", 401)
}

func (o *GetFriendsLeaderboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
