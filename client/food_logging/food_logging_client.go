// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new food logging API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for food logging API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddFavoriteFood adds favorite food

Updates a user's daily activity goals and returns a response using units in the unit system which corresponds to the Accept-Language header provided.
*/
func (a *Client) AddFavoriteFood(params *AddFavoriteFoodParams, authInfo runtime.ClientAuthInfoWriter) (*AddFavoriteFoodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddFavoriteFoodParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addFavoriteFood",
		Method:             "POST",
		PathPattern:        "/1/user/-/foods/log/favorite/{food-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddFavoriteFoodReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddFavoriteFoodOK), nil

}

/*
AddFoods creates food

Creates a new private food for a user and returns a response in the format requested. The created food is found via the Search Foods call.
*/
func (a *Client) AddFoods(params *AddFoodsParams, authInfo runtime.ClientAuthInfoWriter) (*AddFoodsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddFoodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addFoods",
		Method:             "POST",
		PathPattern:        "/1/user/-/foods.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddFoodsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddFoodsCreated), nil

}

/*
AddFoodsLog logs food

Creates food log entries for users with or without foodId value.
*/
func (a *Client) AddFoodsLog(params *AddFoodsLogParams, authInfo runtime.ClientAuthInfoWriter) (*AddFoodsLogCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddFoodsLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addFoodsLog",
		Method:             "POST",
		PathPattern:        "/1/user/-/foods/log.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddFoodsLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddFoodsLogCreated), nil

}

/*
AddMeal creates meal

Creates a meal with the given food contained in the post body.
*/
func (a *Client) AddMeal(params *AddMealParams, authInfo runtime.ClientAuthInfoWriter) (*AddMealOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMealParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addMeal",
		Method:             "POST",
		PathPattern:        "/1/user/-/meals.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddMealReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddMealOK), nil

}

/*
AddUpdateFoodsGoal updates food goal

Updates a user's daily calories consumption goal or food plan and returns a response in the format requested.
*/
func (a *Client) AddUpdateFoodsGoal(params *AddUpdateFoodsGoalParams, authInfo runtime.ClientAuthInfoWriter) (*AddUpdateFoodsGoalCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUpdateFoodsGoalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addUpdateFoodsGoal",
		Method:             "POST",
		PathPattern:        "/1/user/-/foods/log/goal.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddUpdateFoodsGoalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddUpdateFoodsGoalCreated), nil

}

/*
AddUpdateWaterGoal updates water goal

Updates a user's daily calories consumption goal or food plan and returns a response in the format requested.
*/
func (a *Client) AddUpdateWaterGoal(params *AddUpdateWaterGoalParams, authInfo runtime.ClientAuthInfoWriter) (*AddUpdateWaterGoalCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUpdateWaterGoalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addUpdateWaterGoal",
		Method:             "POST",
		PathPattern:        "/1/user/-/foods/log/water/goal.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddUpdateWaterGoalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddUpdateWaterGoalCreated), nil

}

/*
AddWaterLog logs water

Creates a log entry for water using units in the unit systems that corresponds to the Accept-Language header provided.
*/
func (a *Client) AddWaterLog(params *AddWaterLogParams, authInfo runtime.ClientAuthInfoWriter) (*AddWaterLogCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddWaterLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addWaterLog",
		Method:             "POST",
		PathPattern:        "/1/user/-/foods/log/water.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddWaterLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddWaterLogCreated), nil

}

/*
DeleteFavoriteFood deletes favorite food

Deletes a food with the given ID to the user's list of favorite foods.
*/
func (a *Client) DeleteFavoriteFood(params *DeleteFavoriteFoodParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFavoriteFoodNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFavoriteFoodParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFavoriteFood",
		Method:             "DELETE",
		PathPattern:        "/1/user/-/foods/log/favorite/{food-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFavoriteFoodReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFavoriteFoodNoContent), nil

}

/*
DeleteFoods deletes custom food

Deletes custom food for a user and returns a response in the format requested.
*/
func (a *Client) DeleteFoods(params *DeleteFoodsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFoodsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFoodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFoods",
		Method:             "DELETE",
		PathPattern:        "/1/user/-/foods/{food-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFoodsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFoodsNoContent), nil

}

/*
DeleteFoodsLog deletes food log

Deletes a user's food log entry with the given ID.
*/
func (a *Client) DeleteFoodsLog(params *DeleteFoodsLogParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFoodsLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFoodsLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFoodsLog",
		Method:             "DELETE",
		PathPattern:        "/1/user/-/foods/log/{food-log-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFoodsLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFoodsLogOK), nil

}

/*
DeleteMeal deletes meal

Deletes a user's meal with the given meal id.
*/
func (a *Client) DeleteMeal(params *DeleteMealParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteMealOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMealParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMeal",
		Method:             "DELETE",
		PathPattern:        "/1/user/-/meals/{meal-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMealReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMealOK), nil

}

/*
DeleteWaterLog deletes water log

Deletes a user's water log entry with the given ID.
*/
func (a *Client) DeleteWaterLog(params *DeleteWaterLogParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWaterLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWaterLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWaterLog",
		Method:             "DELETE",
		PathPattern:        "/1/user/-/foods/log/water/{water-log-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWaterLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWaterLogOK), nil

}

/*
GetFavoriteFoods gets favorite foods

Returns a list of a user's favorite foods in the format requested. A favorite food in the list provides a quick way to log the food via the Log Food endpoint.
*/
func (a *Client) GetFavoriteFoods(params *GetFavoriteFoodsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFavoriteFoodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFavoriteFoodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFavoriteFoods",
		Method:             "GET",
		PathPattern:        "/1/user/-/foods/log/favorite.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFavoriteFoodsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFavoriteFoodsOK), nil

}

