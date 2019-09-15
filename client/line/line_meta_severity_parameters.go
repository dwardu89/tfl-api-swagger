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

	strfmt "github.com/go-openapi/strfmt"
)

// NewLineMetaSeverityParams creates a new LineMetaSeverityParams object
// with the default values initialized.
func NewLineMetaSeverityParams() *LineMetaSeverityParams {

	return &LineMetaSeverityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineMetaSeverityParamsWithTimeout creates a new LineMetaSeverityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineMetaSeverityParamsWithTimeout(timeout time.Duration) *LineMetaSeverityParams {

	return &LineMetaSeverityParams{

		timeout: timeout,
	}
}

// NewLineMetaSeverityParamsWithContext creates a new LineMetaSeverityParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineMetaSeverityParamsWithContext(ctx context.Context) *LineMetaSeverityParams {

	return &LineMetaSeverityParams{

		Context: ctx,
	}
}

// NewLineMetaSeverityParamsWithHTTPClient creates a new LineMetaSeverityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineMetaSeverityParamsWithHTTPClient(client *http.Client) *LineMetaSeverityParams {

	return &LineMetaSeverityParams{
		HTTPClient: client,
	}
}

/*LineMetaSeverityParams contains all the parameters to send to the API endpoint
for the line meta severity operation typically these are written to a http.Request
*/
type LineMetaSeverityParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line meta severity params
func (o *LineMetaSeverityParams) WithTimeout(timeout time.Duration) *LineMetaSeverityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line meta severity params
func (o *LineMetaSeverityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line meta severity params
func (o *LineMetaSeverityParams) WithContext(ctx context.Context) *LineMetaSeverityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line meta severity params
func (o *LineMetaSeverityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line meta severity params
func (o *LineMetaSeverityParams) WithHTTPClient(client *http.Client) *LineMetaSeverityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line meta severity params
func (o *LineMetaSeverityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LineMetaSeverityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
