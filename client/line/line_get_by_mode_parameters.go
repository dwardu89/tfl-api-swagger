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

// NewLineGetByModeParams creates a new LineGetByModeParams object
// with the default values initialized.
func NewLineGetByModeParams() *LineGetByModeParams {
	var ()
	return &LineGetByModeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineGetByModeParamsWithTimeout creates a new LineGetByModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineGetByModeParamsWithTimeout(timeout time.Duration) *LineGetByModeParams {
	var ()
	return &LineGetByModeParams{

		timeout: timeout,
	}
}

// NewLineGetByModeParamsWithContext creates a new LineGetByModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineGetByModeParamsWithContext(ctx context.Context) *LineGetByModeParams {
	var ()
	return &LineGetByModeParams{

		Context: ctx,
	}
}

// NewLineGetByModeParamsWithHTTPClient creates a new LineGetByModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineGetByModeParamsWithHTTPClient(client *http.Client) *LineGetByModeParams {
	var ()
	return &LineGetByModeParams{
		HTTPClient: client,
	}
}

/*LineGetByModeParams contains all the parameters to send to the API endpoint
for the line get by mode operation typically these are written to a http.Request
*/
type LineGetByModeParams struct {

	/*Modes
	  A comma-separated list of modes e.g. tube,dlr

	*/
	Modes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line get by mode params
func (o *LineGetByModeParams) WithTimeout(timeout time.Duration) *LineGetByModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line get by mode params
func (o *LineGetByModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line get by mode params
func (o *LineGetByModeParams) WithContext(ctx context.Context) *LineGetByModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line get by mode params
func (o *LineGetByModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line get by mode params
func (o *LineGetByModeParams) WithHTTPClient(client *http.Client) *LineGetByModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line get by mode params
func (o *LineGetByModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModes adds the modes to the line get by mode params
func (o *LineGetByModeParams) WithModes(modes []string) *LineGetByModeParams {
	o.SetModes(modes)
	return o
}

// SetModes adds the modes to the line get by mode params
func (o *LineGetByModeParams) SetModes(modes []string) {
	o.Modes = modes
}

// WriteToRequest writes these params to a swagger request
func (o *LineGetByModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
