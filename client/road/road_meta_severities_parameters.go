// Code generated by go-swagger; DO NOT EDIT.

package road

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

// NewRoadMetaSeveritiesParams creates a new RoadMetaSeveritiesParams object
// with the default values initialized.
func NewRoadMetaSeveritiesParams() *RoadMetaSeveritiesParams {

	return &RoadMetaSeveritiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRoadMetaSeveritiesParamsWithTimeout creates a new RoadMetaSeveritiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRoadMetaSeveritiesParamsWithTimeout(timeout time.Duration) *RoadMetaSeveritiesParams {

	return &RoadMetaSeveritiesParams{

		timeout: timeout,
	}
}

// NewRoadMetaSeveritiesParamsWithContext creates a new RoadMetaSeveritiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewRoadMetaSeveritiesParamsWithContext(ctx context.Context) *RoadMetaSeveritiesParams {

	return &RoadMetaSeveritiesParams{

		Context: ctx,
	}
}

// NewRoadMetaSeveritiesParamsWithHTTPClient creates a new RoadMetaSeveritiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRoadMetaSeveritiesParamsWithHTTPClient(client *http.Client) *RoadMetaSeveritiesParams {

	return &RoadMetaSeveritiesParams{
		HTTPClient: client,
	}
}

/*RoadMetaSeveritiesParams contains all the parameters to send to the API endpoint
for the road meta severities operation typically these are written to a http.Request
*/
type RoadMetaSeveritiesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the road meta severities params
func (o *RoadMetaSeveritiesParams) WithTimeout(timeout time.Duration) *RoadMetaSeveritiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the road meta severities params
func (o *RoadMetaSeveritiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the road meta severities params
func (o *RoadMetaSeveritiesParams) WithContext(ctx context.Context) *RoadMetaSeveritiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the road meta severities params
func (o *RoadMetaSeveritiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the road meta severities params
func (o *RoadMetaSeveritiesParams) WithHTTPClient(client *http.Client) *RoadMetaSeveritiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the road meta severities params
func (o *RoadMetaSeveritiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RoadMetaSeveritiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
