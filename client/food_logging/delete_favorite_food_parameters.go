// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteFavoriteFoodParams creates a new DeleteFavoriteFoodParams object
// with the default values initialized.
func NewDeleteFavoriteFoodParams() *DeleteFavoriteFoodParams {
	var ()
	return &DeleteFavoriteFoodParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFavoriteFoodParamsWithTimeout creates a new DeleteFavoriteFoodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFavoriteFoodParamsWithTimeout(timeout time.Duration) *DeleteFavoriteFoodParams {
	var ()
	return &DeleteFavoriteFoodParams{

		timeout: timeout,
	}
}

// NewDeleteFavoriteFoodParamsWithContext creates a new DeleteFavoriteFoodParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFavoriteFoodParamsWithContext(ctx context.Context) *DeleteFavoriteFoodParams {
	var ()
	return &DeleteFavoriteFoodParams{

		Context: ctx,
	}
}

// NewDeleteFavoriteFoodParamsWithHTTPClient creates a new DeleteFavoriteFoodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFavoriteFoodParamsWithHTTPClient(client *http.Client) *DeleteFavoriteFoodParams {
	var ()
	return &DeleteFavoriteFoodParams{
		HTTPClient: client,
	}
}

/*DeleteFavoriteFoodParams contains all the parameters to send to the API endpoint
for the delete favorite food operation typically these are written to a http.Request
*/
type DeleteFavoriteFoodParams struct {

	/*FoodID
	  The ID of the food to be deleted from user's favorites.

	*/
	FoodID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete favorite food params
func (o *DeleteFavoriteFoodParams) WithTimeout(timeout time.Duration) *DeleteFavoriteFoodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete favorite food params
func (o *DeleteFavoriteFoodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete favorite food params
func (o *DeleteFavoriteFoodParams) WithContext(ctx context.Context) *DeleteFavoriteFoodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete favorite food params
func (o *DeleteFavoriteFoodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete favorite food params
func (o *DeleteFavoriteFoodParams) WithHTTPClient(client *http.Client) *DeleteFavoriteFoodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete favorite food params
func (o *DeleteFavoriteFoodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFoodID adds the foodID to the delete favorite food params
func (o *DeleteFavoriteFoodParams) WithFoodID(foodID string) *DeleteFavoriteFoodParams {
	o.SetFoodID(foodID)
	return o
}

// SetFoodID adds the foodId to the delete favorite food params
func (o *DeleteFavoriteFoodParams) SetFoodID(foodID string) {
	o.FoodID = foodID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFavoriteFoodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param food-id
	if err := r.SetPathParam("food-id", o.FoodID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
