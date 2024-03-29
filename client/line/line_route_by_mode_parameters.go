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

// NewLineRouteByModeParams creates a new LineRouteByModeParams object
// with the default values initialized.
func NewLineRouteByModeParams() *LineRouteByModeParams {
	var ()
	return &LineRouteByModeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineRouteByModeParamsWithTimeout creates a new LineRouteByModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineRouteByModeParamsWithTimeout(timeout time.Duration) *LineRouteByModeParams {
	var ()
	return &LineRouteByModeParams{

		timeout: timeout,
	}
}

// NewLineRouteByModeParamsWithContext creates a new LineRouteByModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineRouteByModeParamsWithContext(ctx context.Context) *LineRouteByModeParams {
	var ()
	return &LineRouteByModeParams{

		Context: ctx,
	}
}

// NewLineRouteByModeParamsWithHTTPClient creates a new LineRouteByModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineRouteByModeParamsWithHTTPClient(client *http.Client) *LineRouteByModeParams {
	var ()
	return &LineRouteByModeParams{
		HTTPClient: client,
	}
}

/*LineRouteByModeParams contains all the parameters to send to the API endpoint
for the line route by mode operation typically these are written to a http.Request
*/
type LineRouteByModeParams struct {

	/*Modes
	  A comma-separated list of modes e.g. tube,dlr

	*/
	Modes []string
	/*ServiceTypes
	  A comma seperated list of service types to filter on. Supported values: Regular, Night. Defaulted to 'Regular' if not specified

	*/
	ServiceTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line route by mode params
func (o *LineRouteByModeParams) WithTimeout(timeout time.Duration) *LineRouteByModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line route by mode params
func (o *LineRouteByModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line route by mode params
func (o *LineRouteByModeParams) WithContext(ctx context.Context) *LineRouteByModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line route by mode params
func (o *LineRouteByModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line route by mode params
func (o *LineRouteByModeParams) WithHTTPClient(client *http.Client) *LineRouteByModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line route by mode params
func (o *LineRouteByModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModes adds the modes to the line route by mode params
func (o *LineRouteByModeParams) WithModes(modes []string) *LineRouteByModeParams {
	o.SetModes(modes)
	return o
}

// SetModes adds the modes to the line route by mode params
func (o *LineRouteByModeParams) SetModes(modes []string) {
	o.Modes = modes
}

// WithServiceTypes adds the serviceTypes to the line route by mode params
func (o *LineRouteByModeParams) WithServiceTypes(serviceTypes []string) *LineRouteByModeParams {
	o.SetServiceTypes(serviceTypes)
	return o
}

// SetServiceTypes adds the serviceTypes to the line route by mode params
func (o *LineRouteByModeParams) SetServiceTypes(serviceTypes []string) {
	o.ServiceTypes = serviceTypes
}

// WriteToRequest writes these params to a swagger request
func (o *LineRouteByModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesModes := o.Modes

	joinedModes := swag.JoinByFormat(valuesModes, "")
	// path array param modes
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedModes) > 0 {
		if err := r.SetPathParam("modes", joinedModes[0]); err != nil {
			return err
		}
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
