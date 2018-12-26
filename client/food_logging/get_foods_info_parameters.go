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

// NewGetFoodsInfoParams creates a new GetFoodsInfoParams object
// with the default values initialized.
func NewGetFoodsInfoParams() *GetFoodsInfoParams {
	var ()
	return &GetFoodsInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFoodsInfoParamsWithTimeout creates a new GetFoodsInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFoodsInfoParamsWithTimeout(timeout time.Duration) *GetFoodsInfoParams {
	var ()
	return &GetFoodsInfoParams{

		timeout: timeout,
	}
}

// NewGetFoodsInfoParamsWithContext creates a new GetFoodsInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFoodsInfoParamsWithContext(ctx context.Context) *GetFoodsInfoParams {
	var ()
	return &GetFoodsInfoParams{

		Context: ctx,
	}
}

// NewGetFoodsInfoParamsWithHTTPClient creates a new GetFoodsInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFoodsInfoParamsWithHTTPClient(client *http.Client) *GetFoodsInfoParams {
	var ()
	return &GetFoodsInfoParams{
		HTTPClient: client,
	}
}

/*GetFoodsInfoParams contains all the parameters to send to the API endpoint
for the get foods info operation typically these are written to a http.Request
*/
type GetFoodsInfoParams struct {

	/*FoodID
	  The ID of the food.

	*/
	FoodID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get foods info params
func (o *GetFoodsInfoParams) WithTimeout(timeout time.Duration) *GetFoodsInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get foods info params
func (o *GetFoodsInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get foods info params
func (o *GetFoodsInfoParams) WithContext(ctx context.Context) *GetFoodsInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get foods info params
func (o *GetFoodsInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get foods info params
func (o *GetFoodsInfoParams) WithHTTPClient(client *http.Client) *GetFoodsInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get foods info params
func (o *GetFoodsInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFoodID adds the foodID to the get foods info params
func (o *GetFoodsInfoParams) WithFoodID(foodID string) *GetFoodsInfoParams {
	o.SetFoodID(foodID)
	return o
}

// SetFoodID adds the foodId to the get foods info params
func (o *GetFoodsInfoParams) SetFoodID(foodID string) {
	o.FoodID = foodID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFoodsInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
