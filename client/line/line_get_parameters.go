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

// NewLineGetParams creates a new LineGetParams object
// with the default values initialized.
func NewLineGetParams() *LineGetParams {
	var ()
	return &LineGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineGetParamsWithTimeout creates a new LineGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineGetParamsWithTimeout(timeout time.Duration) *LineGetParams {
	var ()
	return &LineGetParams{

		timeout: timeout,
	}
}

// NewLineGetParamsWithContext creates a new LineGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineGetParamsWithContext(ctx context.Context) *LineGetParams {
	var ()
	return &LineGetParams{

		Context: ctx,
	}
}

// NewLineGetParamsWithHTTPClient creates a new LineGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineGetParamsWithHTTPClient(client *http.Client) *LineGetParams {
	var ()
	return &LineGetParams{
		HTTPClient: client,
	}
}

/*LineGetParams contains all the parameters to send to the API endpoint
for the line get operation typically these are written to a http.Request
*/
type LineGetParams struct {

	/*Ids
	  A comma-separated list of line ids e.g. victoria,circle,N133. Max. approx. 20 ids.

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line get params
func (o *LineGetParams) WithTimeout(timeout time.Duration) *LineGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line get params
func (o *LineGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line get params
func (o *LineGetParams) WithContext(ctx context.Context) *LineGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line get params
func (o *LineGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line get params
func (o *LineGetParams) WithHTTPClient(client *http.Client) *LineGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line get params
func (o *LineGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the line get params
func (o *LineGetParams) WithIds(ids []string) *LineGetParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the line get params
func (o *LineGetParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *LineGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
