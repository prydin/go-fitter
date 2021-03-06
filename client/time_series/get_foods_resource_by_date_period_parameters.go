// Code generated by go-swagger; DO NOT EDIT.

package time_series

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

// NewGetFoodsResourceByDatePeriodParams creates a new GetFoodsResourceByDatePeriodParams object
// with the default values initialized.
func NewGetFoodsResourceByDatePeriodParams() *GetFoodsResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("caloriesIn")
	)
	return &GetFoodsResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFoodsResourceByDatePeriodParamsWithTimeout creates a new GetFoodsResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFoodsResourceByDatePeriodParamsWithTimeout(timeout time.Duration) *GetFoodsResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("caloriesIn")
	)
	return &GetFoodsResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		timeout: timeout,
	}
}

// NewGetFoodsResourceByDatePeriodParamsWithContext creates a new GetFoodsResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFoodsResourceByDatePeriodParamsWithContext(ctx context.Context) *GetFoodsResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("caloriesIn")
	)
	return &GetFoodsResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		Context: ctx,
	}
}

// NewGetFoodsResourceByDatePeriodParamsWithHTTPClient creates a new GetFoodsResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFoodsResourceByDatePeriodParamsWithHTTPClient(client *http.Client) *GetFoodsResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("caloriesIn")
	)
	return &GetFoodsResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,
		HTTPClient:   client,
	}
}

/*GetFoodsResourceByDatePeriodParams contains all the parameters to send to the API endpoint
for the get foods resource by date period operation typically these are written to a http.Request
*/
type GetFoodsResourceByDatePeriodParams struct {

	/*Date
	  The end date of the period specified in the format yyyy-MM-dd or today.

	*/
	Date strfmt.Date
	/*Period
	  The range for which data will be returned. Options are 1d, 7d, 30d, 1w, 3m, 6m, 1y, or max.

	*/
	Period string
	/*ResourcePath
	  The resouce path. See options in the Resouce Path Options section in the full documentation.

	*/
	ResourcePath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) WithTimeout(timeout time.Duration) *GetFoodsResourceByDatePeriodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) WithContext(ctx context.Context) *GetFoodsResourceByDatePeriodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) WithHTTPClient(client *http.Client) *GetFoodsResourceByDatePeriodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) WithDate(date strfmt.Date) *GetFoodsResourceByDatePeriodParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) SetDate(date strfmt.Date) {
	o.Date = date
}

// WithPeriod adds the period to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) WithPeriod(period string) *GetFoodsResourceByDatePeriodParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) SetPeriod(period string) {
	o.Period = period
}

// WithResourcePath adds the resourcePath to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) WithResourcePath(resourcePath string) *GetFoodsResourceByDatePeriodParams {
	o.SetResourcePath(resourcePath)
	return o
}

// SetResourcePath adds the resourcePath to the get foods resource by date period params
func (o *GetFoodsResourceByDatePeriodParams) SetResourcePath(resourcePath string) {
	o.ResourcePath = resourcePath
}

// WriteToRequest writes these params to a swagger request
func (o *GetFoodsResourceByDatePeriodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param date
	if err := r.SetPathParam("date", o.Date.String()); err != nil {
		return err
	}

	// path param period
	if err := r.SetPathParam("period", o.Period); err != nil {
		return err
	}

	// path param resource-path
	if err := r.SetPathParam("resource-path", o.ResourcePath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
