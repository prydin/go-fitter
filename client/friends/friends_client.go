// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new friends API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for friends API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFriends gets friends

Returns data of a user's friends in the format requested using units in the unit system which corresponds to the Accept-Language header provided.
*/
func (a *Client) GetFriends(params *GetFriendsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFriendsOK, *GetFriendsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFriendsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFriends",
		Method:             "GET",
		PathPattern:        "/1/user/-/friends.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFriendsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetFriendsOK:
		return value, nil, nil
	case *GetFriendsCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetFriendsInvitations invites friend

Creates an invitation to become friends with the authorized user.
*/
func (a *Client) GetFriendsInvitations(params *GetFriendsInvitationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFriendsInvitationsOK, *GetFriendsInvitationsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFriendsInvitationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFriendsInvitations",
		Method:             "GET",
		PathPattern:        "/1/user/-/friends/invitations.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFriendsInvitationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetFriendsInvitationsOK:
		return value, nil, nil
	case *GetFriendsInvitationsCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetFriendsLeaderboard gets friends leaderboard

Returns data of a user's friends in the format requested using units in the unit system which corresponds to the Accept-Language header provided.
*/
func (a *Client) GetFriendsLeaderboard(params *GetFriendsLeaderboardParams, authInfo runtime.ClientAuthInfoWriter) (*GetFriendsLeaderboardOK, *GetFriendsLeaderboardCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFriendsLeaderboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFriendsLeaderboard",
		Method:             "GET",
		PathPattern:        "/1/user/-/friends/leaderboard.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFriendsLeaderboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetFriendsLeaderboardOK:
		return value, nil, nil
	case *GetFriendsLeaderboardCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
RespondFriendsInvitation responds to friend invitation

Accepts or rejects an invitation to become friends wit inviting user.
*/
func (a *Client) RespondFriendsInvitation(params *RespondFriendsInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*RespondFriendsInvitationOK, *RespondFriendsInvitationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRespondFriendsInvitationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "respondFriendsInvitation",
		Method:             "GET",
		PathPattern:        "/1/user/-/friends/invitations/{from-user-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RespondFriendsInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RespondFriendsInvitationOK:
		return value, nil, nil
	case *RespondFriendsInvitationCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}