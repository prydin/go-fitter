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

// NewUpdateWaterLogParams creates a new UpdateWaterLogParams object
// with the default values initialized.
func NewUpdateWaterLogParams() *UpdateWaterLogParams {
	var ()
	return &UpdateWaterLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWaterLogParamsWithTimeout creates a new UpdateWaterLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWaterLogParamsWithTimeout(timeout time.Duration) *UpdateWaterLogParams {
	var ()
	return &UpdateWaterLogParams{

		timeout: timeout,
	}
}

// NewUpdateWaterLogParamsWithContext creates a new UpdateWaterLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWaterLogParamsWithContext(ctx context.Context) *UpdateWaterLogParams {
	var ()
	return &UpdateWaterLogParams{

		Context: ctx,
	}
}

// NewUpdateWaterLogParamsWithHTTPClient creates a new UpdateWaterLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWaterLogParamsWithHTTPClient(client *http.Client) *UpdateWaterLogParams {
	var ()
	return &UpdateWaterLogParams{
		HTTPClient: client,
	}
}

/*UpdateWaterLogParams contains all the parameters to send to the API endpoint
for the update water log operation typically these are written to a http.Request
*/
type UpdateWaterLogParams struct {

	/*Amount
	  Amount consumed; in the format X.X and in the specified waterUnit or in the unit system that corresponds to the Accept-Language header provided.

	*/
	Amount string
	/*Unit
	  Water measurement unit. 'ml', 'fl oz', or 'cup'.

	*/
	Unit *string
	/*WaterLogID
	  The ID of the waterUnit log entry to be deleted.

	*/
	WaterLogID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update water log params
func (o *UpdateWaterLogParams) WithTimeout(timeout time.Duration) *UpdateWaterLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update water log params
func (o *UpdateWaterLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update water log params
func (o *UpdateWaterLogParams) WithContext(ctx context.Context) *UpdateWaterLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update water log params
func (o *UpdateWaterLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update water log params
func (o *UpdateWaterLogParams) WithHTTPClient(client *http.Client) *UpdateWaterLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update water log params
func (o *UpdateWaterLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the update water log params
func (o *UpdateWaterLogParams) WithAmount(amount string) *UpdateWaterLogParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the update water log params
func (o *UpdateWaterLogParams) SetAmount(amount string) {
	o.Amount = amount
}

// WithUnit adds the unit to the update water log params
func (o *UpdateWaterLogParams) WithUnit(unit *string) *UpdateWaterLogParams {
	o.SetUnit(unit)
	return o
}

// SetUnit adds the unit to the update water log params
func (o *UpdateWaterLogParams) SetUnit(unit *string) {
	o.Unit = unit
}

// WithWaterLogID adds the waterLogID to the update water log params
func (o *UpdateWaterLogParams) WithWaterLogID(waterLogID string) *UpdateWaterLogParams {
	o.SetWaterLogID(waterLogID)
	return o
}

// SetWaterLogID adds the waterLogId to the update water log params
func (o *UpdateWaterLogParams) SetWaterLogID(waterLogID string) {
	o.WaterLogID = waterLogID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWaterLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param amount
	qrAmount := o.Amount
	qAmount := qrAmount
	if qAmount != "" {
		if err := r.SetQueryParam("amount", qAmount); err != nil {
			return err
		}
	}

	if o.Unit != nil {

		// query param unit
		var qrUnit string
		if o.Unit != nil {
			qrUnit = *o.Unit
		}
		qUnit := qrUnit
		if qUnit != "" {
			if err := r.SetQueryParam("unit", qUnit); err != nil {
				return err
			}
		}

	}

	// path param water-log-id
	if err := r.SetPathParam("water-log-id", o.WaterLogID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}