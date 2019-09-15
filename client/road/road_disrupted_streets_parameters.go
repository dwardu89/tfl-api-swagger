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

// NewRoadDisruptedStreetsParams creates a new RoadDisruptedStreetsParams object
// with the default values initialized.
func NewRoadDisruptedStreetsParams() *RoadDisruptedStreetsParams {
	var ()
	return &RoadDisruptedStreetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRoadDisruptedStreetsParamsWithTimeout creates a new RoadDisruptedStreetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRoadDisruptedStreetsParamsWithTimeout(timeout time.Duration) *RoadDisruptedStreetsParams {
	var ()
	return &RoadDisruptedStreetsParams{

		timeout: timeout,
	}
}

// NewRoadDisruptedStreetsParamsWithContext creates a new RoadDisruptedStreetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRoadDisruptedStreetsParamsWithContext(ctx context.Context) *RoadDisruptedStreetsParams {
	var ()
	return &RoadDisruptedStreetsParams{

		Context: ctx,
	}
}

// NewRoadDisruptedStreetsParamsWithHTTPClient creates a new RoadDisruptedStreetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRoadDisruptedStreetsParamsWithHTTPClient(client *http.Client) *RoadDisruptedStreetsParams {
	var ()
	return &RoadDisruptedStreetsParams{
		HTTPClient: client,
	}
}

/*RoadDisruptedStreetsParams contains all the parameters to send to the API endpoint
for the road disrupted streets operation typically these are written to a http.Request
*/
type RoadDisruptedStreetsParams struct {

	/*EndDate
	  Optional, The end time to filter on.

	*/
	EndDate strfmt.DateTime
	/*StartDate
	  Optional, the start time to filter on.

	*/
	StartDate strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) WithTimeout(timeout time.Duration) *RoadDisruptedStreetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) WithContext(ctx context.Context) *RoadDisruptedStreetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) WithHTTPClient(client *http.Client) *RoadDisruptedStreetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) WithEndDate(endDate strfmt.DateTime) *RoadDisruptedStreetsParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) SetEndDate(endDate strfmt.DateTime) {
	o.EndDate = endDate
}

// WithStartDate adds the startDate to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) WithStartDate(startDate strfmt.DateTime) *RoadDisruptedStreetsParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the road disrupted streets params
func (o *RoadDisruptedStreetsParams) SetStartDate(startDate strfmt.DateTime) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *RoadDisruptedStreetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endDate
	qrEndDate := o.EndDate
	qEndDate := qrEndDate.String()
	if qEndDate != "" {
		if err := r.SetQueryParam("endDate", qEndDate); err != nil {
			return err
		}
	}

	// query param startDate
	qrStartDate := o.StartDate
	qStartDate := qrStartDate.String()
	if qStartDate != "" {
		if err := r.SetQueryParam("startDate", qStartDate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}