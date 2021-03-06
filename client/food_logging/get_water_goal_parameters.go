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

// NewGetWaterGoalParams creates a new GetWaterGoalParams object
// with the default values initialized.
func NewGetWaterGoalParams() *GetWaterGoalParams {

	return &GetWaterGoalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWaterGoalParamsWithTimeout creates a new GetWaterGoalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWaterGoalParamsWithTimeout(timeout time.Duration) *GetWaterGoalParams {

	return &GetWaterGoalParams{

		timeout: timeout,
	}
}

// NewGetWaterGoalParamsWithContext creates a new GetWaterGoalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWaterGoalParamsWithContext(ctx context.Context) *GetWaterGoalParams {

	return &GetWaterGoalParams{

		Context: ctx,
	}
}

// NewGetWaterGoalParamsWithHTTPClient creates a new GetWaterGoalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWaterGoalParamsWithHTTPClient(client *http.Client) *GetWaterGoalParams {

	return &GetWaterGoalParams{
		HTTPClient: client,
	}
}

/*GetWaterGoalParams contains all the parameters to send to the API endpoint
for the get water goal operation typically these are written to a http.Request
*/
type GetWaterGoalParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get water goal params
func (o *GetWaterGoalParams) WithTimeout(timeout time.Duration) *GetWaterGoalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get water goal params
func (o *GetWaterGoalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get water goal params
func (o *GetWaterGoalParams) WithContext(ctx context.Context) *GetWaterGoalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get water goal params
func (o *GetWaterGoalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get water goal params
func (o *GetWaterGoalParams) WithHTTPClient(client *http.Client) *GetWaterGoalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get water goal params
func (o *GetWaterGoalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWaterGoalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
