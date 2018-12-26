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

// NewDeleteMealParams creates a new DeleteMealParams object
// with the default values initialized.
func NewDeleteMealParams() *DeleteMealParams {
	var ()
	return &DeleteMealParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMealParamsWithTimeout creates a new DeleteMealParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMealParamsWithTimeout(timeout time.Duration) *DeleteMealParams {
	var ()
	return &DeleteMealParams{

		timeout: timeout,
	}
}

// NewDeleteMealParamsWithContext creates a new DeleteMealParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMealParamsWithContext(ctx context.Context) *DeleteMealParams {
	var ()
	return &DeleteMealParams{

		Context: ctx,
	}
}

// NewDeleteMealParamsWithHTTPClient creates a new DeleteMealParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMealParamsWithHTTPClient(client *http.Client) *DeleteMealParams {
	var ()
	return &DeleteMealParams{
		HTTPClient: client,
	}
}

/*DeleteMealParams contains all the parameters to send to the API endpoint
for the delete meal operation typically these are written to a http.Request
*/
type DeleteMealParams struct {

	/*MealID
	  Id of the meal to delete.

	*/
	MealID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete meal params
func (o *DeleteMealParams) WithTimeout(timeout time.Duration) *DeleteMealParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete meal params
func (o *DeleteMealParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete meal params
func (o *DeleteMealParams) WithContext(ctx context.Context) *DeleteMealParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete meal params
func (o *DeleteMealParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete meal params
func (o *DeleteMealParams) WithHTTPClient(client *http.Client) *DeleteMealParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete meal params
func (o *DeleteMealParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMealID adds the mealID to the delete meal params
func (o *DeleteMealParams) WithMealID(mealID string) *DeleteMealParams {
	o.SetMealID(mealID)
	return o
}

// SetMealID adds the mealId to the delete meal params
func (o *DeleteMealParams) SetMealID(mealID string) {
	o.MealID = mealID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMealParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param meal-id
	if err := r.SetPathParam("meal-id", o.MealID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}