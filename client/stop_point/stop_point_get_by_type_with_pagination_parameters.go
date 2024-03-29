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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStopPointGetByTypeWithPaginationParams creates a new StopPointGetByTypeWithPaginationParams object
// with the default values initialized.
func NewStopPointGetByTypeWithPaginationParams() *StopPointGetByTypeWithPaginationParams {
	var ()
	return &StopPointGetByTypeWithPaginationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopPointGetByTypeWithPaginationParamsWithTimeout creates a new StopPointGetByTypeWithPaginationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopPointGetByTypeWithPaginationParamsWithTimeout(timeout time.Duration) *StopPointGetByTypeWithPaginationParams {
	var ()
	return &StopPointGetByTypeWithPaginationParams{

		timeout: timeout,
	}
}

// NewStopPointGetByTypeWithPaginationParamsWithContext creates a new StopPointGetByTypeWithPaginationParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopPointGetByTypeWithPaginationParamsWithContext(ctx context.Context) *StopPointGetByTypeWithPaginationParams {
	var ()
	return &StopPointGetByTypeWithPaginationParams{

		Context: ctx,
	}
}

// NewStopPointGetByTypeWithPaginationParamsWithHTTPClient creates a new StopPointGetByTypeWithPaginationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopPointGetByTypeWithPaginationParamsWithHTTPClient(client *http.Client) *StopPointGetByTypeWithPaginationParams {
	var ()
	return &StopPointGetByTypeWithPaginationParams{
		HTTPClient: client,
	}
}

/*StopPointGetByTypeWithPaginationParams contains all the parameters to send to the API endpoint
for the stop point get by type with pagination operation typically these are written to a http.Request
*/
type StopPointGetByTypeWithPaginationParams struct {

	/*Page*/
	Page int32
	/*Types*/
	Types []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) WithTimeout(timeout time.Duration) *StopPointGetByTypeWithPaginationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) WithContext(ctx context.Context) *StopPointGetByTypeWithPaginationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) WithHTTPClient(client *http.Client) *StopPointGetByTypeWithPaginationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) WithPage(page int32) *StopPointGetByTypeWithPaginationParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) SetPage(page int32) {
	o.Page = page
}

// WithTypes adds the types to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) WithTypes(types []string) *StopPointGetByTypeWithPaginationParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the stop point get by type with pagination params
func (o *StopPointGetByTypeWithPaginationParams) SetTypes(types []string) {
	o.Types = types
}

// WriteToRequest writes these params to a swagger request
func (o *StopPointGetByTypeWithPaginationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param page
	if err := r.SetPathParam("page", swag.FormatInt32(o.Page)); err != nil {
		return err
	}

	valuesTypes := o.Types

	joinedTypes := swag.JoinByFormat(valuesTypes, "")
	// path array param types
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedTypes) > 0 {
		if err := r.SetPathParam("types", joinedTypes[0]); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
