// Code generated by go-swagger; DO NOT EDIT.

package bike_point

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

// NewBikePointGetParams creates a new BikePointGetParams object
// with the default values initialized.
func NewBikePointGetParams() *BikePointGetParams {
	var ()
	return &BikePointGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBikePointGetParamsWithTimeout creates a new BikePointGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBikePointGetParamsWithTimeout(timeout time.Duration) *BikePointGetParams {
	var ()
	return &BikePointGetParams{

		timeout: timeout,
	}
}

// NewBikePointGetParamsWithContext creates a new BikePointGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewBikePointGetParamsWithContext(ctx context.Context) *BikePointGetParams {
	var ()
	return &BikePointGetParams{

		Context: ctx,
	}
}

// NewBikePointGetParamsWithHTTPClient creates a new BikePointGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBikePointGetParamsWithHTTPClient(client *http.Client) *BikePointGetParams {
	var ()
	return &BikePointGetParams{
		HTTPClient: client,
	}
}

/*BikePointGetParams contains all the parameters to send to the API endpoint
for the bike point get operation typically these are written to a http.Request
*/
type BikePointGetParams struct {

	/*ID
	  A bike point id (a list of ids can be obtained from the above BikePoint call)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bike point get params
func (o *BikePointGetParams) WithTimeout(timeout time.Duration) *BikePointGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bike point get params
func (o *BikePointGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bike point get params
func (o *BikePointGetParams) WithContext(ctx context.Context) *BikePointGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bike point get params
func (o *BikePointGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bike point get params
func (o *BikePointGetParams) WithHTTPClient(client *http.Client) *BikePointGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bike point get params
func (o *BikePointGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the bike point get params
func (o *BikePointGetParams) WithID(id string) *BikePointGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bike point get params
func (o *BikePointGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BikePointGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
