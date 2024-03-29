// Code generated by go-swagger; DO NOT EDIT.

package stop_point

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

// NewStopPointArrivalsParams creates a new StopPointArrivalsParams object
// with the default values initialized.
func NewStopPointArrivalsParams() *StopPointArrivalsParams {
	var ()
	return &StopPointArrivalsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopPointArrivalsParamsWithTimeout creates a new StopPointArrivalsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopPointArrivalsParamsWithTimeout(timeout time.Duration) *StopPointArrivalsParams {
	var ()
	return &StopPointArrivalsParams{

		timeout: timeout,
	}
}

// NewStopPointArrivalsParamsWithContext creates a new StopPointArrivalsParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopPointArrivalsParamsWithContext(ctx context.Context) *StopPointArrivalsParams {
	var ()
	return &StopPointArrivalsParams{

		Context: ctx,
	}
}

// NewStopPointArrivalsParamsWithHTTPClient creates a new StopPointArrivalsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopPointArrivalsParamsWithHTTPClient(client *http.Client) *StopPointArrivalsParams {
	var ()
	return &StopPointArrivalsParams{
		HTTPClient: client,
	}
}

/*StopPointArrivalsParams contains all the parameters to send to the API endpoint
for the stop point arrivals operation typically these are written to a http.Request
*/
type StopPointArrivalsParams struct {

	/*ID
	  A StopPoint id (station naptan code e.g. 940GZZLUASL, you can use /StopPoint/Search/{query} endpoint to find a stop point id from a station name)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop point arrivals params
func (o *StopPointArrivalsParams) WithTimeout(timeout time.Duration) *StopPointArrivalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop point arrivals params
func (o *StopPointArrivalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop point arrivals params
func (o *StopPointArrivalsParams) WithContext(ctx context.Context) *StopPointArrivalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop point arrivals params
func (o *StopPointArrivalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop point arrivals params
func (o *StopPointArrivalsParams) WithHTTPClient(client *http.Client) *StopPointArrivalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop point arrivals params
func (o *StopPointArrivalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the stop point arrivals params
func (o *StopPointArrivalsParams) WithID(id string) *StopPointArrivalsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stop point arrivals params
func (o *StopPointArrivalsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StopPointArrivalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