/*
GetFoodsByDate gets food logs

Retreives a summary and list of a user's food log entries for a given day in the format requested.
*/
func (a *Client) GetFoodsByDate(params *GetFoodsByDateParams, authInfo runtime.ClientAuthInfoWriter) (*GetFoodsByDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoodsByDateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFoodsByDate",
		Method:             "GET",
		PathPattern:        "/1/user/-/foods/log/date/{date}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoodsByDateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFoodsByDateOK), nil

}

/*
GetFoodsGoal gets food goals

Returns a user's current daily calorie consumption goal and/or foodPlan value in the format requested.
*/
func (a *Client) GetFoodsGoal(params *GetFoodsGoalParams, authInfo runtime.ClientAuthInfoWriter) (*GetFoodsGoalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoodsGoalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFoodsGoal",
		Method:             "GET",
		PathPattern:        "/1/user/-/foods/log/goal.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoodsGoalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFoodsGoalOK), nil

}

/*
GetFoodsInfo gets food

Returns the details of a specific food in the Fitbit food databases or a private food that an authorized user has entered in the format requested.
*/
func (a *Client) GetFoodsInfo(params *GetFoodsInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetFoodsInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoodsInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFoodsInfo",
		Method:             "GET",
		PathPattern:        "/1/foods/{food-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoodsInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFoodsInfoOK), nil

}

/*
GetFoodsList searches foods

Returns a list of public foods from the Fitbit food database and private food the user created in the format requested.
*/
func (a *Client) GetFoodsList(params *GetFoodsListParams, authInfo runtime.ClientAuthInfoWriter) (*GetFoodsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoodsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFoodsList",
		Method:             "GET",
		PathPattern:        "/1/foods/search.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoodsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFoodsListOK), nil

}

/*
GetFoodsLocales gets food locales

Returns the food locales that the user may choose to search, log, and create food in.
*/
func (a *Client) GetFoodsLocales(params *GetFoodsLocalesParams, authInfo runtime.ClientAuthInfoWriter) (*GetFoodsLocalesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoodsLocalesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFoodsLocales",
		Method:             "GET",
		PathPattern:        "/1/foods/locales.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoodsLocalesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFoodsLocalesOK), nil

}

/*
GetFoodsUnits gets food units

Returns a list of all valid Fitbit food units in the format requested.
*/
func (a *Client) GetFoodsUnits(params *GetFoodsUnitsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFoodsUnitsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoodsUnitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFoodsUnits",
		Method:             "GET",
		PathPattern:        "/1/foods/units.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFoodsUnitsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFoodsUnitsOK), nil

}

/*
GetFrequentFoods gets frequent foods

Returns a list of a user's frequent foods in the format requested. A frequent food in the list provides a quick way to log the food via the Log Food endpoint.
*/
func (a *Client) GetFrequentFoods(params *GetFrequentFoodsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFrequentFoodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFrequentFoodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFrequentFoods",
		Method:             "GET",
		PathPattern:        "/1/user/-/foods/log/frequent.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFrequentFoodsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFrequentFoodsOK), nil

}

/*
GetMeals gets meals

Returns a list of meals created by user in the user's food log in the format requested. User creates and manages meals on the Food Log tab on the website.
*/
func (a *Client) GetMeals(params *GetMealsParams, authInfo runtime.ClientAuthInfoWriter) (*GetMealsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMealsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMeals",
		Method:             "GET",
		PathPattern:        "/1/user/-/meals.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMealsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMealsOK), nil

}

/*
GetRecentFoods gets recent foods

Returns a list of a user's frequent foods in the format requested. A frequent food in the list provides a quick way to log the food via the Log Food endpoint.
*/
func (a *Client) GetRecentFoods(params *GetRecentFoodsParams, authInfo runtime.ClientAuthInfoWriter) (*GetRecentFoodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecentFoodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRecentFoods",
		Method:             "GET",
		PathPattern:        "/1/user/-/foods/log/recent.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRecentFoodsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecentFoodsOK), nil

}

/*
GetWaterByDate gets water logs

Retreives a summary and list of a user's water log entries for a given day in the requested using units in the unit system that corresponds to the Accept-Language header provided.
*/
func (a *Client) GetWaterByDate(params *GetWaterByDateParams, authInfo runtime.ClientAuthInfoWriter) (*GetWaterByDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWaterByDateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWaterByDate",
		Method:             "GET",
		PathPattern:        "/1/user/-/foods/log/water/date/{date}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWaterByDateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWaterByDateOK), nil

}

/*
GetWaterGoal gets water goal

Retreives a summary and list of a user's water goal entries for a given day in the requested using units in the unit system that corresponds to the Accept-Language header provided.
*/
func (a *Client) GetWaterGoal(params *GetWaterGoalParams, authInfo runtime.ClientAuthInfoWriter) (*GetWaterGoalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWaterGoalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWaterGoal",
		Method:             "GET",
		PathPattern:        "/1/user/-/foods/log/water/goal.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWaterGoalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWaterGoalOK), nil

}

/*
UpdateMeal edits meal

Replaces an existing meal with the contents of the request. The response contains the updated meal.
*/
func (a *Client) UpdateMeal(params *UpdateMealParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateMealOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMealParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMeal",
		Method:             "POST",
		PathPattern:        "/1/user/-/meals/{meal-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMealReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateMealOK), nil

}

/*
UpdateWaterLog updates water log

Updates a user's water log entry with the given ID.
*/
func (a *Client) UpdateWaterLog(params *UpdateWaterLogParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWaterLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWaterLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateWaterLog",
		Method:             "POST",
		PathPattern:        "/1/user/-/foods/log/water/{water-log-id}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWaterLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateWaterLogOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
