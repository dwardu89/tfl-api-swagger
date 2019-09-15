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

// NewLineRouteParams creates a new LineRouteParams object
// with the default values initialized.
func NewLineRouteParams() *LineRouteParams {
	var ()
	return &LineRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineRouteParamsWithTimeout creates a new LineRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineRouteParamsWithTimeout(timeout time.Duration) *LineRouteParams {
	var ()
	return &LineRouteParams{

		timeout: timeout,
	}
}

// NewLineRouteParamsWithContext creates a new LineRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineRouteParamsWithContext(ctx context.Context) *LineRouteParams {
	var ()
	return &LineRouteParams{

		Context: ctx,
	}
}

// NewLineRouteParamsWithHTTPClient creates a new LineRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineRouteParamsWithHTTPClient(client *http.Client) *LineRouteParams {
	var ()
	return &LineRouteParams{
		HTTPClient: client,
	}
}

/*LineRouteParams contains all the parameters to send to the API endpoint
for the line route operation typically these are written to a http.Request
*/
type LineRouteParams struct {

	/*ServiceTypes
	  A comma seperated list of service types to filter on. Supported values: Regular, Night. Defaulted to 'Regular' if not specified

	*/
	ServiceTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line route params
func (o *LineRouteParams) WithTimeout(timeout time.Duration) *LineRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line route params
func (o *LineRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line route params
func (o *LineRouteParams) WithContext(ctx context.Context) *LineRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line route params
func (o *LineRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line route params
func (o *LineRouteParams) WithHTTPClient(client *http.Client) *LineRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line route params
func (o *LineRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceTypes adds the serviceTypes to the line route params
func (o *LineRouteParams) WithServiceTypes(serviceTypes []string) *LineRouteParams {
	o.SetServiceTypes(serviceTypes)
	return o
}

// SetServiceTypes adds the serviceTypes to the line route params
func (o *LineRouteParams) SetServiceTypes(serviceTypes []string) {
	o.ServiceTypes = serviceTypes
}

// WriteToRequest writes these params to a swagger request
func (o *LineRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
