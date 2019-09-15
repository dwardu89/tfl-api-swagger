// Code generated by go-swagger; DO NOT EDIT.

package line

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

// NewLineSearchParams creates a new LineSearchParams object
// with the default values initialized.
func NewLineSearchParams() *LineSearchParams {
	var ()
	return &LineSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineSearchParamsWithTimeout creates a new LineSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineSearchParamsWithTimeout(timeout time.Duration) *LineSearchParams {
	var ()
	return &LineSearchParams{

		timeout: timeout,
	}
}

// NewLineSearchParamsWithContext creates a new LineSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineSearchParamsWithContext(ctx context.Context) *LineSearchParams {
	var ()
	return &LineSearchParams{

		Context: ctx,
	}
}

// NewLineSearchParamsWithHTTPClient creates a new LineSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineSearchParamsWithHTTPClient(client *http.Client) *LineSearchParams {
	var ()
	return &LineSearchParams{
		HTTPClient: client,
	}
}

/*LineSearchParams contains all the parameters to send to the API endpoint
for the line search operation typically these are written to a http.Request
*/
type LineSearchParams struct {

	/*Modes
	  Optionally filter by the specified modes

	*/
	Modes []string
	/*Query
	  Search term e.g victoria

	*/
	Query string
	/*ServiceTypes
	  A comma seperated list of service types to filter on. Supported values: Regular, Night. Defaulted to 'Regular' if not specified

	*/
	ServiceTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line search params
func (o *LineSearchParams) WithTimeout(timeout time.Duration) *LineSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line search params
func (o *LineSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line search params
func (o *LineSearchParams) WithContext(ctx context.Context) *LineSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line search params
func (o *LineSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line search params
func (o *LineSearchParams) WithHTTPClient(client *http.Client) *LineSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line search params
func (o *LineSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModes adds the modes to the line search params
func (o *LineSearchParams) WithModes(modes []string) *LineSearchParams {
	o.SetModes(modes)
	return o
}

// SetModes adds the modes to the line search params
func (o *LineSearchParams) SetModes(modes []string) {
	o.Modes = modes
}

// WithQuery adds the query to the line search params
func (o *LineSearchParams) WithQuery(query string) *LineSearchParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the line search params
func (o *LineSearchParams) SetQuery(query string) {
	o.Query = query
}

// WithServiceTypes adds the serviceTypes to the line search params
func (o *LineSearchParams) WithServiceTypes(serviceTypes []string) *LineSearchParams {
	o.SetServiceTypes(serviceTypes)
	return o
}

// SetServiceTypes adds the serviceTypes to the line search params
func (o *LineSearchParams) SetServiceTypes(serviceTypes []string) {
	o.ServiceTypes = serviceTypes
}

// WriteToRequest writes these params to a swagger request
func (o *LineSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesModes := o.Modes

	joinedModes := swag.JoinByFormat(valuesModes, "multi")
	// query array param modes
	if err := r.SetQueryParam("modes", joinedModes...); err != nil {
		return err
	}

	// path param query
	if err := r.SetPathParam("query", o.Query); err != nil {
		return err
	}

	valuesServiceTypes := o.ServiceTypes

	joinedServiceTypes := swag.JoinByFormat(valuesServiceTypes, "multi")
	// query array param serviceTypes
	if err := r.SetQueryParam("serviceTypes", joinedServiceTypes...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}