// Code generated by go-swagger; DO NOT EDIT.

package place

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPlaceGetParams creates a new PlaceGetParams object
// with the default values initialized.
func NewPlaceGetParams() *PlaceGetParams {
	var ()
	return &PlaceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPlaceGetParamsWithTimeout creates a new PlaceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPlaceGetParamsWithTimeout(timeout time.Duration) *PlaceGetParams {
	var ()
	return &PlaceGetParams{

		timeout: timeout,
	}
}

// NewPlaceGetParamsWithContext creates a new PlaceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPlaceGetParamsWithContext(ctx context.Context) *PlaceGetParams {
	var ()
	return &PlaceGetParams{

		Context: ctx,
	}
}

// NewPlaceGetParamsWithHTTPClient creates a new PlaceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPlaceGetParamsWithHTTPClient(client *http.Client) *PlaceGetParams {
	var ()
	return &PlaceGetParams{
		HTTPClient: client,
	}
}

/*PlaceGetParams contains all the parameters to send to the API endpoint
for the place get operation typically these are written to a http.Request
*/
type PlaceGetParams struct {

	/*ID
	  The id of the place, you can use the /Place/Types/{types} endpoint to get a list of places for a given type including their ids

	*/
	ID string
	/*IncludeChildren
	  Defaults to false. If true child places e.g. individual charging stations at a charge point while be included, otherwise just the URLs of any child places will be returned

	*/
	IncludeChildren *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the place get params
func (o *PlaceGetParams) WithTimeout(timeout time.Duration) *PlaceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the place get params
func (o *PlaceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the place get params
func (o *PlaceGetParams) WithContext(ctx context.Context) *PlaceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the place get params
func (o *PlaceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the place get params
func (o *PlaceGetParams) WithHTTPClient(client *http.Client) *PlaceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the place get params
func (o *PlaceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the place get params
func (o *PlaceGetParams) WithID(id string) *PlaceGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the place get params
func (o *PlaceGetParams) SetID(id string) {
	o.ID = id
}

// WithIncludeChildren adds the includeChildren to the place get params
func (o *PlaceGetParams) WithIncludeChildren(includeChildren *bool) *PlaceGetParams {
	o.SetIncludeChildren(includeChildren)
	return o
}

// SetIncludeChildren adds the includeChildren to the place get params
func (o *PlaceGetParams) SetIncludeChildren(includeChildren *bool) {
	o.IncludeChildren = includeChildren
}

// WriteToRequest writes these params to a swagger request
func (o *PlaceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IncludeChildren != nil {

		// query param includeChildren
		var qrIncludeChildren bool
		if o.IncludeChildren != nil {
			qrIncludeChildren = *o.IncludeChildren
		}
		qIncludeChildren := swag.FormatBool(qrIncludeChildren)
		if qIncludeChildren != "" {
			if err := r.SetQueryParam("includeChildren", qIncludeChildren); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
