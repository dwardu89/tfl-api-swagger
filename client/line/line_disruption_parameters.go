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

// NewLineDisruptionParams creates a new LineDisruptionParams object
// with the default values initialized.
func NewLineDisruptionParams() *LineDisruptionParams {
	var ()
	return &LineDisruptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineDisruptionParamsWithTimeout creates a new LineDisruptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineDisruptionParamsWithTimeout(timeout time.Duration) *LineDisruptionParams {
	var ()
	return &LineDisruptionParams{

		timeout: timeout,
	}
}

// NewLineDisruptionParamsWithContext creates a new LineDisruptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineDisruptionParamsWithContext(ctx context.Context) *LineDisruptionParams {
	var ()
	return &LineDisruptionParams{

		Context: ctx,
	}
}

// NewLineDisruptionParamsWithHTTPClient creates a new LineDisruptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineDisruptionParamsWithHTTPClient(client *http.Client) *LineDisruptionParams {
	var ()
	return &LineDisruptionParams{
		HTTPClient: client,
	}
}

/*LineDisruptionParams contains all the parameters to send to the API endpoint
for the line disruption operation typically these are written to a http.Request
*/
type LineDisruptionParams struct {

	/*Ids
	  A comma-separated list of line ids e.g. victoria,circle,N133. Max. approx. 20 ids.

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line disruption params
func (o *LineDisruptionParams) WithTimeout(timeout time.Duration) *LineDisruptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line disruption params
func (o *LineDisruptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line disruption params
func (o *LineDisruptionParams) WithContext(ctx context.Context) *LineDisruptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line disruption params
func (o *LineDisruptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line disruption params
func (o *LineDisruptionParams) WithHTTPClient(client *http.Client) *LineDisruptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line disruption params
func (o *LineDisruptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the line disruption params
func (o *LineDisruptionParams) WithIds(ids []string) *LineDisruptionParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the line disruption params
func (o *LineDisruptionParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *LineDisruptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "")
	// path array param ids
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedIds) > 0 {
		if err := r.SetPathParam("ids", joinedIds[0]); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
