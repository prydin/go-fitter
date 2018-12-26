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

// NewGetActivitiesResourceByDatePeriodParams creates a new GetActivitiesResourceByDatePeriodParams object
// with the default values initialized.
func NewGetActivitiesResourceByDatePeriodParams() *GetActivitiesResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("steps")
	)
	return &GetActivitiesResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActivitiesResourceByDatePeriodParamsWithTimeout creates a new GetActivitiesResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActivitiesResourceByDatePeriodParamsWithTimeout(timeout time.Duration) *GetActivitiesResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("steps")
	)
	return &GetActivitiesResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		timeout: timeout,
	}
}

// NewGetActivitiesResourceByDatePeriodParamsWithContext creates a new GetActivitiesResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActivitiesResourceByDatePeriodParamsWithContext(ctx context.Context) *GetActivitiesResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("steps")
	)
	return &GetActivitiesResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,

		Context: ctx,
	}
}

// NewGetActivitiesResourceByDatePeriodParamsWithHTTPClient creates a new GetActivitiesResourceByDatePeriodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActivitiesResourceByDatePeriodParamsWithHTTPClient(client *http.Client) *GetActivitiesResourceByDatePeriodParams {
	var (
		resourcePathDefault = string("steps")
	)
	return &GetActivitiesResourceByDatePeriodParams{
		ResourcePath: resourcePathDefault,
		HTTPClient:   client,
	}
}

/*GetActivitiesResourceByDatePeriodParams contains all the parameters to send to the API endpoint
for the get activities resource by date period operation typically these are written to a http.Request
*/
type GetActivitiesResourceByDatePeriodParams struct {

	/*Date
	  The end date of the period specified in the format yyyy-MM-dd or today.

	*/
	Date strfmt.Date
	/*Period
	  The range for which data will be returned. Options are 1d, 7d, 30d, 1w, 1m, 3m, 6m, 1y, or max.

	*/
	Period string
	/*ResourcePath
	  The resource-path; see options in the Resource Path Options section in the full documentation.

	*/
	ResourcePath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) WithTimeout(timeout time.Duration) *GetActivitiesResourceByDatePeriodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) WithContext(ctx context.Context) *GetActivitiesResourceByDatePeriodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) WithHTTPClient(client *http.Client) *GetActivitiesResourceByDatePeriodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) WithDate(date strfmt.Date) *GetActivitiesResourceByDatePeriodParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) SetDate(date strfmt.Date) {
	o.Date = date
}

// WithPeriod adds the period to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) WithPeriod(period string) *GetActivitiesResourceByDatePeriodParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) SetPeriod(period string) {
	o.Period = period
}

// WithResourcePath adds the resourcePath to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) WithResourcePath(resourcePath string) *GetActivitiesResourceByDatePeriodParams {
	o.SetResourcePath(resourcePath)
	return o
}

// SetResourcePath adds the resourcePath to the get activities resource by date period params
func (o *GetActivitiesResourceByDatePeriodParams) SetResourcePath(resourcePath string) {
	o.ResourcePath = resourcePath
}

// WriteToRequest writes these params to a swagger request
func (o *GetActivitiesResourceByDatePeriodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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