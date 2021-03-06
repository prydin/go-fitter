// Code generated by go-swagger; DO NOT EDIT.

package activity

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

// NewAddUpdateActivitiesGoalsParams creates a new AddUpdateActivitiesGoalsParams object
// with the default values initialized.
func NewAddUpdateActivitiesGoalsParams() *AddUpdateActivitiesGoalsParams {
	var ()
	return &AddUpdateActivitiesGoalsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddUpdateActivitiesGoalsParamsWithTimeout creates a new AddUpdateActivitiesGoalsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddUpdateActivitiesGoalsParamsWithTimeout(timeout time.Duration) *AddUpdateActivitiesGoalsParams {
	var ()
	return &AddUpdateActivitiesGoalsParams{

		timeout: timeout,
	}
}

// NewAddUpdateActivitiesGoalsParamsWithContext creates a new AddUpdateActivitiesGoalsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddUpdateActivitiesGoalsParamsWithContext(ctx context.Context) *AddUpdateActivitiesGoalsParams {
	var ()
	return &AddUpdateActivitiesGoalsParams{

		Context: ctx,
	}
}

// NewAddUpdateActivitiesGoalsParamsWithHTTPClient creates a new AddUpdateActivitiesGoalsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddUpdateActivitiesGoalsParamsWithHTTPClient(client *http.Client) *AddUpdateActivitiesGoalsParams {
	var ()
	return &AddUpdateActivitiesGoalsParams{
		HTTPClient: client,
	}
}

/*AddUpdateActivitiesGoalsParams contains all the parameters to send to the API endpoint
for the add update activities goals operation typically these are written to a http.Request
*/
type AddUpdateActivitiesGoalsParams struct {

	/*Period
	  daily or weekly.

	*/
	Period string
	/*Type
	  goal type

	*/
	Type string
	/*Value
	  goal value

	*/
	Value string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) WithTimeout(timeout time.Duration) *AddUpdateActivitiesGoalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) WithContext(ctx context.Context) *AddUpdateActivitiesGoalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) WithHTTPClient(client *http.Client) *AddUpdateActivitiesGoalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeriod adds the period to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) WithPeriod(period string) *AddUpdateActivitiesGoalsParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) SetPeriod(period string) {
	o.Period = period
}

// WithType adds the typeVar to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) WithType(typeVar string) *AddUpdateActivitiesGoalsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithValue adds the value to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) WithValue(value string) *AddUpdateActivitiesGoalsParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the add update activities goals params
func (o *AddUpdateActivitiesGoalsParams) SetValue(value string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *AddUpdateActivitiesGoalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param period
	if err := r.SetPathParam("period", o.Period); err != nil {
		return err
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	// query param value
	qrValue := o.Value
	qValue := qrValue
	if qValue != "" {
		if err := r.SetQueryParam("value", qValue); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
