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

// NewLineLineRoutesByIdsParams creates a new LineLineRoutesByIdsParams object
// with the default values initialized.
func NewLineLineRoutesByIdsParams() *LineLineRoutesByIdsParams {
	var ()
	return &LineLineRoutesByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineLineRoutesByIdsParamsWithTimeout creates a new LineLineRoutesByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineLineRoutesByIdsParamsWithTimeout(timeout time.Duration) *LineLineRoutesByIdsParams {
	var ()
	return &LineLineRoutesByIdsParams{

		timeout: timeout,
	}
}

// NewLineLineRoutesByIdsParamsWithContext creates a new LineLineRoutesByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineLineRoutesByIdsParamsWithContext(ctx context.Context) *LineLineRoutesByIdsParams {
	var ()
	return &LineLineRoutesByIdsParams{

		Context: ctx,
	}
}

// NewLineLineRoutesByIdsParamsWithHTTPClient creates a new LineLineRoutesByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineLineRoutesByIdsParamsWithHTTPClient(client *http.Client) *LineLineRoutesByIdsParams {
	var ()
	return &LineLineRoutesByIdsParams{
		HTTPClient: client,
	}
}

/*LineLineRoutesByIdsParams contains all the parameters to send to the API endpoint
for the line line routes by ids operation typically these are written to a http.Request
*/
type LineLineRoutesByIdsParams struct {

	/*Ids
	  A comma-separated list of line ids e.g. victoria,circle,N133. Max. approx. 20 ids.

	*/
	Ids []string
	/*ServiceTypes
	  A comma seperated list of service types to filter on. Supported values: Regular, Night. Defaulted to 'Regular' if not specified

	*/
	ServiceTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) WithTimeout(timeout time.Duration) *LineLineRoutesByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) WithContext(ctx context.Context) *LineLineRoutesByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) WithHTTPClient(client *http.Client) *LineLineRoutesByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) WithIds(ids []string) *LineLineRoutesByIdsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithServiceTypes adds the serviceTypes to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) WithServiceTypes(serviceTypes []string) *LineLineRoutesByIdsParams {
	o.SetServiceTypes(serviceTypes)
	return o
}

// SetServiceTypes adds the serviceTypes to the line line routes by ids params
func (o *LineLineRoutesByIdsParams) SetServiceTypes(serviceTypes []string) {
	o.ServiceTypes = serviceTypes
}

// WriteToRequest writes these params to a swagger request
func (o *LineLineRoutesByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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