// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new activity API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for activity API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddActivitiesLog logs activity

The Log Activity endpoint creates log entry for an activity or user's private custom activity using units in the unit system which corresponds to the Accept-Language header provided (or using optional custom distanceUnit) and get a response in the format requested.
*/
func (a *Client) AddActivitiesLog(params *AddActivitiesLogParams, authInfo runtime.ClientAuthInfoWriter) (*AddActivitiesLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddActivitiesLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addActivitiesLog",
		Method:             "POST",
		PathPattern:        "/1/user/-/activities.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddActivitiesLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddActivitiesLogOK), nil

}

/*
AddFavoriteActivities adds favorite activity

Adds the activity with the given ID to user's list of favorite activities.
*/
func (a *Client) AddFavoriteActivities(params *AddFavoriteActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*AddFavoriteActivitiesOK, *AddFavoriteActivitiesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddFavoriteActivitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addFavoriteActivities",
		Method:             "POST",
		PathPattern:        "/1/user/-/activities/favorite/{activity-id}.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddFavoriteActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *AddFavoriteActivitiesOK:
		return value, nil, nil
	case *AddFavoriteActivitiesCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
AddUpdateActivitiesGoals updates activity goals

Updates a user's daily or weekly activity goals and returns a response using units in the unit system which corresponds to the Accept-Language header provided.
*/
func (a *Client) AddUpdateActivitiesGoals(params *AddUpdateActivitiesGoalsParams, authInfo runtime.ClientAuthInfoWriter) (*AddUpdateActivitiesGoalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUpdateActivitiesGoalsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addUpdateActivitiesGoals",
		Method:             "POST",
		PathPattern:        "/1/user/-/activities/goals/{period}.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddUpdateActivitiesGoalsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddUpdateActivitiesGoalsOK), nil

}

/*
DeleteActivitiesLog deletes activity log

Deletes a user's activity log entry with the given ID.
*/
func (a *Client) DeleteActivitiesLog(params *DeleteActivitiesLogParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteActivitiesLogNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteActivitiesLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteActivitiesLog",
		Method:             "DELETE",
		PathPattern:        "/1/user/-/activities/{activity-log-id}.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteActivitiesLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteActivitiesLogNoContent), nil

}

/*
DeleteFavoriteActivities deletes favorite activity

Removes the activity with the given ID from a user's list of favorite activities.
*/
func (a *Client) DeleteFavoriteActivities(params *DeleteFavoriteActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFavoriteActivitiesOK, *DeleteFavoriteActivitiesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFavoriteActivitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFavoriteActivities",
		Method:             "DELETE",
		PathPattern:        "/1/user/-/activities/favorite/{activity-id}.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFavoriteActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteFavoriteActivitiesOK:
		return value, nil, nil
	case *DeleteFavoriteActivitiesNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetActivitiesByDate gets daily activity summary

Retrieves a summary and list of a user's activities and activity log entries for a given day.
*/
func (a *Client) GetActivitiesByDate(params *GetActivitiesByDateParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivitiesByDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivitiesByDateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivitiesByDate",
		Method:             "GET",
		PathPattern:        "/1/user/-/activities/date/{date}.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivitiesByDateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivitiesByDateOK), nil

}

/*
GetActivitiesGoals gets activity goals

Retreives a user's current daily or weekly activity goals using measurement units as defined in the unit system, which corresponds to the Accept-Language header provided.
*/
func (a *Client) GetActivitiesGoals(params *GetActivitiesGoalsParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivitiesGoalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivitiesGoalsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivitiesGoals",
		Method:             "GET",
		PathPattern:        "/1/user/-/activities/goals/{period}.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivitiesGoalsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivitiesGoalsOK), nil

}

/*
GetActivitiesLog gets lifetime stats

Updates a user's daily activity goals and returns a response using units in the unit system which corresponds to the Accept-Language header provided.
*/
func (a *Client) GetActivitiesLog(params *GetActivitiesLogParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivitiesLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivitiesLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivitiesLog",
		Method:             "GET",
		PathPattern:        "/1/user/-/activities.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivitiesLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivitiesLogOK), nil

}

/*
GetActivitiesLogList gets activity log list

Retreives a list of user's activity log entries before or after a given day with offset and limit using units in the unit system which corresponds to the Accept-Language header provided.
*/
func (a *Client) GetActivitiesLogList(params *GetActivitiesLogListParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivitiesLogListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivitiesLogListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivitiesLogList",
		Method:             "GET",
		PathPattern:        "/1/user/-/activities/list.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivitiesLogListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivitiesLogListOK), nil

}

/*
GetActivitiesTCX gets activity t c x

Retreives the details of a user's location and heart rate data during a logged exercise activity.
*/
func (a *Client) GetActivitiesTCX(params *GetActivitiesTCXParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivitiesTCXOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivitiesTCXParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivitiesTCX",
		Method:             "GET",
		PathPattern:        "/1/user/-/activities/{log-id}.tcx",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivitiesTCXReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivitiesTCXOK), nil

}

/*
GetActivitiesTypeDetail gets activity type

Returns the detail of a specific activity in the Fitbit activities database in the format requested. If activity has levels, it also returns a list of activity level details.
*/
func (a *Client) GetActivitiesTypeDetail(params *GetActivitiesTypeDetailParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivitiesTypeDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivitiesTypeDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivitiesTypeDetail",
		Method:             "GET",
		PathPattern:        "/1/activities/{activity-id}.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivitiesTypeDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivitiesTypeDetailOK), nil

}

/*
GetActivitiesTypes browses activity types

Retreives a tree of all valid Fitbit public activities from the activities catelog as well as private custom activities the user created in the format requested.
*/
func (a *Client) GetActivitiesTypes(params *GetActivitiesTypesParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivitiesTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivitiesTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivitiesTypes",
		Method:             "GET",
		PathPattern:        "/1/activities.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivitiesTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActivitiesTypesOK), nil

}

/*
GetFavoriteActivities gets favorite activities

Returns a list of a user's favorite activities.
*/
func (a *Client) GetFavoriteActivities(params *GetFavoriteActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetFavoriteActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFavoriteActivitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFavoriteActivities",
		Method:             "GET",
		PathPattern:        "/1/user/-/activities/favorite.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFavoriteActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFavoriteActivitiesOK), nil

}

/*
GetFrequentActivities gets frequent activities

Retreives a list of a user's frequent activities in the format requested using units in the unit system which corresponds to the Accept-Language header provided.
*/
func (a *Client) GetFrequentActivities(params *GetFrequentActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetFrequentActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFrequentActivitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFrequentActivities",
		Method:             "GET",
		PathPattern:        "/1/user/-/activities/frequent.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFrequentActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFrequentActivitiesOK), nil

}

/*
GetRecentActivities gets recent activity types

Retreives a list of a user's recent activities types logged with some details of the last activity log of that type using units in the unit system which corresponds to the Accept-Language header provided.
*/
func (a *Client) GetRecentActivities(params *GetRecentActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetRecentActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecentActivitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRecentActivities",
		Method:             "GET",
		PathPattern:        "/1/user/-/activities/recent.json",
		ProducesMediaTypes: []string{"application/x-www-form-urlencoded"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRecentActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecentActivitiesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
